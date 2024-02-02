package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type modelConfig struct {
	Model  string `yaml:"model"`
	Fields []struct {
		Name       string `yaml:"name"`
		Type       string `yaml:"type"`
		Binding    string `yaml:"binding"`
		Constraint string `yaml:"constraint"`
	} `yaml:"fields"`
	Routers []struct {
		Method string `yaml:"method"`
		Path   string `yaml:"path"`
	} `yaml:"routers"`
}

var config modelConfig

func init() {
	yamlFile, err := ioutil.ReadFile("animal.yaml")
	if err != nil {
		log.Fatalf("error load yaml")
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error unmarshalling yaml")
	}
}

func main() {
	modelReplacement := generateModelReplacements(config)
	dbReplacement := generateDbReplacements(config)
	routeReplacement := generateRouteReplacements(config)
	generateFile(config, "model", modelReplacement)
	generateFile(config, "db", dbReplacement)
	generateFile(config, "service", dbReplacement)
	generateRouteFile(config, routeReplacement)
}

func generateFile(config modelConfig, file string, Modelreplacement map[string]string) {

	templateFile := fmt.Sprintf("./template/%s_generated.go.txt", file)
	content, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Convert content to string
	contentStr := string(content)
	// Perform replacements
	for placeholder, replacement := range Modelreplacement {
		contentStr = strings.ReplaceAll(contentStr, placeholder, replacement)
	}
	// Define the target folder
	targetFolder := "./" + file + "/"
	modifiedContent := []byte(contentStr)
	// Write the modified content to a new file in the target folder
	targetFile := targetFolder + Modelreplacement["{.model}"] + ".go"
	err = ioutil.WriteFile(targetFile, modifiedContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File successfully created:", targetFile)
}

func generateRouteFile(config modelConfig, Modelreplacement map[string]string) {
	content, err := ioutil.ReadFile("./route/route.txt")
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return
	}

	// Convert content to string
	contentStr := string(content)
	// Perform replacements
	for placeholder, replacement := range Modelreplacement {
		contentStr = strings.ReplaceAll(contentStr, placeholder, replacement)
	}
	// Define the target folder
	targetFolder := "./route/"
	modifiedContent := []byte(contentStr)
	// Write the modified content to a new file in the target folder
	targetFile := targetFolder + "route.txt"
	err = ioutil.WriteFile(targetFile, modifiedContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File successfully created:", targetFile)
}

func generateModelReplacements(config modelConfig) map[string]string {
	var model_definition string
	var model_input string
	var PopulateFromDTOInput string

	for _, field := range config.Fields {
		tag := "`gorm:\"" + field.Constraint + "\"`"
		model_definition += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, tag)
	}

	for _, field := range config.Fields {
		tag := fmt.Sprintf("`json:\"%s\" binding:\"%s\"`", field.Name, field.Binding)
		model_input += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, tag)
	}

	replacements := map[string]string{
		"{.model}":                config.Model,
		"{.model_definition}":     model_definition,
		"{.model_input}":          model_input,
		"{.PopulateFromDTOInput}": PopulateFromDTOInput,
	}
	return replacements
}

func generateDbReplacements(config modelConfig) map[string]string {
	replacements := map[string]string{
		"{.model}": config.Model,
	}
	return replacements
}

func generateRouteReplacements(config modelConfig) map[string]string {
	var handler string
	var route string

	for _, r := range config.Routers {
		switch r.Method {
		case "get":
			route += fmt.Sprintf("\tr.GET(\"%s\", %shandler.GetList)\n", r.Path, config.Model)
		case "post":
			route += fmt.Sprintf("\tr.POST(\"%s\", %shandler.Insert)\n", r.Path, config.Model)
		}
	}
	replacements := map[string]string{
		"{.model}":        config.Model,
		"//{.NewHandler}": handler + "\n //{.NewHandler}",
		"//{.NewRoute}":   route + "\n //{.NewRoute}",
	}
	return replacements
}
