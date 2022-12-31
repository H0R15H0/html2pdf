package values

import "github.com/google/uuid"

type UserID uuid.UUID

var ZeroUserID UserID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

func UserIDString(s string) (UserID, error) {
	userId, err := uuid.Parse(s)
	if err != nil {
		return ZeroUserID, err
	}
	return UserID(userId), nil
}

func MustUserIDString(s string) UserID {
	userId, err := UserIDString(s)
	if err != nil {
		panic("UserID must be valid")
	}
	return userId
}
