package common

import (
	"net/http"
	"payeasy/shared/model"

	"github.com/gin-gonic/gin"
)

func SendPagedResponse(c *gin.Context, data []interface{}, paging model.Paging, message string) {
	c.JSON(http.StatusOK, &model.PagedResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}