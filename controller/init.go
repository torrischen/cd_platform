package controller

import "github.com/gin-gonic/gin"

func InitController() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	_ = NewBaseController()
	watchController := NewWatchController()
	execController := NewExecController()

	dep := engine.Group("/deployment")
	dep.GET("/get/:namespace/:name", watchController.GetDeploymentByName)
	dep.POST("/getByLabel", watchController.GetDeploymentByLabel)
	dep.POST("/create", execController.CreateDeployment)

	po := engine.Group("/pod")
	po.GET("/get/:namespace/:name", watchController.GetPodByName)
	po.POST("/getByLabel", watchController.GetPodByLabel)

	sts := engine.Group("/statefulset")
	sts.GET("/get/:namespace/:name", watchController.GetStatefulsetByName)
	sts.POST("/getByLabel", watchController.GetStatefulsetByLabel)

	svc := engine.Group("/service")
	svc.GET("/get/:namespace/:name", watchController.GetServiceByName)
	svc.POST("/getByLabel", watchController.GetServiceByLabel)

	igs := engine.Group("/ingress")
	igs.GET("/get/:namespace/:name", watchController.GetIngressByName)
	igs.POST("/getByLabel", watchController.GetIngressByLabel)

	return engine
}
