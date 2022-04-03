package log

import "path/filepath"

type Options struct {
	DevFlag   bool
	FileDir   string
	MaxSize   int //文件切分
	MaxBackUp int //文件保留个数
	MaxAge    int
	Level     string
	CtxKey    string
}

type ZapLogOption func(option *Options)

func NewOption(opts ...ZapLogOption) *Options {
	opt := &Options{
		DevFlag:   true,
		MaxSize:   100,
		MaxBackUp: 60,
		MaxAge:    30,
		CtxKey:    "zapLog_key",
	}
	opt.FileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
	opt.FileDir += "/logs/"
	for _, f := range opts {
		f(opt)
	}
	return opt
}

func SetLevel(level string) ZapLogOption {
	return func(option *Options) {
		option.Level = level
	}
}
