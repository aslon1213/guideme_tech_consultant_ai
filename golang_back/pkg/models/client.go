package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role,omitempty" bson:"role"`
}

type Client struct {
	User         `bson:",inline"`
	Documents    []Document    `json:"documents,omitempty" bson:"documents"`
	Actionslist  []Action      `json:"actionslist,omitempty" bson:"actionslist"`
	ApiKeys      []string      `json:"api_keys,omitempty" bson:"api_keys"`
	Balance      float64       `json:"balance,omitempty" bson:"balance"`
	Transactions []Transaction `json:"transactions,omitempty" bson:"transactions"`
	UsageInfo    UsageInfo     `json:"usage_info,omitempty" bson:"usage_info"`
}

type UsageInfo struct {
	TotalTokenUsage float64 `json:"total_usage" bson:"total_usage"`
	TotalRequests   int     `json:"total_requests" bson:"total_requests"`
}

type Transaction struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Positive bool    `json:"positive" bson:"positive"`
}
