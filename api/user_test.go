package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"user_assignment/database"
)

type apiError struct {
	Error string
}

type ServiceTestSuite struct {
	suite.Suite
	router  *gin.Engine
	queries *database.Queries
}

// Create a normal test function and pass the suite to suite.Run
func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetUpSuite() {

	postgres := database.SetUpPostgres()
	suite.queries = database.New(postgres.DB)
	service := NewService(suite.queries)

	suite.router = gin.Default()
	service.RegisterHandlers(suite.router)
}

func (suite *ServiceTestSuite) SetUpTest() {
	err := suite.queries.TruncateUsers(context.Background())
	if err != nil {
		return
	}
}

// Test Create: success case
func (suite *ServiceTestSuite) TestCreate() {
	request := apiUser{
		Name: "andy",
	}

	var buffer bytes.Buffer
	suite.Require().NoError(json.NewEncoder(&buffer).Encode(request))

	// Create request
	req, err := http.NewRequest("POST", "/user", &buffer)
	suite.Require().NoError(err)

	// Record response
	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	// Compare status code
	suite.Require().Equal(http.StatusCreated, res.Result().StatusCode)

	// Decode RequestBody
	var result apiUser
	suite.Require().NoError(json.NewDecoder(res.Result().Body).Decode(&result))

	// Compare RequestBody
	suite.Require().Equal(request.Name, result.Name)
}

// Test Create: failed case - Bad Request
func (suite *ServiceTestSuite) TestCreateBadRequest() {
	// Create wrong RequestBody
	request := apiUser{
		Name: "the name exceeds twenty characters and failed the validation",
	}

	var buffer bytes.Buffer
	suite.Require().NoError(json.NewEncoder(&buffer).Encode(request))

	req, err := http.NewRequest("POST", "/user", &buffer)
	suite.Require().NoError(err)

	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	suite.Require().Equal(http.StatusBadRequest, res.Result().StatusCode)

	var result apiError
	suite.Require().NoError(json.NewDecoder(res.Result().Body).Decode(&result))
	suite.Require().Contains(result.Error, "max")
}

// Test Get: success case
func (suite *ServiceTestSuite) TestGet() {
	user, err := suite.queries.CreateUser(context.Background(), "testName")
	suite.Require().NoError(err)

	req, err := http.NewRequest("GET", fmt.Sprintf("/user/%d", user.ID), nil)
	suite.Require().NoError(err)

	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	suite.Require().Equal(http.StatusOK, res.Result().StatusCode)

	var result apiUser
	suite.Require().NoError(json.NewDecoder(res.Result().Body).Decode(&result))
	suite.Require().Equal(user.ID, result.ID)
	suite.Require().Equal(user.Name, result.Name)
}

// Test Get: failed case - Not Found
func (suite *ServiceTestSuite) TestGetNotFound() {
	req, err := http.NewRequest("GET", "/user/123", nil)
	suite.Require().NoError(err)

	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	suite.Require().Equal(http.StatusNotFound, res.Result().StatusCode)
}

// Test Get: failed case - Bad Request
func (suite *ServiceTestSuite) TestGetBadRequest() {
	req, err := http.NewRequest("GET", "/user/bad-request-test", nil)
	suite.Require().NoError(err)

	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	suite.Require().Equal(http.StatusBadRequest, res.Result().StatusCode)
}

// Test Update: success case
func (suite *ServiceTestSuite) TestUpdate() {
	user, err := suite.queries.CreateUser(context.Background(), "testOriginalName")

	suite.Require().NoError(err)

	update := apiUser{Name: "testUpdatedName"}

	var buffer bytes.Buffer
	suite.Require().NoError(json.NewEncoder(&buffer).Encode(update))

	req, err := http.NewRequest("PUT", fmt.Sprintf("/user/%d", user.ID), &buffer)
	suite.Require().NoError(err)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusOK, rec.Result().StatusCode)
	var result apiUser
	suite.Require().NoError(json.NewDecoder(rec.Result().Body).Decode(&result))
	suite.Require().Equal(user.ID, result.ID)
	suite.Require().Equal(update.Name, result.Name)
}

// Test Update: failed case - Bad Request
func (suite *ServiceTestSuite) TestUpdateBadRequest() {
	user, err := suite.queries.CreateUser(context.Background(), "testName")
	suite.Require().NoError(err)

	request := apiUser{
		Name: "the name of this user is too long that fails the validation",
	}
	var buffer bytes.Buffer
	suite.Require().NoError(json.NewEncoder(&buffer).Encode(request))

	req, err := http.NewRequest("PUT", fmt.Sprintf("/user/%d", user.ID), &buffer)
	suite.Require().NoError(err)

	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	suite.Require().Equal(http.StatusBadRequest, res.Result().StatusCode)
}

// Test Delete: success case
func (suite *ServiceTestSuite) TestDelete() {
	user, err := suite.queries.CreateUser(context.Background(), "testName")
	suite.Require().NoError(err)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user/%d", user.ID), nil)
	suite.Require().NoError(err)

	res := httptest.NewRecorder()
	suite.router.ServeHTTP(res, req)

	suite.Require().Equal(http.StatusOK, res.Result().StatusCode)
	_, err = suite.queries.GetUser(context.Background(), user.ID)
	suite.Require().Error(err)
}

// Test List: SUCCESS CASE
func (suite *ServiceTestSuite) TestList() {
	req, err := http.NewRequest("GET", "/user", nil)
	suite.Require().NoError(err)

	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, req)

	suite.Require().Equal(http.StatusNotFound, rec.Result().StatusCode)
}
