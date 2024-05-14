package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Thwani47/react-go-webapp/types"
	"github.com/gin-gonic/gin"
)

func CreateChannel(c *gin.Context, db *sql.DB) {

	fmt.Println("called...")
	var channel types.Channel
	if err := c.ShouldBindJSON(&channel); err != nil {
		writeBadRequestError(c, err)
		return
	}

	result, err := db.Exec("INSERT INTO channels (name) VALUES (?)", channel.Name)

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

func ListChannels(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM channels")

	if err != nil {
		writeInternalServerError(c, "", err)
		return
	}

	var channels []types.Channel

	fmt.Println("ROws", rows)

	for rows.Next() {
		var channel types.Channel

		err := rows.Scan(&channel.ID, &channel.Name)

		if err != nil {
			writeInternalServerError(c, "", err)
			return
		}

		channels = append(channels, channel)
	}

	c.JSON(http.StatusOK, channels)
}
