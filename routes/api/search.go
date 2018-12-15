package api

import (
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zeuxisoo/go-zenwords/pkg/keywords"
)

// SearchPost to handle the search action
func SearchPost(c *gin.Context) {
	kind := c.PostForm("kind")
	key  := c.PostForm("key")

	var result string
	var matched bool

	if strings.ToLower(kind) == "extra" {
		result, matched = keywords.ExtraSearch(key)
	}else{
		values, state := keywords.PrefixSearch(key)

		result  = strings.Join(values, ",")
		matched = state
	}

	c.JSON(http.StatusOK, gin.H{
		"search": result,
		"matched": matched,
	})
}
