package helpers

import (
	"github.com/rinaldypasya/src/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type ErrorHelper struct {
}

func ErrorHelperHandler() (ErrorHelper) {
	return ErrorHelper{}
}

func (handler *ErrorHelper) HTTPResponseError(context *gin.Context, e error, defaultErrorCode int) {

	if nil != e {
		switch e.Error() {
		case "record not found":
			defaultErrorCode = constants.ResourceNotFound
			break
		}
	}

	errorConstant := constants.GetErrorConstant(defaultErrorCode)
	context.JSON(errorConstant.HttpCode, gin.H{
		"code":    defaultErrorCode,
		"message": errorConstant.Message,
	})

	if nil != e {
		if _, ok := e.(*mysql.MySQLError); !ok {
			fmt.Println(e)
		}
		panic(e)
	}

}
