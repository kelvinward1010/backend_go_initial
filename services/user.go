package services

import (
	"backend_go/core"
	"backend_go/models"
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	core.DB.Find(&users)
	utils.SendResponse(c, http.StatusOK, "Users retrieved successfully", users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := core.DB.First(&user, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	utils.SendResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := core.DB.First(&user, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	core.DB.Save(&user)
	utils.SendResponse(c, http.StatusOK, "User updated successfully", user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := core.DB.First(&user, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	core.DB.Delete(&user)
	utils.SendResponse(c, http.StatusOK, "User deleted successfully", nil)
}
