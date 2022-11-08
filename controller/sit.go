package controller

import (
	"cd_platform/common"
	"cd_platform/ext"
	"cd_platform/pkg/sit"
	"cd_platform/util"
	"github.com/gin-gonic/gin"
)

type SitController struct {
	BaseController
	SitService sit.SitService
}

func NewSitController() *SitController {
	psvc := sit.NewService(ext.MiddleWare)
	return &SitController{
		SitService: psvc,
	}
}

func (ctrl *SitController) CreateDeployment(c *gin.Context) {
	var args common.CreateArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateDeployment err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	err := ctrl.SitService.CreateDeployment(c, args.Project, args.Metadata)
	if err != nil {
		util.Logger.Errorf("controller.CreateDeployment err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, nil, "success")
}
