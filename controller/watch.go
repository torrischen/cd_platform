package controller

import (
	"cd_platform/common"
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
		return
	}
	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetDeploymentByLabel(c *gin.Context) {
	para := &common.SelectorCondList{}
	if err := c.BindJSON(&para); err != nil {
		util.Logger.Errorf("controller.GetDeploymentByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ret, err := ctrl.WatchService.GetDeploymentByLabel(c, para)
	if err != nil {
		util.Logger.Errorf("controller.GetDeploymentByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetPodByName(c *gin.Context) {
	n := c.Param("name")
	ns := c.Param("namespace")
	ret, err := ctrl.WatchService.GetPodByName(c, ns, n)
	if err != nil {
		util.Logger.Errorf("controller.GetPodByName err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}
	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetPodByLabel(c *gin.Context) {
	para := &common.SelectorCondList{}
	if err := c.BindJSON(&para); err != nil {
		util.Logger.Errorf("controller.GetPodByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ret, err := ctrl.WatchService.GetPodByLabel(c, para)
	if err != nil {
		util.Logger.Errorf("controller.GetPodByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetStatefulsetByName(c *gin.Context) {
	n := c.Param("name")
	ns := c.Param("namespace")
	ret, err := ctrl.WatchService.GetStatefulSetByName(c, ns, n)
	if err != nil {
		util.Logger.Errorf("controller.GetStatefulSetByName err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}
	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetStatefulsetByLabel(c *gin.Context) {
	para := &common.SelectorCondList{}
	if err := c.BindJSON(&para); err != nil {
		util.Logger.Errorf("controller.GetStatefulSetByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ret, err := ctrl.WatchService.GetStatefulSetByLabel(c, para)
	if err != nil {
		util.Logger.Errorf("controller.GetStatefulSetByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetServiceByName(c *gin.Context) {
	n := c.Param("name")
	ns := c.Param("namespace")
	ret, err := ctrl.WatchService.GetServiceByName(c, ns, n)
	if err != nil {
		util.Logger.Errorf("controller.GetServiceByName err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}
	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetServiceByLabel(c *gin.Context) {
	para := &common.SelectorCondList{}
	if err := c.BindJSON(&para); err != nil {
		util.Logger.Errorf("controller.GetServiceByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ret, err := ctrl.WatchService.GetServiceByLabel(c, para)
	if err != nil {
		util.Logger.Errorf("controller.GetServiceByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetIngressByName(c *gin.Context) {
	n := c.Param("name")
	ns := c.Param("namespace")
	ret, err := ctrl.WatchService.GetIngressByName(c, ns, n)
	if err != nil {
		util.Logger.Errorf("controller.GetIngressByName err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}
	ctrl.Jsonify(c, 200, ret, "")
}

func (ctrl *WatchController) GetIngressByLabel(c *gin.Context) {
	para := &common.SelectorCondList{}
	if err := c.BindJSON(&para); err != nil {
		util.Logger.Errorf("controller.GetIngressByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ret, err := ctrl.WatchService.GetIngressByLabel(c, para)
	if err != nil {
		util.Logger.Errorf("controller.GetIngressByLabel err: %s", err)
		ctrl.Jsonify(c, 400, nil, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "")
}
