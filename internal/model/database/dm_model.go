package database

import "time"

type Conversation struct {
	Id        int
	User1Id   int
	User2Id   int
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DMessage struct {
	Id             int
	ConversationId int
	SenderId       int
	Content        string
	CreatedAt      time.Time
}
