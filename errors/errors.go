package errors

import "errors"

// Op オペレーション
type Op string

// Kind エラー種別
type Kind string

// Error エラー構造体
type Error struct {
	Op   Op    // operation
	Kind Kind  // category of errors
	Err  error // the wrapped error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// E エラーをラップする
func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case string:
			e.Err = errors.New(arg)
		case Kind:
			e.Kind = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}

// Ops スタックトレース取得
func Ops(e *Error) []Op {
	res := []Op{e.Op}
	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}
	res = append(res, Ops(subErr)...)
	return res
}
