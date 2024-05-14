package types

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int `json:"id"`
	ChannelID int `json:"channel_id"`
	UserID    int `json:"userId"`
	Username  int `json:"userName"`
	Text      int `json:"text"`
}
