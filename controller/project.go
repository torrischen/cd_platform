package controller

import (
	"bufio"
	"cd_platform/common"
	"cd_platform/ext"
	"cd_platform/pkg/workload"
	"cd_platform/pkg/workload/watch"
	"cd_platform/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type ProjectController struct {
	BaseController
	ExecService  workload.ExecService
	WatchService watch.WatchService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		ExecService:  workload.NewService(),
		WatchService: watch.NewService(),
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

	if err := ctrl.ExecService.CreateProjectIngress(c, args.Name); err != nil {
		util.Logger.Errorf("controller.InitProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetProjectList(c *gin.Context) {
	args := new(common.GetProjectListQueryArgs)
	if err := c.BindQuery(args); err != nil {
		util.Logger.Errorf("controller.GetProjectList page or pagesize err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ret, err := ext.MiddleWare.MysqlClient.GetProjectList(args.Page, args.PageSize)
	if err != nil {
		util.Logger.Errorf("controller.GetProjectList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetImageList(c *gin.Context) {
	project := c.Param("project")
	repo := c.Param("repo")

	ret, err := ext.MiddleWare.HarborClient.GetRepoTag(project, repo)
	if err != nil {
		util.Logger.Errorf("controller.GetImageList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetRepoList(c *gin.Context) {
	project := c.Param("project")

	ret, err := ext.MiddleWare.HarborClient.ListRepo(project)
	if err != nil {
		util.Logger.Errorf("controller.GetRepoList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) CreateApplication(c *gin.Context) {
	var args common.CreateApplicationArgs
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

	if err := ctrl.ExecService.CreateApplicationConfigmap(c, args.Project, args.DeploymentRaw.Name); err != nil {
		util.Logger.Errorf("controller.CreateApplication createConfigmap err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	//if err := ctrl.ExecService.InsertIngressRule(c, args.Project, &args.IngressRule); err != nil {
	//	util.Logger.Errorf("controller.CreateApplication UpdateIngress err: %s", err)
	//	ctrl.Jsonify(c, 400, struct{}{}, err.Error())
	//	return
	//}

	ctrl.Jsonify(c, 200, args.Project, "success")
}

func (ctrl *ProjectController) InsertApplicationIngressPath(c *gin.Context) {
	var args common.IngressRule
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.InsertApplicationIngressPath err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	if err := ctrl.ExecService.InsertIngressRule(c, &args); err != nil {
		util.Logger.Errorf("controller.InsertApplicationIngressPath err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) DestroyApplication(c *gin.Context) {
	var args common.DestroyApplicationArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DestroyProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
	}

	if err := ctrl.ExecService.DeleteIngressRule(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteIngressRule err: %s", err)
	}

	if err := ctrl.ExecService.DeleteService(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteService err: %s", err)
	}

	if err := ctrl.ExecService.DeleteDeployment(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteDeployment err: %s", err)
	}

	if err := ctrl.ExecService.DeleteConfigmap(c, args.Project, args.Application); err != nil {
		util.Logger.Errorf("controller.DestroyApplication DeleteDeployment err: %s", err)
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

	if err := ctrl.ExecService.UpdateDeploymentImage(c, args.Project, args.Application, args.Image); err != nil {
		util.Logger.Errorf("controller.DeployApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetApplicationList(c *gin.Context) {
	project := c.Param("project")

	ret, err := ctrl.WatchService.GetDeploymentListByProject(c, project)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetApplicationDetailsByApplication(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")

	ret, err := ctrl.WatchService.GetPodByApplication(c, project, application)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationDetailsByApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetApplicationDetailsByProject(c *gin.Context) {
	project := c.Param("project")

	ret, err := ctrl.WatchService.GetPodByProject(c, project)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationDetailsByProject err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetApplicationIngress(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")

	ret, err := ctrl.WatchService.GetIngressByApplication(c, project, application)
	if err != nil {
		util.Logger.Errorf("controller.GetSitApplicationIngress err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetPodLog(c *gin.Context) {
	project := c.Param("project")
	podname := c.Param("podname")

	stream, err := ctrl.WatchService.GetPodLog(c, project, podname)
	if err != nil {
		util.Logger.Errorf("controller.GetPodLog err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	reader := bufio.NewReader(stream)
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		util.Logger.Errorf("controller.GetPodLog newconn err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	go func() {
		for {
			msgType, _, err := conn.ReadMessage()
			if err != nil || msgType == -1 {
				stream.Close()
				conn.Close()
				return
			}
		}
	}()

	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() != "http2: response body closed" {
				util.Logger.Errorf("controller.GetPodLog readline err: %s", err)
			}
			return
		}

		err = conn.WriteMessage(1, data)
		if err != nil {
			util.Logger.Errorf("controller.GetPodLog getlog err: %s", err)
			return
		}
	}
}

func (ctrl *ProjectController) DeleteSpecifiedIngressRule(c *gin.Context) {
	var args common.DeleteSpecifiedIngressRuleArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DeleteSpecifiedIngressRule err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.DeleteSpecifiedIngressRule(c, args.Project, args.Path)
	if err != nil {
		util.Logger.Errorf("controller.DeleteSpecifiedIngressRule err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) PatchApplicationReplica(c *gin.Context) {
	var args common.PatchReplicaArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.PatchApplicationReplica err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.PatchDeploymentReplica(c, args.Project, args.Application, args.Replica)
	if err != nil {
		util.Logger.Errorf("controller.PatchApplicationReplica err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) RestartApplication(c *gin.Context) {
	var args common.RestartDeploymentArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.RestartApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.RestartDeployment(c, args.Project, args.Application)
	if err != nil {
		util.Logger.Errorf("controller.RestartApplication err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetApplicationYaml(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")

	ret, err := ctrl.WatchService.GetDeploymentYaml(c, project, application)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationYaml err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) SetApplicationEnvs(c *gin.Context) {
	var args common.SetEnvArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.SetApplicationEnvs err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.SetDeploymentEnv(c, args.Project, args.Application, args.Envs)
	if err != nil {
		util.Logger.Errorf("controller.SetApplicationEnvs err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetApplicationEnvs(c *gin.Context) {
	project := c.Param("project")
	appliation := c.Param("application")

	ret, err := ctrl.WatchService.GetDeploymentEnvs(c, project, appliation)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationEnvs err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) CreateApplicationConfigmap(c *gin.Context) {
	var args common.CreateConfigmapArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.CreateApplicationConfigmap err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.CreateApplicationConfigmap(c, args.Project, args.Application)
	if err != nil {
		util.Logger.Errorf("controller.CreateApplicationConfigmap err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) AddConfigToConfigmap(c *gin.Context) {
	var args common.AddConfigToConfigmapArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.AddConfigToConfigmap err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.AddConfigToConfigmap(c, args.Project, args.Application, args.Config)
	if err != nil {
		util.Logger.Errorf("controller.AddConfigToConfigmap err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) DeleteSpecifiedConfig(c *gin.Context) {
	var args common.DeleteSpecifiedConfigArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.DeleteSpecifiedConfig err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.DeleteSpecifiedConfig(c, args.Project, args.Application, args.ConfigName)
	if err != nil {
		util.Logger.Errorf("controller.DeleteSpecifiedConfig err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}

func (ctrl *ProjectController) GetApplicationConfigList(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")

	ret, err := ctrl.WatchService.GetApplicationConfigList(c, project, application)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationConfigList err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) GetApplicationConfigDetail(c *gin.Context) {
	project := c.Param("project")
	application := c.Param("application")
	cmname := c.Param("cmname")

	ret, err := ctrl.WatchService.GetApplicationConfigDetail(c, project, application, cmname)
	if err != nil {
		util.Logger.Errorf("controller.GetApplicationConfigDetail err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, ret, "success")
}

func (ctrl *ProjectController) UpdateSpecifiedConfig(c *gin.Context) {
	var args common.UpdateSpecifiedConfigArgs
	if err := c.BindJSON(&args); err != nil {
		util.Logger.Errorf("controller.UpdateSpecifiedConfig err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	err := ctrl.ExecService.UpdateSpecifiedConfig(c, args.Project, args.Application, args.ConfigName, args.NewVal)
	if err != nil {
		util.Logger.Errorf("controller.UpdateSpecifiedConfig err: %s", err)
		ctrl.Jsonify(c, 400, struct{}{}, err.Error())
		return
	}

	ctrl.Jsonify(c, 200, struct{}{}, "success")
}
