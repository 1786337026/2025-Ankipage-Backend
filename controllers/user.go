package controllers

import (
	"Ankipage/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var jwtKey = []byte("your_secret_key") // 替换为安全的密钥

// @Summary Register a new user
// @Description Create a new user account with a username and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body object{username=string,password=string,email=string} true "User credentials"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Failed to hash password or create user"
// @Router /register [post]
func RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"      binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//fmt.Println("haha")
		return
	}
	/*
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// 创建用户
		user := models.User{
			Username: input.Username,
			Password: string(hashedPassword),
			Email:    input.Email,
		}
		if err := db.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	*/
	user, err := services.Register(input.Username, input.Password, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "User registered successfully",
		"data":    user,
	})
}

// 登录用户
// @Summary Login a user
// @Description Authenticate a user with username and password, return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body object{username=string,password=string} true "User credentials"
// @Success 200 {object} web.Response4 map[string]interface{} "Logged in successfully, token generated"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 401 {object} map[string]interface{} "Invalid username or password"
// @Failure 500 {object} map[string]interface{} "Failed to generate token"
// @Router /login [post]
func LoginUser(c *gin.Context) {
	var input services.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := services.LoginService(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    1,
			"message": "username or email already exisits",
			"error":   err.Error(),
		})
		return
	}

	// 设置 Cookie
	c.SetCookie("token", token, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Logged in successfully",
		"user":    user,
	})
}
