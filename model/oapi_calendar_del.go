package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCalendarDeleteRequest() *OapiCalendarDeleteRequest {
	return &OapiCalendarDeleteRequest{}
}

type OapiCalendarDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarDeleteResponse
	CalendarId      string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCalendarDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarDeleteRequest) SetCalendarId(calendarId2 string) {
	this.CalendarId = calendarId2
}
func (this *OapiCalendarDeleteRequest) GetCalendarId() string {
	return this.CalendarId
}
func (this *OapiCalendarDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCalendarDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCalendarDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.delete"
}
func (this *OapiCalendarDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["calendar_id"] = this.CalendarId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCalendarDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCalendarDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCalendarDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
