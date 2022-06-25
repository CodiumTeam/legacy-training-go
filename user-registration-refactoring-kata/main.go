package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserPostResponse struct {
	Id    int
	Email string
	Name  string
}

type inMemoryUserRepository struct {
	users map[string]*User
}

func NewInMemoryUserRepository() inMemoryUserRepository {
	repository := inMemoryUserRepository{}
	repository.users = make(map[string]*User)
	return repository
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func (repository inMemoryUserRepository) find(email string) *User {
	return repository.users[email]
}

func (repository inMemoryUserRepository) save(user User) {
	repository.users[user.Email] = &user
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	repo := NewInMemoryUserRepository()
	r.POST("/users", func(c *gin.Context) {
		if len(c.Query("password")) < 8 || !strings.Contains(c.Query("password"), "_") {
			c.String(http.StatusBadRequest, "Password is not valid")
			return
		}
		if repo.find(c.Query("email")) != nil {
			c.String(http.StatusBadRequest, "The email is already in use")
			return
		}
		userId := rand.Int()
		repo.save(User{userId, c.Query("name"), c.Query("email"), c.Query("password")})
		response := UserPostResponse{
			Id:    userId,
			Email: c.Query("Email"),
			Name:  c.Query("name"),
		}
		c.JSON(http.StatusCreated, response)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
