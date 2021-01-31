package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceTicketCreateRequest() *OapiCustomerserviceTicketCreateRequest {
	return &OapiCustomerserviceTicketCreateRequest{}
}

type OapiCustomerserviceTicketCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceTicketCreateResponse
	TicketCreate    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceTicketCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceTicketCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceTicketCreateRequest) SetTicketCreate(ticketCreate2 string) {
	this.TicketCreate = ticketCreate2
}
func (this *OapiCustomerserviceTicketCreateRequest) GetTicketCreate() string {
	return this.TicketCreate
}
func (this *OapiCustomerserviceTicketCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.ticket.create"
}
func (this *OapiCustomerserviceTicketCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceTicketCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceTicketCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceTicketCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceTicketCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ticket_create"] = this.TicketCreate
	return txtParams
}
func (this *OapiCustomerserviceTicketCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceTicketCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceTicketCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TicketCreateDto struct {
	ForeignId   string           `json:"foreign_id,omitempty"`
	ForeignName string           `json:"foreign_name,omitempty"`
	Properties  []TicketFieldDto `json:"properties,omitempty"`
	SourceId    string           `json:"source_id,omitempty"`
	TemplateId  string           `json:"template_id,omitempty"`
	Title       string           `json:"title,omitempty"`
}
type OapiCustomerserviceTicketCreateResponse struct {
	taobao.TaobaoResponse
	Result  TicketCreateResultDto `json:"result,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
type TicketCreateResultDto struct {
	TicketId string `json:"ticket_id,omitempty"`
}
