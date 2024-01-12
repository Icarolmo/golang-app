package graphicsendpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GraphicsPage(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}