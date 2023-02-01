package report

type ResReportData struct {
	Data []GenerateReportResponse `json:"data"`
}

type GenerateReportResponse struct {
	Id          int64  `json:"id"`          // id
	Name        string `json:"name"`        //name
	Blog        string `json:"blog"`        //个人网址
	PhoneNumber string `json:"phoneNumber"` //电话号码
}

//kafka发送来的数据
type ResReportKafkaData struct {
	Data []Company `json:"data"`
}

type Company struct {
	Name     string `json:"name"`
	Salary   string `json:"salary"`
	Employee string `json:"employee"`
	Email    string `json:"email"`
	Remark   string `json:"remark"`
}

type KafkaRes struct {
	Data ResStatus `json:"data"`
}

type ResStatus struct {
	Status int32 `json:"status"` //状态  1.成功  2.失败
}
