/*
 *     @Copyright
 *     @file: Request.go
 *     @author: Equationzhao
 *     @email: equationzhao@foxmail.com
 *     @time: 2023/3/17 下午9:54
 *     @last modified: 2023/3/17 下午8:07
 *
 *
 *
 */

package DDNS

import (
	"errors"
	"fmt"

	"github.com/robfig/cron/v3"
)

const (
	Success = iota
	NotExecute
	Failed
)

type Request interface {
	cron.Job
	ToParameters() Parameters
	GetName() string    // return like "dnspod"
	MakeRequest() error // MakeRequest will return error if exist
	Status() Status
}

type Status struct {
	Name    string
	Msg     string
	Success int
}

//const Success = "success"
//const Failed = "failed"

// AppendMsg
// append msg to Status.Msg, using fmt.Sprint
func (s *Status) AppendMsg(msg ...any) *Status {
	s.Msg += fmt.Sprint(msg...)
	return s
}

// AppendMsgF
// append msg to Status.Msg, using fmt.Sprintf
func (s *Status) AppendMsgF(format string, msg ...any) *Status {
	s.Msg += fmt.Sprintf(format, msg...)
	return s
}

func ExecuteRequest(request Request) error {
	return request.MakeRequest()
}

func ExecuteRequestList(request ...Request) ([]Status, error) {
	var content []Status
	var err error
	for _, r := range request {
		errTemp := r.MakeRequest()
		content = append(content, r.Status())
		if errTemp != nil {
			err = errors.Join(err, errTemp)
		}
	}
	return content, err
}
