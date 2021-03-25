package models

import (
	"errors"

	"go-chat/app/utils/dto"
	"go-chat/app/utils/helper"
	"go-chat/app/utils/interfaces"
)

// UserModel ...
type UserModel struct {
	Redis interfaces.Redis
}

// CreateUser ...
func (User UserModel) CreateUser(user dto.User) error {
	userKey := helper.KeyBuilder(`user`, user.ID)

	_user, err := User.Redis.HGetAll(userKey)

	if len(_user) > 0 || _user[`ID`] == user.ID {
		err = errors.New(`User Already Exists`)
	} else {
		err = User.Redis.HSet(userKey, map[string]string{
			`id`:   user.ID,
			`name`: user.Name,
		})
	}

	return err
}

// InsertGlobalContact ...
func (User UserModel) InsertGlobalContact(userID string) error {
	contactIDKey := helper.KeyBuilder("user", userID)
	contactKey := helper.KeyBuilder(`contacts`)

	return User.Redis.SAdd(contactKey, contactIDKey)
}

// InsertContact ...
func (User UserModel) InsertContact(userID string, contactID string) error {
	contactIDKey := helper.KeyBuilder("user", contactID)
	contactListKey := helper.KeyBuilder("user", userID, "contacts")

	return User.Redis.SAdd(contactListKey, contactIDKey)
}

// ListContact ...
func (User UserModel) ListContact(userID string) ([]dto.User, error) {
	contacts := []dto.User{}

	userKey := helper.KeyBuilder("user", userID)
	contactListKey := helper.KeyBuilder(userKey, "contacts")

	queryContactIDList, err := User.Redis.SMembers(contactListKey)
	for _, contactID := range queryContactIDList {
		contact, err := User.Redis.HGetAll(contactID)

		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, dto.User{
			DefaultUserGroup: dto.DefaultUserGroup{
				ID:   contact["id"],
				Name: contact["name"],
			},
		})
	}

	return contacts, err
}
