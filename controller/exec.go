package controller

import (
	"cd_platform/ext"
	"cd_platform/pkg/exec"
	"github.com/gin-gonic/gin"
)

type ExecController struct {
	BaseController
	ExecService exec.ExecService
}

func NewExecController() *ExecController {
	exc := exec.NewService(ext.MiddleWare)
	return &ExecController{
		ExecService: exc,
	}
}

func (ctrl *ExecController) CreateDeployment(c *gin.Context) {

}
