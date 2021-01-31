package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCalendarV2AttendeeUpdateRequest() *OapiCalendarV2AttendeeUpdateRequest {
	return &OapiCalendarV2AttendeeUpdateRequest{}
}

type OapiCalendarV2AttendeeUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarV2AttendeeUpdateResponse
	Agentid         int64
	Attendees       string
	CalendarId      string
	EventId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCalendarV2AttendeeUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetAttendees(attendees2 string) {
	this.Attendees = attendees2
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetAttendees() string {
	return this.Attendees
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetCalendarId(calendarId2 string) {
	this.CalendarId = calendarId2
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetCalendarId() string {
	return this.CalendarId
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetEventId(eventId2 string) {
	this.EventId = eventId2
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetEventId() string {
	return this.EventId
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.v2.attendee.update"
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarV2AttendeeUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["attendees"] = this.Attendees
	txtParams["calendar_id"] = this.CalendarId
	txtParams["event_id"] = this.EventId
	return txtParams
}
func (this *OapiCalendarV2AttendeeUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Attendees, 20, "attendees"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarV2AttendeeUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Attendee struct {
	AttendeeStatus string `json:"attendee_status,omitempty"`
	Userid         string `json:"userid,omitempty"`
}
type OapiCalendarV2AttendeeUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
