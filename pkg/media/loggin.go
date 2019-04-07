package media

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

func (mw logginMiddleware) CreateContent(content *Content) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "createContent",
			"input", content,
			"output", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.CreateContent(content)
	return
}
func (mw logginMiddleware) FindContentByID(id uint) (output *Content, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "findContentByID",
			"input", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.FindContentByID(id)
	return
}
func (mw logginMiddleware) FindAllContents() (output []Content, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "findAllContents",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.FindAllContents()
	return
}
func (mw logginMiddleware) SetToLive(content *Content) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "setToLive",
			"input", content,
			"output", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.SetToLive(content)
	return
}
