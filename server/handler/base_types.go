package handler

import (
	"jobfair2024/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrCode int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ResponseWithPagination struct {
	Response
	Total      int         `json:"total"`
	Pagination *Pagination `json:"pagination"`
}

type BaseResponse struct {
	TraceId string `json:"trace_id"`
}

func responseBadRequestError(c *gin.Context, serviceErr pkg.ServiceError) {
	c.JSON(http.StatusBadRequest, Response{
		ErrCode: serviceErr.ErrCode(),
		Msg:     serviceErr.Error(),
		Data:    "",
	})
}

func responseNotFoundError(c *gin.Context, serviceErr pkg.ServiceError) {
	c.JSON(http.StatusNotFound, Response{
		ErrCode: serviceErr.ErrCode(),
		Msg:     serviceErr.Error(),
		Data:    "",
	})
}

func responseServerError(c *gin.Context, serviceErr pkg.ServiceError) {
	c.JSON(http.StatusInternalServerError, Response{
		ErrCode: serviceErr.ErrCode(),
		Msg:     serviceErr.Error(),
		Data:    "",
	})
}

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Msg:  "200 OK",
		Data: data,
	})
}

func responseNotAuthorized(c *gin.Context, serviceErr pkg.ServiceError) {
	c.JSON(http.StatusUnauthorized, Response{
		ErrCode: serviceErr.ErrCode(),
		Msg:     serviceErr.Error(),
		Data:    "",
	})
}
