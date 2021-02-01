package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceTicketQueryRequest() *OapiCustomerserviceTicketQueryRequest {
	return &OapiCustomerserviceTicketQueryRequest{}
}

type OapiCustomerserviceTicketQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiCustomerserviceTicketQueryResponse
	TicketPageQueryDto string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiCustomerserviceTicketQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceTicketQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceTicketQueryRequest) SetTicketPageQueryDto(ticketPageQueryDto2 string) {
	this.TicketPageQueryDto = ticketPageQueryDto2
}
func (this *OapiCustomerserviceTicketQueryRequest) GetTicketPageQueryDto() string {
	return this.TicketPageQueryDto
}
func (this *OapiCustomerserviceTicketQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.ticket.query"
}
func (this *OapiCustomerserviceTicketQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceTicketQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceTicketQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceTicketQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceTicketQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ticket_page_query_dto"] = this.TicketPageQueryDto
	return txtParams
}
func (this *OapiCustomerserviceTicketQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceTicketQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceTicketQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TicketPageQueryDto struct {
	Cursor           int64     `json:"cursor,omitempty"`
	EndDate          time.Time `json:"end_date,omitempty"`
	ForeignId        string    `json:"foreign_id,omitempty"`
	ForeignName      string    `json:"foreign_name,omitempty"`
	Size             int64     `json:"size,omitempty"`
	SourceId         string    `json:"source_id,omitempty"`
	StartDate        time.Time `json:"start_date,omitempty"`
	TicketId         string    `json:"ticket_id,omitempty"`
	TicketStatus     string    `json:"ticket_status,omitempty"`
	TicketTemplateId string    `json:"ticket_template_id,omitempty"`
}
type OapiCustomerserviceTicketQueryResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type TicketBaseDto struct {
	GmtCreate    time.Time `json:"gmt_create,omitempty"`
	GmtModified  time.Time `json:"gmt_modified,omitempty"`
	TemplateId   string    `json:"template_id,omitempty"`
	TicketId     string    `json:"ticket_id,omitempty"`
	TicketStatus string    `json:"ticket_status,omitempty"`
	Title        string    `json:"title,omitempty"`
}
