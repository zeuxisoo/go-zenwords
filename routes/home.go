package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/utils"
)

// HomeIndexGet return the default home page
func HomeIndexGet(c *gin.Context) {
	c.String(http.StatusOK, "ZenWords")
}

// HomeRobotsTxtGet return the robots.txt from current directory if file exists
func HomeRobotsTxtGet(c *gin.Context) {
	if utils.IsFileExists("robots.txt") {
		c.File("robots.txt")
	} else {
		c.String(http.StatusOK, "User-agent: *\nDisallow:")
	}
}
