package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/db"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/helpers"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/models"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	database := db.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = database, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := database.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	database := db.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = database, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := database.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email / password",
		})

		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email / password",
		})

		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	database := db.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = database, contentType

	User := models.User{}
	NewUser := models.User{}

	id := c.Param("userId")

	err := database.First(&User, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&NewUser)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = database.Model(&User).Updates(NewUser).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"updated_at": User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	database := db.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = database, contentType
	User := models.User{}

	id := c.Param("userId")

	err := database.First(&User, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = database.Delete(&User).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}