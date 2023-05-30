package main

import (
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

func main() {
	router := gin.Default()
	router.Run("localhost:9090")
}
