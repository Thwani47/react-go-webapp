package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Thwani47/react-go-webapp/types"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, db *sql.DB) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		writeBadRequestError(c, err)
		return
	}

	result, err := db.Exec("INSERT INTO uses (username, password) VALUES (?, ?)", user.Username, user.Password)

	if err != nil {
		writeInternalServerError(c, "", err)
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		writeInternalServerError(c, "", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
