package controller

import "github.com/gin-gonic/gin"

func InitController() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	_ = NewBaseController()
	projectController := NewProjectController()
	sitController := NewSitController()

	pc := engine.Group("/api/dev/project")
	pc.POST("/init", projectController.InitProject)
	pc.GET("/list", projectController.GetProjectList)
	pc.GET("/:project/:repo/list", projectController.GetRepoList)
	pc.POST("/application/create", projectController.CreateApplication)
	pc.POST("/application/destroy", projectController.DestroyApplication)
	pc.POST("/application/deploy", projectController.DeployApplication)
	pc.GET("/:project/detail", projectController.GetApplicationDetails)
	pc.POST("/application/ingress/insert", projectController.InsertApplicationIngressPath)
	pc.GET("/:project/application/:application/ingress", projectController.GetApplicationIngress)
	pc.GET("/:project/pod/:podname/log", projectController.GetPodLog)

	sit := engine.Group("/api/dev/sit")
	sit.POST("/application/create", sitController.CreateSitApplication)
	sit.POST("/application/deploy", sitController.DeploySitApplication)
	sit.POST("/application/destroy", sitController.DestroySitApplication)
	sit.GET("/application/:application/detail", sitController.GetSitApplicationDetails)
	sit.POST("/application/ingress/insert", sitController.InsertSitApplicationIngressPath)
	sit.GET("/application/:application/ingress", sitController.GetSitApplicationIngress)
	sit.GET("/application/:application/log/:podname/log", sitController.GetSitPodLog)

	return engine
}
