package controller

import (
	"cd_platform/common"
	"cd_platform/ext"
	"cd_platform/pkg/exec"
	"cd_platform/util"
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
	var args common.CreateArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateDeployment err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	err := ctrl.ExecService.CreateDeployment(c, args.Project, args.Metadata)
	if err != nil {
		util.Logger.Errorf("controller.CreateDeployment err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, nil, "success")
}

func (ctrl *ExecController) CreateNamespace(c *gin.Context) {
	var args common.CreateArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateNamespace err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	err := ctrl.ExecService.CreateNamespace(c, args.Project, args.Metadata)
	if err != nil {
		util.Logger.Errorf("controller.CreateNamespace err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, nil, "success")
}
