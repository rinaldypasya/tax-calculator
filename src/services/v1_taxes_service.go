package services

import (
	"../objects"
	"../repositories"
	"fmt"
	"strconv"
)

type V1TaxesService struct {
	repository                  repositories.V1TaxesRepository
	taxConfigurationsRepository repositories.V1TaxConfigurationsRepository
}

func V1TaxesServiceHandler() (V1TaxesService) {
	service := V1TaxesService{
		repository:                  repositories.V1TaxesRepositoryHandler(),
		taxConfigurationsRepository: repositories.V1TaxConfigurationsRepositoryHandler(),
	}
	return service
}

func (service *V1TaxesService) CalculateTax(requestObject objects.V1TaxesObjectRequest) (objects.V1TaxesObjectResponse, error) {
	return service.calculateTax(requestObject)
}

func (service *V1TaxesService) CalculateTaxBulk(requestObject []objects.V1TaxesObjectRequest) ([]objects.V1TaxesObjectResponse, error) {

	var response []objects.V1TaxesObjectResponse

	for _, element := range requestObject {
		result, err := service.calculateTax(element)
		if err != nil {
			return []objects.V1TaxesObjectResponse{}, err
		}
		response = append(response, result)
	}

	return response, nil

}

func (service *V1TaxesService) calculateTax(requestObject objects.V1TaxesObjectRequest) (objects.V1TaxesObjectResponse, error) {

	var err error

	taxes, err := service.repository.GetByCode(requestObject.TaxCode)
	if nil != err {
		return objects.V1TaxesObjectResponse{}, err
	}

	taxId, err := strconv.ParseInt(fmt.Sprintf("%d", taxes.ID), 10, 64)
	if nil != err {
		return objects.V1TaxesObjectResponse{}, err
	}

	taxConfigurations, err := service.taxConfigurationsRepository.GetListByTaxId(taxId)
	if nil != err {
		return objects.V1TaxesObjectResponse{}, err
	}

	var totalTax = requestObject.Price

	for _, element := range taxConfigurations {

		if element.Type == "percentage_of_price" {
			totalTax = totalTax * element.Value / 100
		}

		if element.Type == "additional_charge" {
			totalTax = totalTax + element.Value
		}

	}

	refundable := "NO"
	if taxes.IsRefundable {
		refundable = "YES"
	}

	response := objects.V1TaxesObjectResponse{
		Name:       requestObject.Name,
		TaxCode:    requestObject.TaxCode,
		Type:       taxes.Name,
		Refundable: refundable,
		Price:      requestObject.Price,
		Tax:        totalTax,
		Amount:     requestObject.Price + totalTax,
	}

	return response, nil

}
