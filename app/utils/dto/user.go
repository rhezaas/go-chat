package dto

// User ...
type User struct {
	DefaultUserGroup
	InContact bool   `json:"in_contact"`
	Status    string `json:"status"`
}
