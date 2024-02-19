package handler

import (
	"jobfair2024/middleware"
	"jobfair2024/pkg"
	"jobfair2024/usecase"

	"github.com/gin-gonic/gin"
)

type NotificationReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	ReceiverID int64  `json:"receiver_id"`
}

func (h *JobFairHandler) SendNotification(c *gin.Context) {
	var req NotificationReq

	if err := c.ShouldBind(&req); err != nil {
		responseBadRequestError(c, pkg.BindingFailure)
		return
	}

	userInfo := middleware.GetUserInfoFromContext(c)
	if userInfo == nil {
		return
	}

	err := h.notificationUsecase.CreateNotification(c, *userInfo, usecase.NotificationInfo{
		Title:      req.Title,
		Content:    req.Content,
		ReceiverID: req.ReceiverID,
		SenderID:   &userInfo.ID,
	})

	if err != nil {
		responseServerError(c, pkg.ParseError(err))
		return
	}

	responseSuccess(c, "ok")
}
