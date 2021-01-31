package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceScheduleListbyusersRequest() *OapiAttendanceScheduleListbyusersRequest {
	return &OapiAttendanceScheduleListbyusersRequest{}
}

type OapiAttendanceScheduleListbyusersRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceScheduleListbyusersResponse
	FromDateTime    int64
	OpUserId        string
	ToDateTime      int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiAttendanceScheduleListbyusersRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetFromDateTime(fromDateTime2 int64) {
	this.FromDateTime = fromDateTime2
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetFromDateTime() int64 {
	return this.FromDateTime
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetToDateTime(toDateTime2 int64) {
	this.ToDateTime = toDateTime2
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetToDateTime() int64 {
	return this.ToDateTime
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.schedule.listbyusers"
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceScheduleListbyusersRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["from_date_time"] = this.FromDateTime
	txtParams["op_user_id"] = this.OpUserId
	txtParams["to_date_time"] = this.ToDateTime
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiAttendanceScheduleListbyusersRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FromDateTime, "fromDateTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceScheduleListbyusersRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceScheduleListbyusersResponse struct {
	taobao.TaobaoResponse
	Result  []TopScheduleVo `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
