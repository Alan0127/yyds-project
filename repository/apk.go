package repository

import (
	"yyds-pro/model"
	"yyds-pro/trace"
)

type ApkRepo interface {
	FindApkById(ctx *trace.Trace, id int) (model.App, error)
}
