package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/normos/qrresume/database"
	"github.com/normos/qrresume/models"
	"golang.org/x/crypto/bcrypt"
)

var t http.Cookie

type head struct {
	Token string `header:Token`
}

type Claims struct {
	Email      string `json:"email"`
	Authorized string `json:"authorized"`
	Exp        string `json:"exp"`
}

func generatehashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte("secretkeyqrresume")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "testing the apis",
	})
}

func CreateAccount(c *gin.Context) {
	body := models.Usercreateac{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db := database.CreateDbConn()
	err := db.Create(&models.User{Name: body.Name, EmailId: body.EmailId, Password: generatehashPassword(body.Password)}).Error
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"code": 201})
		return
	}
	fmt.Println(body)
	c.JSON(http.StatusAccepted, gin.H{"code": 200})
}

func Login(c *gin.Context) {
	body := models.Login{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(body)
	db := database.CreateDbConn()

	var authuser models.User
	err := db.First(&authuser, "email_id = ?", body.EmailId).Error
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"code": 202})
		return
	}

	check := CheckPasswordHash(body.Password, authuser.Password)

	if !check {
		c.JSON(http.StatusAccepted, gin.H{"code": 202})
		return
	}

	validToken, err := GenerateJWT(authuser.EmailId)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"code": 203})
		return
	}
	expirationTime := time.Now().Add(time.Hour * 20)
	fmt.Println(body)
	//cookie, _ := c.Cookie(validToken)
	//c.SetCookie("token", validToken, 60*60*20, "", "localhost", true, true)

	t = http.Cookie{Name: "token", Value: validToken, Expires: expirationTime, HttpOnly: true}
	http.SetCookie(c.Writer, &t)
	c.JSON(http.StatusAccepted, gin.H{"code": 200})
}

func GetResumes(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"code": 204})
	}
	fmt.Printf("Cookie value: %v", cookie.Value)

	hmacSecret := []byte("secretkeyqrresume")
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	db := database.CreateDbConn()
	var users []models.User
	//db.Where("email_id = ?", claims["email"]).First(&users)
	db.Model(&models.User{}).Preload("Resumes").Where("email_id = ?", claims["email"]).Find(&users)
	c.JSON(http.StatusAccepted, gin.H{"code": 200, "resumes": users[0].Resumes})
	//var user models.User
	//var resumes []models.Resume
	//db := database.CreateDbConn()
	//db.Where("email = ?", email).First(&user)
	//db.Model(&models.User{}).Preload("Resumes").Find(&resumes)

}

func DeleteResume(c *gin.Context) {
	if err := c.ShouldBindHeader(&head{}); err != nil {
		c.JSON(http.StatusOK, err)
	}

}

func EditResume(c *gin.Context) {
	if err := c.ShouldBindHeader(&head{}); err != nil {
		c.JSON(http.StatusOK, err)
	}

}

func GenerateResumePdf(c *gin.Context) {
	if err := c.ShouldBindHeader(&head{}); err != nil {
		c.JSON(http.StatusOK, err)
	}

}

func GenerateResumeQr(c *gin.Context) {
	if err := c.ShouldBindHeader(&head{}); err != nil {
		c.JSON(http.StatusOK, err)
	}

}

func GetTemplates(c *gin.Context) {
	db := database.CreateDbConn()
	var templates []models.Template
	db.Select("id", "name", "url").Find(&templates)
	c.JSON(http.StatusOK, gin.H{"code": 200, "templates": templates})
}

func GetTemplate(c *gin.Context) {
	name := c.Param("name")
	db := database.CreateDbConn()
	var template models.Template
	db.Where("name = ?", name).First(&template)
	var tabsMap map[string]interface{}
	json.Unmarshal([]byte(template.Tabs), &tabsMap)
	var strucMap map[string]interface{}
	json.Unmarshal([]byte(template.Tabs), &strucMap)
	c.JSON(http.StatusOK, gin.H{"code": 200, "tabs": tabsMap, "struc": strucMap})
}
