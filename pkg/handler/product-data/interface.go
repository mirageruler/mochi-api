package productdata

import "github.com/gin-gonic/gin"

type IHandler interface {
	ProductBotCommand(c *gin.Context)
}
