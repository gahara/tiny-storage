package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet(DATABASE_KEY).(*gorm.DB)
}
