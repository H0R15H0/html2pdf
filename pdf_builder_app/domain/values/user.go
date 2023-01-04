package values

import "github.com/google/uuid"

type UserID string

var ZeroUserID UserID

func (u UserID) String() string {
	return string(u)
}

func UserIDString(s string) (UserID, error) {
	userId, err := uuid.Parse(s)
	if err != nil {
		return ZeroUserID, err
	}
	return UserID(userId.String()), nil
}

func MustUserIDString(s string) UserID {
	userId, err := UserIDString(s)
	if err != nil {
		panic("UserID must be valid")
	}
	return userId
}
