package dto

type User struct {
	ID   int64    `json:"UserID"`
	Name string   `json:"UserName"`
	Age  int      `json:"UserAge"`
	Tags []string `json:"UserTags"`
}
