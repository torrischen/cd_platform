package controller

import (
	"cd_platform/common"
	"cd_platform/ext"
	"cd_platform/pkg"
	"cd_platform/pkg/watch"
	"cd_platform/util"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	BaseController
	ExecService  pkg.ExecService
	WatchService watch.WatchService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		ExecService:  pkg.ExService,
		WatchService: watch.NewService(ext.MiddleWare),
	}
}

func (ctrl *ProjectController) InitProject(c *gin.Context) {
	var args common.InitProjectArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ext.MiddleWare.MysqlClient.CreateProject(args.Name); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ext.MiddleWare.HarborClient.CreateProject(args.Name); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.CreateNamespace(c, args.Name); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetProjectList(c *gin.Context) {
	var args common.GetProjectListQueryArgs
	if err := c.BindQuery(&args); err != nil {
		util.Logger.Errorf("controller.GetProjectList page or pagesize err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ret, err := ext.MiddleWare.MysqlClient.GetProjectList(args.Page, args.Pagesize)
	if err != nil {
		util.Logger.Errorf("controller.GetProjectList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetRepoList(c *gin.Context) {
	project := c.Param("project")
	repo := c.Param("repo")

	ret, err := ext.MiddleWare.HarborClient.GetRepoTag(project, repo)
	if err != nil {
		util.Logger.Errorf("controller.GetRepoList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) CreateApplication(c *gin.Context) {
	var args common.CreateProjectArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.CreateDeployment(c, args.Project, args.DeploymentRaw); err != nil {
		util.Logger.Errorf("controller.CreateApplication createDeployment err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.CreateService(c, args.Project, args.ServiceRaw); err != nil {
		util.Logger.Errorf("controller.CreateApplication createService err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.InsertIngressRule(c, args.Project, &args.IngressRule); err != nil {
		util.Logger.Errorf("controller.CreateApplication UpdateIngress err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, args.Project, "success")
}

func (ctrl *ProjectController) DestroyApplication(c *gin.Context) {
	var args common.DestroyProjectArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DestroyProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.DeleteIngressRule(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteIngressRule err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.DeleteService(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteService err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.DeleteDeployment(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteDeployment err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) DeployApplication(c *gin.Context) {
	var args common.DeployApplicationArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DeployApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.UpdateDeployment(c, args.Project, args.Application, args.Image); err != nil {
		util.Logger.Errorf("controller.DeployApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetApplicationDetails(c *gin.Context) {
	project := c.Param("project")

	ret, err := ctrl.WatchService.GetPodByProject(c, project)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationDetails err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}
