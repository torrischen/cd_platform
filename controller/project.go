package controller

import (
	"cd_platform/common"
	"cd_platform/pkg"
	"cd_platform/util"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	BaseController
	ExecService pkg.ExecService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		ExecService: pkg.ExService,
	}
}

func (ctrl *ProjectController) InitProject(c *gin.Context) {
	var args common.InitProjectArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.CreateNamespace(c, args.Project); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) DeployProject(c *gin.Context) {

}
