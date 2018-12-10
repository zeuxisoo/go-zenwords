package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// SearchPost to handle the search action
func SearchPost(c *gin.Context) {
	content := c.PostForm("content")

	fmt.Println(content)

	result, _ := keywords.Search(content)

	c.JSON(http.StatusOK, gin.H{
		"search": result,
	})
}
