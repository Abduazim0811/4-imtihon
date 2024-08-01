package models

type User struct {
	ID       string `bson:"_id"`
	UserName string `bson:"user_name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Profile  struct {
		Name    string `bson:"name"`
		Address string `bson:"address"`
	} `bson:"profile"`
}
