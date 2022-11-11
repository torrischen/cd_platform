package controller

import (
	"cd_platform/common"
	"cd_platform/ext"
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

	if err := ext.MiddleWare.MysqlClient.CreateProject(args.Project); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ext.MiddleWare.HarborClient.CreateProject(args.Project); err != nil {
		ext.MiddleWare.K8sclient.ClientSet.CoreV1()
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

func (ctrl *ProjectController) GetProjectList(c *gin.Context) {
	ret, err := ext.MiddleWare.MysqlClient.GetProjectList()
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

func (ctrl *ProjectController) DeployApplication(c *gin.Context) {
	var args common.DeployProjectArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.CreateDeployment(c, args.Project, args.DeploymentRaw); err != nil {
		util.Logger.Errorf("controller.CreateProject createDeployment err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.CreateService(c, args.Project, args.ServiceRaw); err != nil {
		util.Logger.Errorf("controller.CreateProject createService err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.InsertIngressRule(c, &args.IngressRule); err != nil {
		util.Logger.Errorf("controller.CreateProject UpdateIngress err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, args.Project, "success")
}
