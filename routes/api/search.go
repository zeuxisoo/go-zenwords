package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// SearchPost to handle the search action
func SearchPost(c *gin.Context) {
	content := c.PostForm("content")

	result, matched := keywords.ExtraSearch(content)

	c.JSON(http.StatusOK, gin.H{
		"search": result,
		"matched": matched,
	})
}
