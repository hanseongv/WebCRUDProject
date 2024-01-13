package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	for _, user := range users {
		//strconv.Itoa(int(user.ID)) == userID  다른 방법 찾기
		if strconv.Itoa(int(user.ID)) == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, "User not found")
}

// 양식
// {"name":"hs j","email":"hanseong@abcd.com"}
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	lastID++
	newUser.ID = lastID
	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": newUser})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	for i, user := range users {
		if strconv.Itoa(int(user.ID)) == userID {
			var updatedUser User
			if err := c.BindJSON(&updatedUser); err != nil {
				c.JSON(http.StatusBadRequest, "Invalid JSON format")
				return
			}
			updatedUser.ID = user.ID
			users[i] = updatedUser
			c.JSON(http.StatusOK, "User updated")
			return
		}
	}
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(int(user.ID)) == userID {
			//users[:i]는 0부터 i-1까지의 요소
			//users[i+1:]는 i+1부터 끝까지의 요소
			users = append(users[:i], users[i+1:]...)

			c.JSON(http.StatusOK, "User deleted")
			return
		}
	}

	c.JSON(http.StatusNotFound, "User not found")
}
