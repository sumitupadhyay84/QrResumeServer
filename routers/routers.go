package routers

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/normos/qrresume/handlers"
)

func InitializeRoute(r *gin.Engine) {
	r.GET("/test", handlers.Test)
	r.POST("/createacc", handlers.CreateAccount)
	r.POST("/login", handlers.Login)
	r.GET("/resumes", handlers.GetResumes)
	r.GET("/templates", handlers.GetTemplates)
	r.GET("/template/:name", handlers.GetTemplate)
	r.GET("/create/:name", handlers.CreateResume)
	r.Use(static.Serve("/", static.LocalFile("./ui/build", true)))

}
