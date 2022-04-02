package v1

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func reponseError(ctx *gin.Context, code int, message string) {
	ctx.JSONP(code, gin.H{
		"message": message,
	})
}

func reponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
}

func marshalJson(data interface{}, dataStruct interface{}) error {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(dataJson, &dataStruct); err != nil {
		return err
	}
	return nil
}
