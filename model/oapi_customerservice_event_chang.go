package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceEventChangeRequest() *OapiCustomerserviceEventChangeRequest {
	return &OapiCustomerserviceEventChangeRequest{}
}

type OapiCustomerserviceEventChangeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceEventChangeResponse
	EventDto        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceEventChangeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceEventChangeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceEventChangeRequest) SetEventDto(eventDto2 string) {
	this.EventDto = eventDto2
}
func (this *OapiCustomerserviceEventChangeRequest) GetEventDto() string {
	return this.EventDto
}
func (this *OapiCustomerserviceEventChangeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.event.change"
}
func (this *OapiCustomerserviceEventChangeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceEventChangeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceEventChangeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceEventChangeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceEventChangeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["event_dto"] = this.EventDto
	return txtParams
}
func (this *OapiCustomerserviceEventChangeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceEventChangeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceEventChangeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EventDto struct {
	BizType   string `json:"biz_type,omitempty"`
	BuId      string `json:"bu_id,omitempty"`
	EventBody string `json:"event_body,omitempty"`
	EventCode string `json:"event_code,omitempty"`
	EventId   string `json:"event_id,omitempty"`
	Source    string `json:"source,omitempty"`
}
type OapiCustomerserviceEventChangeResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
