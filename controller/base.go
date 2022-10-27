package controller

import (
	"cd_platform/api"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (ctrl *BaseController) Jsonify(ctx *gin.Context, code int, data interface{}, message string) {
	body := &api.ResponseBody{}
	body.Code = code
	body.Message = message
	body.Data = data
	ctx.JSON(code, body)
}
