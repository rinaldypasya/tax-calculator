package controllers

import (
	"../constants"
	"../helpers"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V1TaxesController struct {
	service     services.V1TaxesService
	errorHelper helpers.ErrorHelper
}

func V1TaxesControllerHandler(router *gin.Engine) {

	handler := &V1TaxesController{
		service:     services.V1TaxesServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("v1/taxes")
	{
		group.POST("", handler.CalculateTax)
		group.POST("bulk", handler.CalculateTaxBulk)
	}

}

func (handler *V1TaxesController) CalculateTax(context *gin.Context) {

	requestObject := objects.V1TaxesObjectRequest{}
	context.ShouldBind(&requestObject)

	if requestObject.Price == 0 || requestObject.TaxCode == 0 {
		handler.errorHelper.HTTPResponseError(context, nil, constants.RequestParameterInvalid)
		return
	}

	result, err := handler.service.CalculateTax(requestObject)

	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
		return
	}

	context.JSON(http.StatusOK, result)

}

func (handler *V1TaxesController) CalculateTaxBulk(context *gin.Context) {

	var requestObject []objects.V1TaxesObjectRequest
	context.ShouldBind(&requestObject)

	for _, element := range requestObject {

		if element.Price == 0 || element.TaxCode == 0 {
			handler.errorHelper.HTTPResponseError(context, nil, constants.RequestParameterInvalid)
			return
		}

	}

	result, err := handler.service.CalculateTaxBulk(requestObject)

	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
		return
	}

	context.JSON(http.StatusOK, result)

}
