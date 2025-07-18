package models

type User struct {
	ID       string `json:"id"  bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email"  bson:"email"`
	Password string `json:"-"  bson:"password"`
	Role     string `json:"role" bson:"role"`
}
