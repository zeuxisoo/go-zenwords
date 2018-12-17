package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// ContentReplacePost to handle the search action
func ContentReplacePost(c *gin.Context) {
	content := c.PostForm("content")

	result := keywords.Filter(content)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
