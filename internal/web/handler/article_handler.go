package handler

import "github.com/gin-gonic/gin"

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (h *ArticleHandler) Edit(ctx *gin.Context) {

}
