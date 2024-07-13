package handler

import (
	"fmt"
	"net/http"
	pb "user_service/genproto/authentication"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register user
// @Description create new users
// @Tags auth
// @Param info body authentication.UserDetails true "User info"
// @Success 200 {object} authentication.ID
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /user/register [post]
func (h Handler) Register(c *gin.Context) {
	h.Log.Info("Register is starting")
	req := pb.UserDetails{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.Auth.Register(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	h.Log.Info("Register ended")
	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Login user
// @Description it generates new acces and refresh tokens
// @Tags auth
// @Param userinfo body authentication.LoginRequest true "user name and password"
// @Success 200 {object} authentication.LoginResponse "accestoken and refreshtoken"
// @Failure 400 {object} string "Invalid date"
// @Failure 500 {object} string "error while reading from server"
// @Router /user/login [post]
func (h Handler) Login(c *gin.Context) {
	h.Log.Info("Login is working")
	req := pb.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
	}
	res, err := h.Auth.Login(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  res.Access.Accesstoken,
		"refresh_token": res.Refresh.Refreshtoken,
	})
	h.Log.Info("login is succesfully ended")

}

// RefreshToken godoc
// @Summary Check refresh token
// @Description Checks refresh token. If valid, it returns a new access token.
// @Tags auth
// @Param refreshToken body authentication.CheckRefreshTokenRequest true "Refresh Token"
// @Success 200 {object} string "accessToken"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Server error"
// @Router /user/refresh [post]
func (h Handler) RefreshToken(c *gin.Context) {
	h.Log.Info("refreshtoken is working")
	var data map[string]string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	refreshTokenStr := data["token"]
	s := fmt.Sprintf("Received username: %s", refreshTokenStr)
	h.Log.Info(s)

	res, err := h.Auth.CheckRefreshToken(c, &pb.CheckRefreshTokenRequest{Token: refreshTokenStr})
	if err != nil {
		h.Log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if !res.Acces {
		h.Log.Error("Refresh token is invalid")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid"})
	}

	c.JSON(http.StatusOK, gin.H{
		"accestoken": res.Accestoken,
	})
	h.Log.Info("refreshtoken is succesful")
}
