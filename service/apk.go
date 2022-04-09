package service

import "yyds-pro/trace"

type ApkService interface {
	GetApkById(ctx *trace.Trace, id int)
}
