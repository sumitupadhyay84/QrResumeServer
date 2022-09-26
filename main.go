package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/normos/qrresume/database"
	"github.com/normos/qrresume/models"
	"github.com/normos/qrresume/routers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	routers.InitializeRoute(router)
	db := database.CreateDbConn()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Resume{})
	db.AutoMigrate(&models.Template{})
	//db.Create(&models.Template{
	//	Name:  "elegant",
	//	Tabs:  "{\"tabs\":[{\"name\":\"Basic Info\",\"fields\":[{\"name\":\"Name\"},{\"name\":\"Ph Number\"}]},{\"name\":\"Skills\",\"fields\":[{\"name\":\"Skill List\"}]},{\"name\":\"Education\",\"fields\":[{\"name\":\"Passout Year\"},{\"name\":\"Institute Name\"},{\"name\":\"Achievements\"}]},{\"name\":\"Experience\",\"fields\":[{\"name\":\"Leaving Year\"},{\"name\":\"Company Name\"},{\"name\":\"Achievements\"}]}]}",
	//	Struc: "{\"name\":\"\",\"phNumber\":\"\",\"passoutYear\":\"\",\"instituteName\":\"\",\"instituteAchivements\":\"\",\"leavingYear\":\"\",\"companyName\":\"\",\"companyAchivements\":\"\",\"skillList\":\"\"}",
	//})
	router.Run(":8080")

}
