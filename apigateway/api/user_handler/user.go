package userhandler

import (
	"Api-GateWay/protos/userproto"
	"context"
	"net/http"
	"time"

	_ "Api-GateWay/docs"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ClientUser userproto.UserServiceClient
}

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @securityDefinitions.apikey Bearer
// @in 				header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host localhost:7777
// @BasePath /

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.UserRequest true "User request body"
// @Success 200 {object} userproto.UserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	var req userproto.UserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.CreateUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// VerfiyCode godoc
// @Summary VerifyCode a user
// @Description Login a user and get a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param login body userproto.LoginRequest true "Login request body"
// @Success 200 {object} userproto.Res
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /verifycode [post]
func (u *UserHandler) VerifyCode(c *gin.Context) {
	var req userproto.Req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.VerifyCode(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Login godoc
// @Summary Login a user
// @Description Login a user and get a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param login body userproto.LoginRequest true "Login request body"
// @Success 200 {object} userproto.UserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /login [post]
func (u *UserHandler) Login(c *gin.Context) {
	var req userproto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetByIdUser godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} userproto.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/{id} [get]
func (u *UserHandler) GetByIdUser(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.GetByIdUser(ctx, &userproto.GetUserRequest{UserId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} userproto.ListUser
// @Failure 500 {object} string
// @Security Bearer
// @Router /users [get]
func (u *UserHandler) GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := u.ClientUser.GetUsers(ctx, &userproto.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.Request true "User request body"
// @Success 200 {object} userproto.UpdateResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var req userproto.Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.UpdateUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePassword godoc
// @Summary Update a user's password
// @Description Update a user's password
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body userproto.UpdateUserRequest true "Password update request body"
// @Success 200 {object} userproto.UpdateResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/password [put]
func (u *UserHandler) UpdatePassword(c *gin.Context) {
	var req userproto.UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.UpdatePassword(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateEmail godoc
// @Summary Update a user's email
// @Description Update a user's email
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.UpdateEmailRequest true "Email update request body"
// @Success 200 {object} userproto.UpdateResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/email [put]
func (u *UserHandler) UpdateEmail(c *gin.Context) {
	var req userproto.UpdateEmailRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.UpdateEmail(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userproto.DeleteUserRequest true "Delete user request body"
// @Success 200 {object} userproto.UpdateResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users [delete]
func (u *UserHandler) DeleteUser(c *gin.Context) {

	var req userproto.DeleteUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := u.ClientUser.DeleteUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
