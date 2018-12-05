package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchPost to handle the search action
func SearchPost(c *gin.Context) {
	c.String(http.StatusOK, "ApiSearchPost")
}
