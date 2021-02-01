package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCalendarV2EventUpdateRequest() *OapiCalendarV2EventUpdateRequest {
	return &OapiCalendarV2EventUpdateRequest{}
}

type OapiCalendarV2EventUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarV2EventUpdateResponse
	Agentid         int64
	Event           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCalendarV2EventUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarV2EventUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarV2EventUpdateRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiCalendarV2EventUpdateRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiCalendarV2EventUpdateRequest) SetEvent(event2 string) {
	this.Event = event2
}
func (this *OapiCalendarV2EventUpdateRequest) GetEvent() string {
	return this.Event
}
func (this *OapiCalendarV2EventUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.v2.event.update"
}
func (this *OapiCalendarV2EventUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarV2EventUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarV2EventUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarV2EventUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarV2EventUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["event"] = this.Event
	return txtParams
}
func (this *OapiCalendarV2EventUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCalendarV2EventUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarV2EventUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCalendarV2EventUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
