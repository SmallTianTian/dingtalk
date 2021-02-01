package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectmetaContactDescribeRequest() *OapiCrmObjectmetaContactDescribeRequest {
	return &OapiCrmObjectmetaContactDescribeRequest{}
}

type OapiCrmObjectmetaContactDescribeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectmetaContactDescribeResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectmetaContactDescribeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectmetaContactDescribeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectmetaContactDescribeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectmeta.contact.describe"
}
func (this *OapiCrmObjectmetaContactDescribeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectmetaContactDescribeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectmetaContactDescribeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectmetaContactDescribeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectmetaContactDescribeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCrmObjectmetaContactDescribeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectmetaContactDescribeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectmetaContactDescribeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectmetaContactDescribeResponse struct {
	taobao.TaobaoResponse
	Result DObject `json:"result,omitempty"`
}
type SelectOptions struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
type ReferenceFields struct {
	Format        string          `json:"format,omitempty"`
	Label         string          `json:"label,omitempty"`
	Name          string          `json:"name,omitempty"`
	Nillable      bool            `json:"nillable,omitempty"`
	SelectOptions []SelectOptions `json:"select_options,omitempty"`
	Type          string          `json:"type,omitempty"`
	Unit          string          `json:"unit,omitempty"`
}
type RollUpSummaryFields struct {
	Aggregator string `json:"aggregator,omitempty"`
	Name       string `json:"name,omitempty"`
}
type Fields struct {
	Customized          bool                  `json:"customized,omitempty"`
	Format              string                `json:"format,omitempty"`
	Label               string                `json:"label,omitempty"`
	Name                string                `json:"name,omitempty"`
	Nillable            bool                  `json:"nillable,omitempty"`
	Quote               bool                  `json:"quote,omitempty"`
	ReferenceFields     []ReferenceFields     `json:"reference_fields,omitempty"`
	ReferenceTo         string                `json:"reference_to,omitempty"`
	RollUpSummaryFields []RollUpSummaryFields `json:"roll_up_summary_fields,omitempty"`
	SelectOptions       []SelectOptions       `json:"select_options,omitempty"`
	Type                string                `json:"type,omitempty"`
	Unit                string                `json:"unit,omitempty"`
}
type DObject struct {
	Customized bool     `json:"customized,omitempty"`
	Fields     []Fields `json:"fields,omitempty"`
	Name       string   `json:"name,omitempty"`
}
