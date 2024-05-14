package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Thwani47/react-go-webapp/types"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, db *sql.DB) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		writeBadRequestError(c, err)
		return
	}

	row := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", user.Username, user.Password)

	var id int

	err := row.Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			writeInternalServerError(c, "invalid username or password", err)
			return
		}

		writeInternalServerError(c, "", err)
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
