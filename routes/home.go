package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeGet return the default home page
func HomeGet(c *gin.Context) {
	c.String(http.StatusOK, "ZenWords")
}
