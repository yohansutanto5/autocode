package model

type Pet struct {
		ID int `gorm:"primaryKey;autoIncrement"`
	Stupid string `gorm:"primaryKey;autoIncrement"`
	Arm string `gorm:"notNull"`
	Isfurry bool `gorm:"notNull"`

}

// DTO input and func to populate it
type PetInput struct {
		ID int `json:"ID" binding:"required"`
	Stupid string `json:"Stupid" binding:"required"`
	Arm string `json:"Arm" binding:"none"`
	Isfurry bool `json:"Isfurry" binding:"none"`

}
func (m *Pet) PopulateFromDTOInput(input PetInput) {
  
}