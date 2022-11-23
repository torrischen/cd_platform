package controller

import (
	"cd_platform/common"
	"cd_platform/pkg/watch"
	"cd_platform/pkg/workload/sit"
	"cd_platform/util"
	"github.com/gin-gonic/gin"
)

type SitController struct {
	BaseController
	SitService   sit.SitService
	WatchService watch.WatchService
}

func NewSitController() *SitController {
	return &SitController{
		SitService:   sit.NewService(),
		WatchService: watch.NewService(),
	}
}

func (ctrl *SitController) CreateSitApplication(c *gin.Context) {
	var args common.CreateSitApplicationArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateSitApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.CreateSitNamespace(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.CreateSitApplication CreateNS err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.CreateSitDeployment(c, args.Project, args.DeploymentRaw); err != nil {
		util.Logger.Errorf("controller.CreateSitApplication CreateDeployment err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.CreateSitService(c, args.Project, args.ServiceRaw); err != nil {
		util.Logger.Errorf("controller.CreateSitApplication CreateService err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	//if err := ctrl.SitService.InsertSitIngressRule(c, args.Application, &args.IngressRule); err != nil {
	//	util.Logger.Errorf("controller.CreateSitApplication InsertIngressRule err: %s", err)
	//	ctrl.Jsonify(c, 400, struct{}{}, err.Error())
	//	return
	//}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *SitController) InsertSitApplicationIngressPath(c *gin.Context) {
	var args common.IngressRule
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.InsertApplicationIngressPath err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.InsertSitIngressRule(c, &args); err != nil {
		util.Logger.Errorf("controller.InsertApplicationIngressPath err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *SitController) DeploySitApplication(c *gin.Context) {
	var args common.DeploySitApplicationArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DeploySitApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.UpdateSitDeployment(c, args.Project, args.Application, args.Image); err != nil {
		util.Logger.Errorf("controller.DeploySitApplication UpdateSitDeployment err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *SitController) DestroySitApplication(c *gin.Context) {
	var args common.DestroySitApplicationArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DestroySitApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.DeleteSitIngressRule(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroySitApplication DeleteSitIngressRule err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.DeleteSitService(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroySitApplication DeleteSitService err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.DeleteSitDeployment(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroySitApplication DeleteSitDeployment err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.SitService.DeleteSitNamespace(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroySitApplication DeleteSitNS err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *SitController) GetSitApplicationDetails(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")
	ret, err := ctrl.WatchService.GetSitPodByApplication(c, project, application)
	if err != nil {
		util.Logger.Errorf("controller.GetSitApplicationDetails err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *SitController) GetSitApplicationIngress(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")

	ret, err := ctrl.WatchService.GetSitIngressByApplication(c, project, application)
	if err != nil {
		util.Logger.Errorf("controller.GetSitApplicationIngress err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *SitController) GetSitPodLog(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")
	podname := c.Param("podname")

	ret, err := ctrl.WatchService.GetSitPodLog(c, project, application, podname)
	if err != nil {
		util.Logger.Errorf("controller.GetSitPodLog err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, util.ByteToString(ret), "success")
}

func (ctrl *SitController) DeleteSpecifiedSitIngressRule(c *gin.Context) {
	var args common.DeleteSpecifiedIngressRuleArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DeleteSpecifiedSitIngressRule err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.SitService.DeleteSpecifiedSitIngressRule(c, args.Path)
	if err != nil {
		util.Logger.Errorf("controller.DeleteSpecifiedSitIngressRule err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}
