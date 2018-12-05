package home

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/utils"
)

// IndexGet return the default home page
func IndexGet(c *gin.Context) {
	c.String(http.StatusOK, "ZenWords")
}

// RobotsTxtGet return the robots.txt from current directory if file exists
func RobotsTxtGet(c *gin.Context) {
	if utils.IsFileExists("robots.txt") {
		c.File("robots.txt")
	} else {
		c.String(http.StatusOK, "User-agent: *\nDisallow:")
	}
}
