package auth

import (
	"time"

	"github.com/go-kit/kit/log"
)

type logginMiddleware struct {
	logger log.Logger
	next   Service
}

//NewLogginMiddleware Constructor
func NewLogginMiddleware(logger log.Logger, next Service) Service {
	return &logginMiddleware{
		logger: logger,
		next:   next,
	}
}

func (mw logginMiddleware) RegisterUser(user *User) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "register",
			"input", user,
			"output", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.RegisterUser(user)
	return
}

func (mw logginMiddleware) Login(email string, password string) (output *User, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "register",
			"input_mail", email,
			"input_pass", password,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Login(email, password)
	return
}
