package db

import (
	"github.com/gosagawa/go-error-handling-sample/model"
)

//LookUpUser DBからデータ取得
func LookUpUser(ID int) (*model.User, error) {

	//DBのコネクション

	//データの取得

	return &model.User{
		ID:       1,
		Name:     "Go Sagawa",
		IsActive: true,
	}, nil

}
