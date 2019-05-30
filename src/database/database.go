package database

import (
	"../models"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"net/url"
	"os"
	"time"
)

var connection *gorm.DB

func init() {
	godotenv.Load()
}

func Initialize() (*gorm.DB, error) {

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	defaultTimezone := os.Getenv("SERVER_TIMEZONE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape(defaultTimezone),
	)

	var err error

	connection, err = gorm.Open("mysql", connectionString)
	if nil != err {

		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)

		errorOutput.Println("")
		errorOutput.Println("!!! Warning")
		errorOutput.Println(fmt.Sprintf("Failed connected to database %s", connectionString))
		errorOutput.Println("")

	} else {

		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)

		successOutput.Println("")
		successOutput.Println("!!! Info")
		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", connectionString))
		successOutput.Println("")

		if err == nil {

			// create schemas
			connection.AutoMigrate(&models.Taxes{})
			connection.AutoMigrate(&models.TaxConfigurations{})

			// insert data seeder

			/*
			('1', '1', 'Food and Beverage', '1', '2019-01-27 09:47:10', '2019-01-27 09:47:10', NULL),
			('2', '2', 'Tobacco', '0', '2019-01-27 09:47:10', '2019-01-27 09:47:10', NULL),
			('3', '3', 'Entertainment', '0', '2019-01-27 09:47:10', '2019-01-27 09:47:10', NULL);
			 */

			connection.Table("taxes").Create(models.Taxes{
				ID:           1,
				Code:         1,
				Name:         "Food and Beverage",
				IsRefundable: true,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			})

			connection.Table("taxes").Create(models.Taxes{
				ID:           2,
				Code:         2,
				Name:         "Tobacco",
				IsRefundable: false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			})

			connection.Table("taxes").Create(models.Taxes{
				ID:           3,
				Code:         3,
				Name:         "Entertainment",
				IsRefundable: false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			})

			/*
			('1', '1', NULL, NULL, 'percentage_of_price', '10', '1', '2019-01-27 09:49:14', '2019-01-27 09:49:14', NULL),
			('2', '2', NULL, NULL, 'percentage_of_price', '2', '1', '2019-01-27 09:49:14', '2019-01-27 09:49:14', NULL),
			('3', '2', NULL, NULL, 'additional_charge', '10', '2', '2019-01-27 09:49:14', '2019-01-27 09:49:14', NULL),
			('4', '3', '100', NULL, 'additional_charge', '-100', '1', '2019-01-27 09:49:14', '2019-01-27 09:49:14', NULL),
			('5', '3', '100', NULL, 'percentage_of_price', '1', '2', '2019-01-27 09:49:14', '2019-01-27 09:49:14', NULL);
			 */

			connection.Table("tax_configurations").Create(models.TaxConfigurations{
				ID:        1,
				TaxId:     1,
				MinPrice:  nil,
				MaxPrice:  nil,
				Type:      "percentage_of_price",
				Value:     10,
				Priority:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

			connection.Table("tax_configurations").Create(models.TaxConfigurations{
				ID:        2,
				TaxId:     2,
				MinPrice:  nil,
				MaxPrice:  nil,
				Type:      "percentage_of_price",
				Value:     2,
				Priority:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

			connection.Table("tax_configurations").Create(models.TaxConfigurations{
				ID:        3,
				TaxId:     2,
				MinPrice:  nil,
				MaxPrice:  nil,
				Type:      "additional_charge",
				Value:     10,
				Priority:  2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

			connection.Table("tax_configurations").Create(models.TaxConfigurations{
				ID:        4,
				TaxId:     3,
				MinPrice:  nil,
				MaxPrice:  nil,
				Type:      "additional_charge",
				Value:     -100,
				Priority:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

			connection.Table("tax_configurations").Create(models.TaxConfigurations{
				ID:        5,
				TaxId:     3,
				MinPrice:  nil,
				MaxPrice:  nil,
				Type:      "percentage_of_price",
				Value:     1,
				Priority:  2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})

		}

	}

	zapLog, _ := zap.NewProduction()
	connection.SetLogger(customLogger(zapLog))

	return connection, nil

}

func GetConnection() *gorm.DB {

	for ; connection == nil; {
		time.Sleep(1 * time.Second)
		connection, _ = Initialize()
	}

	return connection

}

func customLogger(zap *zap.Logger) *customLoggerStruct {
	return &customLoggerStruct{
		zap: zap,
	}
}

type customLoggerStruct struct {
	zap *zap.Logger
}

func (l *customLoggerStruct) Print(values ...interface{}) {
	var additionalString = ""
	for _, item := range values {
		if _, ok := item.(string); ok {
			additionalString = additionalString + fmt.Sprintf("\n%v", item)
		}
		if err, ok := item.(*mysql.MySQLError); ok {
			err.Message = err.Message + additionalString
			// integrate 3rd party logging error here :)
			fmt.Println(err)
		}
	}
}
