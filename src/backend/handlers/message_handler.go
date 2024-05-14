package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Thwani47/react-go-webapp/types"
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context, db *sql.DB) {
	var message types.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		writeBadRequestError(c, err)
		return
	}

	result, err := db.Exec("INSERT INTO messages (channel_id, user_id, message) VALUES (?, ?, ?)", message.ChannelID, message.UserID, message.Text)

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

func ListMessages(c *gin.Context, db *sql.DB) {
	channelID, err := strconv.Atoi(c.Query("channelID"))

	if err != nil {
		writeBadRequestError(c, err)
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 100
	}

	lastMessageId, err := strconv.Atoi(c.Query("lastMessageId"))
	if err != nil {
		lastMessageId = 0
	}

	rows, err := db.Query("SELECT m.id, channel_id, user_id, u.username AS userName FROM messages m LEFT JOIN users u on u.id = m.user_id WHERE channel_id = ? AND m.id > ? ORDER  BY m.id ASC LIMIT ?", channelID, lastMessageId, limit)

	if err != nil {
		writeInternalServerError(c, "", err)
		return
	}

	var messages []types.Message

	for rows.Next() {
		var message types.Message

		err := rows.Scan(&message.ID, &message.ChannelID, &message.UserID, &message.Username, &message.Text)

		if err != nil {
			writeInternalServerError(c, "", err)
			return
		}

		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, messages)

}
