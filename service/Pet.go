package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// PetService defines the interface for managing Pets.
type PetService interface {
	Insert(Pet *model.Pet) *error.Error
	GetList() ([]model.Pet, *error.Error)
}

type PetServiceImpl struct {
	db *db.DataStore
}

func NewPetService(db *db.DataStore) PetService {
	return &PetServiceImpl{db: db}
}

// Function Implementation

func (s PetServiceImpl) GetList() ([]model.Pet, *error.Error) {
	return s.db.GetListPet()
}

func (s *PetServiceImpl) Insert(Pet *model.Pet) *error.Error {
	return s.db.InsertPet(Pet)
}