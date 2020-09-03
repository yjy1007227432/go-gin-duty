package app

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logs.Info(err.Key, err.Message)
	}

	return
}
