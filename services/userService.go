package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramezanpour/users/dao"
	"github.com/ramezanpour/users/models"
	"github.com/ramezanpour/users/utilities/database"
	"github.com/ramezanpour/users/utilities/resources"
	"github.com/ramezanpour/users/utilities/security"
)

// UserService contains APIs for user management
type UserService struct{}

// Serve initialize APIs and routes
func (us UserService) Serve(engine *gin.Engine) {
	engine.POST("/users/login", us.login)
	engine.POST("/users/signup", us.signup)
	engine.GET("/users/token", us.getByToken)
	engine.GET("/users/get/:id", us.getUserbyID)
	engine.GET("/users", us.getAll)
	database.Db.AutoMigrate(&models.User{})
}

func (us UserService) login(c *gin.Context) {
	var input dao.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.InvalidRequest)
		return
	}

	var user models.User
	if err := database.Db.First(&user, "username = ?", input.Username).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, resources.InvalidCredintials)
		return
	}

	if user.Password != security.HashString(input.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, resources.InvalidCredintials)
		return
	}

	tokenString, err := security.CreateToken(&user)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusOK, getUserViewModel(&user, tokenString))
}

func (us UserService) signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.InvalidRequest)
		return
	}

	if len(user.Username) < 3 {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.UsernameLengthCannotBeLessThan3Chars)
		return
	}

	if len(user.Password) < 6 {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.PasswordLengthCannotBeLessThan6Chars)
		return
	}

	var dbUser models.User
	if !database.Db.First(&dbUser, "username = ?", user.Username).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusConflict, resources.UsernameAlreadyExists)
		return
	}

	// Hash the password before saving to the database
	user.Password = security.HashString(user.Password)

	if err := database.Db.Create(&user).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, getUserViewModel(&user, ""))
}

func getUserViewModel(user *models.User, token string) *dao.UserViewModel {
	result := &dao.UserViewModel{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

	return result
}

func (us UserService) getByToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("x-token")
	claims, err := security.ParseToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.CannotParseToken)
		return
	}

	userID := claims.Id
	if id, _ := strconv.Atoi(userID); id != 0 {
		var user models.User
		if err := database.Db.First(&user, userID).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, getUserViewModel(&user, ""))
	}

}

func (us UserService) getUserbyID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.InvalidRequest)
		return
	}

	var user models.User
	if database.Db.First(&user, userID).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusNotFound, resources.NotFound)
		return
	}

	c.JSON(http.StatusOK, getUserViewModel(&user, ""))
}

func (us UserService) getAll(c *gin.Context) {
	const pageSize = 10
	pageString := c.Query("page")
	if pageString == "" {
		pageString = "1"
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resources.InvalidPage)
		return
	}

	var users []models.User
	if err := database.Db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &users)
}
