package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/najibjodiansyah/gin-users-api/models"
	"github.com/najibjodiansyah/gin-users-api/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) Get(ctx *gin.Context) {
	users, err := uc.UserService.Get()
	if users == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no users found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}

func (uc *UserController) GetByUsername(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetByUser(username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (uc *UserController) Create(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = uc.UserService.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Update(ctx *gin.Context) {
	var user models.User
	name := ctx.Param("name")
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = uc.UserService.Update(name, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	err := uc.UserService.Delete(name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	userRoute := router.Group("/user")
	userRoute.GET("/", uc.Get)
	userRoute.GET("/:name", uc.GetByUsername)
	userRoute.POST("/", uc.Create)
	userRoute.PUT("/:name", uc.Update)
	userRoute.DELETE("/:name", uc.Delete)
}
