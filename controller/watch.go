package controller

import (
	"cd_platform/api"
	"cd_platform/ext"
	"cd_platform/pkg/watch"
	"cd_platform/util"

	"github.com/gin-gonic/gin"
)

type WatchController struct {
	BaseController
	WatchService watch.WatchService
}

func NewWatchController() *WatchController {
	watchSvc := watch.NewService(ext.MiddleWare)
	return &WatchController{
		WatchService: watchSvc,
	}
}

func (ctrl *WatchController) GetDeploymentByName(c *gin.Context) {
	n := c.Param("name")
	ns := c.Param("namespace")
	ret, err := ctrl.WatchService.GetDeploymentByName(c, ns, n)
	if err != nil {
		util.Logger.Errorf("controller.GetDeploymentByName err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
	}
	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetDeploymentByLabel(c *gin.Context) {
	para := &api.SelectorCondList{}
	if err := c.BindJSON(&para); err != nil {
		util.Logger.Errorf("controller.GetDeploymentByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
	}

	ret, err := ctrl.WatchService.GetDeploymentByLabel(c, para)
	if err != nil {
		util.Logger.Errorf("controller.GetDeploymentByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
	}

	ctrl.Jsonify(c, 200, ret, "")
}
