package report

type GenerateReport struct {
	Id int64 `json:"id" binding:"required"`
}

//启动参数
type StartFlag struct {
	Flag int32 `json:"flag"` //1:启动  2:不处理
}
