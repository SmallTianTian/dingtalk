package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGetleavetimebynamesRequest() *OapiAttendanceGetleavetimebynamesRequest {
	return &OapiAttendanceGetleavetimebynamesRequest{}
}

type OapiAttendanceGetleavetimebynamesRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetleavetimebynamesResponse
	FromDate        time.Time
	LeaveNames      string
	ToDate          time.Time
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceGetleavetimebynamesRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetFromDate(fromDate2 time.Time) {
	this.FromDate = fromDate2
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetFromDate() time.Time {
	return this.FromDate
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetLeaveNames(leaveNames2 string) {
	this.LeaveNames = leaveNames2
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetLeaveNames() string {
	return this.LeaveNames
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetToDate(toDate2 time.Time) {
	this.ToDate = toDate2
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetToDate() time.Time {
	return this.ToDate
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getleavetimebynames"
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetleavetimebynamesRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["from_date"] = this.FromDate
	txtParams["leave_names"] = this.LeaveNames
	txtParams["to_date"] = this.ToDate
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceGetleavetimebynamesRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.LeaveNames, 20, "leaveNames"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetleavetimebynamesRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetleavetimebynamesResponse struct {
	taobao.TaobaoResponse
	Result ColumnValListForTopVo `json:"result,omitempty"`
}
