package utils

import "github.com/gin-gonic/gin"

type PaginationMeta struct {
	TotalData   int64 `json:"total_data"`
	Limit       int   `json:"limit"`
	CurrentPage int   `json:"current_page"`
	TotalPage   int   `json:"total_page"`
}

func ResponseSuccess(c *gin.Context, status int, data interface{}, message string, pagination ...PaginationMeta) {

	response := gin.H{
		"success": true,
	}

	if data != nil {
		response["data"] = data
	}

	if message != "" {
		response["message"] = message
	}

	if len(pagination) > 0 {
		response["meta"] = pagination[0]
	}

	c.JSON(status, response)
}

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
	})
}
