package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	http "net/http"
	database "user_assignment/database"
	"user_assignment/model"
)

type Service struct {
	queries *database.Queries
}

func NewService(queries *database.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/user", s.List)
	router.GET("/user/:user_id", s.Get)
	router.POST("/user", s.Create)
	router.PUT("/user/:user_id", s.Update)
	router.DELETE("/user/:user_id", s.Delete)
}

// error message
const idNotFoundMsg = "The inputted id does not exist"
const wrongIdFormatMsg = "Please provide id in the correct format"
const wrongNameFormatMsg = "Please provide name within 20 alphabetical characters"
const internalServerErrorMsg = "Please contact Madelain for further assistance"

// apiUser api input
type apiUser struct {
	ID   int32  `json:"user_id,numeric"`
	Name string `json:"name" binding:"required,max=20,alpha"`
}

type pathParameter struct {
	ID int32 `uri:"user_id" binding:"required,numeric"`
}

// fromDB result from DE
func fromDB(user model.User) *apiUser {
	return &apiUser{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (s *Service) Create(c *gin.Context) {
	// Parse request
	var request apiUser
	// Abort if the inputted name is in the wrong format
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": wrongNameFormatMsg, "error": err.Error()})
		return
	}
	// Create user to DB
	param := request.Name
	user, err := s.queries.CreateUser(context.Background(), param)
	// Abort should any server errors
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMsg, "error": err.Error()})
		return
	}
	// Build success response
	c.JSON(http.StatusCreated, fromDB(user))
}

func (s *Service) Get(c *gin.Context) {
	// Parse request
	var pathParam pathParameter
	// Abort if the inputted id is in the wrong format
	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": wrongIdFormatMsg, "error": err.Error()})
		return
	}
	// Get user from DB
	user, err := s.queries.GetUser(context.Background(), pathParam.ID)
	if err != nil {
		// Abort if the inputted id is not available
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": idNotFoundMsg, "error": err.Error()})
			return
		}
		// Abort should any server errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMsg, "error": err.Error()})
		return
	}
	// Build success response
	c.JSON(http.StatusOK, fromDB(user))
}

func (s *Service) Update(c *gin.Context) {
	// Parse request
	var pathParam pathParameter
	// Abort if the inputted id is in the wrong format
	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": wrongIdFormatMsg, "error": err.Error()})
		return
	}
	var request apiUser
	// Abort if the inputted name is in the wrong format
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": wrongNameFormatMsg, "error": err.Error()})
		return
	}
	// Update user in DB
	params := database.UpdateUserParams{
		ID:   pathParam.ID,
		Name: request.Name,
	}
	user, err := s.queries.UpdateUser(context.Background(), params)
	if err != nil {
		// Abort if the inputted id is not available
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": idNotFoundMsg, "error": err.Error()})
			return
		}
		// Abort should any server errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMsg, "error": err.Error()})
		return
	}
	// Build success response
	c.JSON(http.StatusOK, fromDB(user))
}

func (s *Service) Delete(c *gin.Context) {
	// Parse request
	var pathParam pathParameter
	// Abort if the inputted id is in the wrong format
	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": wrongIdFormatMsg, "error": err.Error()})
		return
	}
	// Delete user in DB
	user, err := s.queries.DeleteUser(context.Background(), pathParam.ID)
	if err != nil {
		// Abort if the inputted id is not available
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": idNotFoundMsg, "error": err.Error()})
			return
		}
		// Abort should any server errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMsg, "error": err.Error()})
		return
	}
	// Build success response
	result := fromDB(user)
	if result != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	}
}

func (s *Service) List(c *gin.Context) {
	// List users
	users, err := s.queries.GetUserList(context.Background())
	// Abort should any server errors
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": internalServerErrorMsg, "error": err.Error()})
		return
	}
	// Abort if there is none user info available
	if len(users) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Users information unavailable"})
		return
	}
	// Build success response
	var response []*apiUser
	for _, user := range users {
		response = append(response, fromDB(user))
	}
	c.JSON(http.StatusOK, users)
}
