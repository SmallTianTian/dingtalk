package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCalendarV2EventDetailRequest() *OapiCalendarV2EventDetailRequest {
	return &OapiCalendarV2EventDetailRequest{}
}

type OapiCalendarV2EventDetailRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarV2EventDetailResponse
	Agentid         int64
	CalendarId      string
	EventId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCalendarV2EventDetailRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarV2EventDetailRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarV2EventDetailRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiCalendarV2EventDetailRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiCalendarV2EventDetailRequest) SetCalendarId(calendarId2 string) {
	this.CalendarId = calendarId2
}
func (this *OapiCalendarV2EventDetailRequest) GetCalendarId() string {
	return this.CalendarId
}
func (this *OapiCalendarV2EventDetailRequest) SetEventId(eventId2 string) {
	this.EventId = eventId2
}
func (this *OapiCalendarV2EventDetailRequest) GetEventId() string {
	return this.EventId
}
func (this *OapiCalendarV2EventDetailRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.v2.event.detail"
}
func (this *OapiCalendarV2EventDetailRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarV2EventDetailRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarV2EventDetailRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarV2EventDetailRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarV2EventDetailRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["calendar_id"] = this.CalendarId
	txtParams["event_id"] = this.EventId
	return txtParams
}
func (this *OapiCalendarV2EventDetailRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CalendarId, "calendarId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCalendarV2EventDetailRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarV2EventDetailRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCalendarV2EventDetailResponse struct {
	taobao.TaobaoResponse
	Result  Event `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
