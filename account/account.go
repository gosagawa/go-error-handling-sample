package account

import (
	"github.com/gosagawa/go-error-handling-sample/errors"
	"github.com/gosagawa/go-error-handling-sample/login"
	"github.com/gosagawa/go-error-handling-sample/model"
)

// GetUser ユーザを取得
func GetUser(ID int) (*model.User, error) {
	const op errors.Op = "account.GetUser"
	err := login.Validate(ID)
	if err != nil {
		return nil, errors.E(op, err)
	}

	user, err := login.GetUser(ID)
	if err != nil {
		return nil, errors.E(op, err)
	}
	return user, nil
}
