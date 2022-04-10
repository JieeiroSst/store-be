package v1

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type ErroMessage struct {
	Timestamp int    `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

func reponseError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, ErroMessage{
		Timestamp: int(time.Now().Unix()),
		Status:    400,
		Message:   message,
		Path:      ctx.Request.URL.Path,
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
