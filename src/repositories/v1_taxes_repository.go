package repositories

import (
	"../database"
	"../models"
	"github.com/jinzhu/gorm"
)

type V1TaxesRepository struct {
	DB gorm.DB
}

func V1TaxesRepositoryHandler() (V1TaxesRepository) {
	repository := V1TaxesRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V1TaxesRepository) GetByCode(code int64) (models.Taxes, error) {

	taxesResponse := models.Taxes{}

	query := repository.DB.Table("taxes")
	query = query.Where("code=?", code)
	query = query.First(&taxesResponse)

	return taxesResponse, query.Error

}
