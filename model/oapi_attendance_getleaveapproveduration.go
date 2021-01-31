package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAttendanceGetleaveapprovedurationRequest() *OapiAttendanceGetleaveapprovedurationRequest {
	return &OapiAttendanceGetleaveapprovedurationRequest{}
}

type OapiAttendanceGetleaveapprovedurationRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetleaveapprovedurationResponse
	FromDate        time.Time
	ToDate          time.Time
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceGetleaveapprovedurationRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) SetFromDate(fromDate2 time.Time) {
	this.FromDate = fromDate2
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetFromDate() time.Time {
	return this.FromDate
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) SetToDate(toDate2 time.Time) {
	this.ToDate = toDate2
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetToDate() time.Time {
	return this.ToDate
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getleaveapproveduration"
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["from_date"] = this.FromDate
	txtParams["to_date"] = this.ToDate
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FromDate, "fromDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetleaveapprovedurationRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetleaveapprovedurationResponse struct {
	taobao.TaobaoResponse
	Result ApproveDurationForTopVo `json:"result,omitempty"`
}
type ApproveDurationForTopVo struct {
	DurationInMinutes int64 `json:"duration_in_minutes,omitempty"`
}
