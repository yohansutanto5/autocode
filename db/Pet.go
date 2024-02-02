package db

import (
	"app/model"
	"app/pkg/error"
)

func (d *DataStore) GetListPet() (Pets []model.Pet, err *error.Error) {
	e := d.Db.Find(&Pets).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) InsertPet(Pet *model.Pet) (err *error.Error) {
	e := d.Db.Create(Pet).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) DeletePetByID(id int) (err *error.Error) {
	e := d.Db.Where("id = ?", id).Delete(&model.Pet{}).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) UpdatePet(Pet *model.Pet) (err *error.Error) {
	e := d.Db.Save(&Pet).Error
	err.ParseMysqlError(e)
	return
}