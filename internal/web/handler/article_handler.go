package handler

import (
	jwtx "github.com/LEILEI0628/GinPro/middleware/jwt"
	loggerx "github.com/LEILEI0628/GinPro/middleware/logger"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/LEILEI0628/GoWeb/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleHandler struct {
	svc *service.ArticleService
	l   loggerx.Logger
}

func NewArticleHandler(svc *service.ArticleService, logger loggerx.Logger) *ArticleHandler {
	return &ArticleHandler{svc: svc, l: logger}
}

func (h *ArticleHandler) Edit(ctx *gin.Context) {
	type Req struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	//TODO 检测输入

	c := ctx.MustGet("claims")
	claims, ok := c.(*jwtx.UserClaims)
	if !ok {
		//ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
		h.l.Error("未发现用户的session信息")
		return
	}
	UID := claims.UID

	// 调用svc
	id, err := h.svc.Save(ctx.Request.Context(), domain.Article{Id: req.Id, Title: req.Title, Content: req.Content,
		Author: domain.Author{Id: UID}})
	if err != nil {
		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
		h.l.Error("article保存失败", loggerx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, Result{Msg: "OK", Data: id})
}
