package repositories

import (
	"../database"
	"../models"
	"github.com/jinzhu/gorm"
)

type V1TaxConfigurationsRepository struct {
	DB gorm.DB
}

func V1TaxConfigurationsRepositoryHandler() (V1TaxConfigurationsRepository) {
	repository := V1TaxConfigurationsRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V1TaxConfigurationsRepository) GetListByTaxId(taxId int64) ([]models.TaxConfigurations, error) {

	var taxConfigurationsResponse []models.TaxConfigurations

	query := repository.DB.Table("tax_configurations")
	query = query.Where("tax_id=?", taxId)
	query = query.Order("priority")
	query = query.Find(&taxConfigurationsResponse)

	return taxConfigurationsResponse, query.Error

}
