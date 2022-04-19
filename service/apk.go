package service

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

type ApkService interface {
	GetApkById(ctx *trace.Trace, id int) (model.App, error)
}
