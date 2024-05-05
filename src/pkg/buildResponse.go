package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"s3/src/internal/customTypes"
)

func ResponseOK(c *gin.Context, text string) {
	c.JSON(http.StatusOK, gin.H{
		"message": text,
	})
}

func BuildFilesResponse(message customTypes.Message, data customTypes.Files) customTypes.FilesResponse {
	res := customTypes.FilesResponse{}
	res.Results.Data = data
	res.Results.Message = message
	return res
}
