package handlers

import (
	"bytes"
	"encoding/json"
	"http-demo/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type MockUserDB struct{}

type IUserDB interface {
	Create(user *models.User) (*models.User, error)
	GetById(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	DeleteById(id int) (int, error)
}

func (mdb *MockUserDB) Create(user *models.User) (*models.User, error) {
	return &models.User{Id: 1, Name: "jiten", Email: "Jitenp@outlook.com"}, nil
}

func (mdb *MockUserDB) GetById(id int) (*models.User, error) {
	return &models.User{Id: 1, Name: "jiten", Email: "Jitenp@outlook.com"}, nil
}

func (mdb *MockUserDB) GetAll() ([]models.User, error) {
	return []models.User{models.User{Id: 1, Name: "jiten", Email: "Jitenp@outlook.com"}}, nil
}

func (mdb *MockUserDB) DeleteById(id int) (int, error) {
	return 1, nil
}

func setUpUserRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	ch := NewUserHandler(&MockUserDB{})
	user_router := r.Group("/v1/public")
	user_router.POST("/users", ch.Insert)
	user_router.GET("/users/:id", ch.GetUserBy)
	user_router.GET("/users/all", ch.GetAllUsers)
	user_router.DELETE("/users/:id", ch.GetDeleteBy)
	return r
}

func TestCreateUser(t *testing.T) {

	r := setUpUserRouter()

	user := &models.User{Id: 1, Name: "jiten", Email: "Jitenp@outlook.com"}

	buff, err := json.Marshal(user)

	require.NoError(t, err)

	rec := httptest.NewRecorder()

	t.Log(string(buff))

	req := httptest.NewRequest(http.MethodPost, "/v1/public/users", bytes.NewReader(buff))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(rec, req)
	require.Equal(t, 201, rec.Code)
}
