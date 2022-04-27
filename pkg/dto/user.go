package dto

type User struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Group Group    `json:"group"`
	Tags  []string `json:"tags"`
}
