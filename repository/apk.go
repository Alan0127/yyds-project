package repository

import "yyds-pro/trace"

type ApkRepo interface {
	FindApkById(ctx *trace.Trace, id int)
}
