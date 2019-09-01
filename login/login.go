package login

import (
	"github.com/gosagawa/go-error-handling-sample/db"
	"github.com/gosagawa/go-error-handling-sample/errors"
	"github.com/gosagawa/go-error-handling-sample/model"
)

// Validate ユーザをバリデート
func Validate(ID int) error {
	const op errors.Op = "login.Validate"
	user, err := db.LookUpUser(ID)
	if err != nil {
		return errors.E(op, err)
	}
	if !user.IsValid() {
		return errors.E(op, "User Invalid")
	}
	return nil
}

// GetUser ユーザを取得
func GetUser(ID int) (*model.User, error) {
	const op errors.Op = "login.GetUser"
	user, err := db.LookUpUser(ID)
	if err != nil {
		return nil, errors.E(op, err)
	}
	return user, nil
}
