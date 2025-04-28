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
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}
