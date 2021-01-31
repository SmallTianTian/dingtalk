package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceActionQueryRequest() *OapiCustomerserviceActionQueryRequest {
	return &OapiCustomerserviceActionQueryRequest{}
}

type OapiCustomerserviceActionQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCustomerserviceActionQueryResponse
	TicketActionPageQuery string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCustomerserviceActionQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceActionQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceActionQueryRequest) SetTicketActionPageQuery(ticketActionPageQuery2 string) {
	this.TicketActionPageQuery = ticketActionPageQuery2
}
func (this *OapiCustomerserviceActionQueryRequest) GetTicketActionPageQuery() string {
	return this.TicketActionPageQuery
}
func (this *OapiCustomerserviceActionQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.action.query"
}
func (this *OapiCustomerserviceActionQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceActionQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceActionQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceActionQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceActionQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ticket_action_page_query"] = this.TicketActionPageQuery
	return txtParams
}
func (this *OapiCustomerserviceActionQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceActionQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceActionQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TicketActionPageQueryDto struct {
	Cursor   int64  `json:"cursor,omitempty"`
	Size     int64  `json:"size,omitempty"`
	TicketId string `json:"ticket_id,omitempty"`
}
type OapiCustomerserviceActionQueryResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type TicketFieldDto struct {
	DisplayName  string `json:"display_name,omitempty"`
	DisplayValue string `json:"display_value,omitempty"`
	Name         string `json:"name,omitempty"`
	Value        string `json:"value,omitempty"`
	ValueType    string `json:"value_type,omitempty"`
}
type TicketActionDto struct {
	ActionContent []TicketFieldDto `json:"action_content,omitempty"`
	Operator      string           `json:"operator,omitempty"`
	OperatorRole  string           `json:"operator_role,omitempty"`
}
type PageQueryCursorResult struct {
	ActionList []TicketActionDto `json:"action_list,omitempty"`
	HasMore    bool              `json:"has_more,omitempty"`
	NextCursor int64             `json:"next_cursor,omitempty"`
	Total      int64             `json:"total,omitempty"`
}
