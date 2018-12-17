package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// SearchPost to handle the search action
func SearchPost(c *gin.Context) {
	content := c.PostForm("content")

	result := keywords.Filter(content)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
