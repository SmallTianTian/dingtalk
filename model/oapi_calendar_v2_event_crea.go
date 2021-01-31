package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCalendarV2EventCreateRequest() *OapiCalendarV2EventCreateRequest {
	return &OapiCalendarV2EventCreateRequest{}
}

type OapiCalendarV2EventCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarV2EventCreateResponse
	Agentid         int64
	Event           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCalendarV2EventCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarV2EventCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarV2EventCreateRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiCalendarV2EventCreateRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiCalendarV2EventCreateRequest) SetEvent(event2 string) {
	this.Event = event2
}
func (this *OapiCalendarV2EventCreateRequest) GetEvent() string {
	return this.Event
}
func (this *OapiCalendarV2EventCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.v2.event.create"
}
func (this *OapiCalendarV2EventCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarV2EventCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarV2EventCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarV2EventCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarV2EventCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["event"] = this.Event
	return txtParams
}
func (this *OapiCalendarV2EventCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCalendarV2EventCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarV2EventCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type LocationVo struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Place     string `json:"place,omitempty"`
}
type Event struct {
	Attendees        []Attendee             `json:"attendees,omitempty"`
	CalendarId       string                 `json:"calendar_id,omitempty"`
	Description      string                 `json:"description,omitempty"`
	End              DateTime               `json:"end,omitempty"`
	Location         LocationVo             `json:"location,omitempty"`
	NotificationType string                 `json:"notification_type,omitempty"`
	Organizer        Attendee               `json:"organizer,omitempty"`
	Reminder         OpenCalendarReminderVo `json:"reminder,omitempty"`
	Start            DateTime               `json:"start,omitempty"`
	Summary          string                 `json:"summary,omitempty"`
}
type OapiCalendarV2EventCreateResponse struct {
	taobao.TaobaoResponse
	Result  Event `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
