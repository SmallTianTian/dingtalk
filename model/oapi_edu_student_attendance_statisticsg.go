package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduStudentAttendanceStatisticsGetRequest() *OapiEduStudentAttendanceStatisticsGetRequest {
	return &OapiEduStudentAttendanceStatisticsGetRequest{}
}

type OapiEduStudentAttendanceStatisticsGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduStudentAttendanceStatisticsGetResponse
	Date            string
	OpUserid        string
	SchoolCorpid    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) SetDate(date2 string) {
	this.Date = date2
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetDate() string {
	return this.Date
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) SetSchoolCorpid(schoolCorpid2 string) {
	this.SchoolCorpid = schoolCorpid2
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetSchoolCorpid() string {
	return this.SchoolCorpid
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.student.attendance.statistics.get"
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["date"] = this.Date
	txtParams["op_userid"] = this.OpUserid
	txtParams["school_corpid"] = this.SchoolCorpid
	return txtParams
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Date, "date"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduStudentAttendanceStatisticsGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduStudentAttendanceStatisticsGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                                  `json:"errcode,omitempty"`
	Errmsg  string                                 `json:"errmsg,omitempty"`
	Result  GetStudentAttendanceStatisticsResponse `json:"result,omitempty"`
	Success bool                                   `json:"success,omitempty"`
}
type GetStudentAttendanceStatisticsResponse struct {
	Data string `json:"data,omitempty"`
}
