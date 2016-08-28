package services

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/almang45/go-sandbox/amdb/models"
	"github.com/almang45/go-sandbox/amdb/utils"
	"github.com/gin-gonic/gin"
)

var userDbmap = utils.InitUserDb()

func CreateUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	hashedPassword := hash(user.Password)
	user.Email = hashedPassword
	err := userDbmap.Insert(&user)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "Failed saving user!"})
	}
}

func GetUserByEmail(c *gin.Context) {
	email := c.Params.ByName("email")
	user := getUserByEmail(email)
	c.JSON(200, user)
}

func getUserByEmail(email string) models.User {
	var user models.User
	userDbmap.SelectOne(&user, "SELECT * FROM users WHERE email=$1", email)
	return user
}

func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	hashedPassword := hash(user.Password)
	storedUser := getUserByEmail(user.Email)
	if hashedPassword == storedUser.Password {
		c.JSON(200, storedUser)
	} else {
		c.JSON(404, gin.H{"error": "Invalid email or password!"})
	}
}

func hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
