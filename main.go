package main

import (
	"errors"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id      string `json: "id"`
	Name    string `json: "name"`
	Country string `json: "country"`
	Age     int64  `json: "age"`
}

var users = []user{
	{Id: "1", Name: "Joe", Country: "United States", Age: 25},
	{Id: "2", Name: "Maria", Country: "Brasil", Age: 22},
	{Id: "3", Name: "Akemi", Country: "Japan", Age: 45},
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func addUser(context *gin.Context) {
	var newUser user
	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUser(context *gin.Context) {
	id := context.Param("id")
	user, err := getUserById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

// func updateUser(context *gin.Context) {
// 	id := context.Param("id")
// 	userUp, err := getUserById(id)

// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not update"})
// 		return
// 	}
// 	context.IndentedJSON(http.StatusOK, userUp)
// }

func updateUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var u user
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}
	log.Println(u.Name)
}

func getUserById(id string) (*user, error) {
	for i, u := range users {
		if u.Id == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", addUser)
	http.HandleFunc("/users", updateUser)
	// router.POST("/users/:id", updateUser)

	router.Run("localhost:9090")
}
