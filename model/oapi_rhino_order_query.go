package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoOrderQueryRequest() *OapiRhinoOrderQueryRequest {
	return &OapiRhinoOrderQueryRequest{}
}

type OapiRhinoOrderQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoOrderQueryResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoOrderQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoOrderQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoOrderQueryRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoOrderQueryRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoOrderQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.order.query"
}
func (this *OapiRhinoOrderQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoOrderQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoOrderQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoOrderQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoOrderQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoOrderQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoOrderQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoOrderQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiQueryProductOrderReq struct {
	KeyWord       string    `json:"key_word,omitempty"`
	Page          int64     `json:"page,omitempty"`
	PageSize      int64     `json:"page_size,omitempty"`
	PlanTimeBegin time.Time `json:"plan_time_begin,omitempty"`
	PlanTimeEnd   time.Time `json:"plan_time_end,omitempty"`
	Sort          string    `json:"sort,omitempty"`
	SortAsc       bool      `json:"sort_asc,omitempty"`
	StatusList    []string  `json:"status_list,omitempty"`
	TenantId      string    `json:"tenant_id,omitempty"`
	Userid        string    `json:"userid,omitempty"`
}
type OapiRhinoOrderQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Model   PageInfo `json:"model,omitempty"`
}
type PageInfo struct {
	List     []ProductOrderDto `json:"list,omitempty"`
	Page     int64             `json:"page,omitempty"`
	PageSize int64             `json:"page_size,omitempty"`
	Total    int64             `json:"total,omitempty"`
}
