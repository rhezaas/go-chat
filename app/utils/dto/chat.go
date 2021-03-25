package dto

// Chat ...
type Chat struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	User  User   `json:"user"`
}
