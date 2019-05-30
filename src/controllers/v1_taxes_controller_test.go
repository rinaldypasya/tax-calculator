package controllers

import (
	"../database"
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"testing"
)

func init() {
	godotenv.Load("../.env")
	database.Initialize()
}

func (handler *V1TaxesController) TestCalculateTax(t *testing.T) {

	message := map[string]interface{}{
		"name":     "movie",
		"tax_code": 3,
		"price":    150,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("/v1/taxes", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	_, errorAttribute := result["name"]

	if !errorAttribute {
		t.Errorf("Attribute name doesn't exists")
	}

}

func (handler *V1TaxesController) CalculateTaxBulk(t *testing.T) {

	message := map[string]interface{}{
		"name":     "movie",
		"tax_code": 3,
		"price":    150,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("/v1/taxes", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result []map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for _, item := range result {
		_, errorAttribute := item["name"]
		if !errorAttribute {
			t.Errorf("Attribute name doesn't exists")
		}
	}

}
