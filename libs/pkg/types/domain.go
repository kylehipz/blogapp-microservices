package types

type Blog struct {
	ID        string `json:"id"`
	Author    *User  `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	CreatedAt string `json:"createdAt"`
}

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}
