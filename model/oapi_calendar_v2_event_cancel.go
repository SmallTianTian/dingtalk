package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCalendarV2EventCancelRequest() *OapiCalendarV2EventCancelRequest {
	return &OapiCalendarV2EventCancelRequest{}
}

type OapiCalendarV2EventCancelRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarV2EventCancelResponse
	Agentid         int64
	CalendarId      string
	EventId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCalendarV2EventCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarV2EventCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarV2EventCancelRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiCalendarV2EventCancelRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiCalendarV2EventCancelRequest) SetCalendarId(calendarId2 string) {
	this.CalendarId = calendarId2
}
func (this *OapiCalendarV2EventCancelRequest) GetCalendarId() string {
	return this.CalendarId
}
func (this *OapiCalendarV2EventCancelRequest) SetEventId(eventId2 string) {
	this.EventId = eventId2
}
func (this *OapiCalendarV2EventCancelRequest) GetEventId() string {
	return this.EventId
}
func (this *OapiCalendarV2EventCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.v2.event.cancel"
}
func (this *OapiCalendarV2EventCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarV2EventCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarV2EventCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarV2EventCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarV2EventCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["calendar_id"] = this.CalendarId
	txtParams["event_id"] = this.EventId
	return txtParams
}
func (this *OapiCalendarV2EventCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CalendarId, "calendarId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCalendarV2EventCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarV2EventCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCalendarV2EventCancelResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
