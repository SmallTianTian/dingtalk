package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoOrderDetailGetRequest() *OapiRhinoOrderDetailGetRequest {
	return &OapiRhinoOrderDetailGetRequest{}
}

type OapiRhinoOrderDetailGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoOrderDetailGetResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoOrderDetailGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoOrderDetailGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoOrderDetailGetRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoOrderDetailGetRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoOrderDetailGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.order.detail.get"
}
func (this *OapiRhinoOrderDetailGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoOrderDetailGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoOrderDetailGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoOrderDetailGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoOrderDetailGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoOrderDetailGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoOrderDetailGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoOrderDetailGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiGetProductOrderDetailReq struct {
	ProductOrderId int64  `json:"product_order_id,omitempty"`
	TenantId       string `json:"tenant_id,omitempty"`
	Userid         string `json:"userid,omitempty"`
}
type OapiRhinoOrderDetailGetResponse struct {
	taobao.TaobaoResponse
	Model Model `json:"model,omitempty"`
}
type OpenApiProductOrderDetailDto struct {
	Id             int64  `json:"id,omitempty"`
	ProductOrderId int64  `json:"product_order_id,omitempty"`
	Quantity       int64  `json:"quantity,omitempty"`
	SizeId         int64  `json:"size_id,omitempty"`
	SizeName       string `json:"size_name,omitempty"`
	TenantId       string `json:"tenant_id,omitempty"`
}
