package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceApproveDurationCalculateRequest() *OapiAttendanceApproveDurationCalculateRequest {
	return &OapiAttendanceApproveDurationCalculateRequest{}
}

type OapiAttendanceApproveDurationCalculateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceApproveDurationCalculateResponse
	BizType         int64
	CalculateModel  int64
	DurationUnit    string
	FromTime        string
	ToTime          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceApproveDurationCalculateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetCalculateModel(calculateModel2 int64) {
	this.CalculateModel = calculateModel2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetCalculateModel() int64 {
	return this.CalculateModel
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetDurationUnit(durationUnit2 string) {
	this.DurationUnit = durationUnit2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetDurationUnit() string {
	return this.DurationUnit
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetFromTime(fromTime2 string) {
	this.FromTime = fromTime2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetFromTime() string {
	return this.FromTime
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetToTime(toTime2 string) {
	this.ToTime = toTime2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetToTime() string {
	return this.ToTime
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.approve.duration.calculate"
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceApproveDurationCalculateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["calculate_model"] = this.CalculateModel
	txtParams["duration_unit"] = this.DurationUnit
	txtParams["from_time"] = this.FromTime
	txtParams["to_time"] = this.ToTime
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceApproveDurationCalculateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizType, "bizType"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceApproveDurationCalculateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceApproveDurationCalculateResponse struct {
	taobao.TaobaoResponse
	Result TopDurationVo `json:"result,omitempty"`
}
type TopDayDurationVo struct {
	Date     string `json:"date,omitempty"`
	Duration string `json:"duration,omitempty"`
}
type TopDurationVo struct {
	Duration        string             `json:"duration,omitempty"`
	DurationDetails []TopDayDurationVo `json:"duration_details,omitempty"`
}
