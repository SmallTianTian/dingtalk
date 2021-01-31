package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGetleavestatusRequest() *OapiAttendanceGetleavestatusRequest {
	return &OapiAttendanceGetleavestatusRequest{}
}

type OapiAttendanceGetleavestatusRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetleavestatusResponse
	EndTime         int64
	Offset          int64
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiAttendanceGetleavestatusRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetleavestatusRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetleavestatusRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiAttendanceGetleavestatusRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiAttendanceGetleavestatusRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiAttendanceGetleavestatusRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiAttendanceGetleavestatusRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceGetleavestatusRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceGetleavestatusRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiAttendanceGetleavestatusRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiAttendanceGetleavestatusRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiAttendanceGetleavestatusRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiAttendanceGetleavestatusRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getleavestatus"
}
func (this *OapiAttendanceGetleavestatusRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetleavestatusRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetleavestatusRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetleavestatusRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetleavestatusRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["end_time"] = this.EndTime
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiAttendanceGetleavestatusRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EndTime, "endTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetleavestatusRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetleavestatusRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetleavestatusResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  LeaveStatusListVO `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type LeaveStatusVO struct {
	DurationPercent int64  `json:"duration_percent,omitempty"`
	DurationUnit    string `json:"duration_unit,omitempty"`
	EndTime         int64  `json:"end_time,omitempty"`
	StartTime       int64  `json:"start_time,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
type LeaveStatusListVO struct {
	HasMore     bool            `json:"has_more,omitempty"`
	LeaveStatus []LeaveStatusVO `json:"leave_status,omitempty"`
}
