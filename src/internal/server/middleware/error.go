package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"s3/src/pkg"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		for _, err := range context.Errors {
			switch e := err.Err.(type) {
			case pkg.HttpError:
				log.Println(err)
				context.AbortWithStatusJSON(e.StatusCode, gin.H{"results": gin.H{"error": e}})
			default:
				log.Println(err)
				context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"results": gin.H{"error": gin.H{"description": "Something went wrong", "statusCode": 500}}})
			}
		}
	}
}
