package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceActivityExecuteRequest() *OapiCustomerserviceActivityExecuteRequest {
	return &OapiCustomerserviceActivityExecuteRequest{}
}

type OapiCustomerserviceActivityExecuteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceActivityExecuteResponse
	TicketActivity  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceActivityExecuteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceActivityExecuteRequest) SetTicketActivity(ticketActivity2 string) {
	this.TicketActivity = ticketActivity2
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetTicketActivity() string {
	return this.TicketActivity
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.activity.execute"
}
func (this *OapiCustomerserviceActivityExecuteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceActivityExecuteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceActivityExecuteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ticket_activity"] = this.TicketActivity
	return txtParams
}
func (this *OapiCustomerserviceActivityExecuteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceActivityExecuteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TicketActivityDto struct {
	ActivityCode string           `json:"activity_code,omitempty"`
	ForeignId    string           `json:"foreign_id,omitempty"`
	ForeignName  string           `json:"foreign_name,omitempty"`
	Properties   []TicketFieldDto `json:"properties,omitempty"`
	SourceId     string           `json:"source_id,omitempty"`
	TicketId     string           `json:"ticket_id,omitempty"`
}
type OapiCustomerserviceActivityExecuteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
