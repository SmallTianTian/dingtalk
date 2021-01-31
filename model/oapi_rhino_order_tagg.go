package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoOrderTagGetRequest() *OapiRhinoOrderTagGetRequest {
	return &OapiRhinoOrderTagGetRequest{}
}

type OapiRhinoOrderTagGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoOrderTagGetResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoOrderTagGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoOrderTagGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoOrderTagGetRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoOrderTagGetRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoOrderTagGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.order.tag.get"
}
func (this *OapiRhinoOrderTagGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoOrderTagGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoOrderTagGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoOrderTagGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoOrderTagGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoOrderTagGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoOrderTagGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoOrderTagGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type QueryTagByProductOrderReq struct {
	ProductOrderId int64  `json:"product_order_id,omitempty"`
	TenantId       string `json:"tenant_id,omitempty"`
	Userid         string `json:"userid,omitempty"`
}
type OapiRhinoOrderTagGetResponse struct {
	taobao.TaobaoResponse
	Model OrderTagDto `json:"model,omitempty"`
}
type OrderTagDto struct {
	BomReady        bool   `json:"bom_ready,omitempty"`
	CustomizeOrder  bool   `json:"customize_order,omitempty"`
	EmbroideryReady bool   `json:"embroidery_ready,omitempty"`
	GoodsNo         string `json:"goods_no,omitempty"`
	MarkerReady     bool   `json:"marker_ready,omitempty"`
	NoBom           bool   `json:"no_bom,omitempty"`
	NoGsd           bool   `json:"no_gsd,omitempty"`
	NoLaser         bool   `json:"no_laser,omitempty"`
	NoMarker        bool   `json:"no_marker,omitempty"`
	SkipSap         bool   `json:"skip_sap,omitempty"`
	SkipSupplyChain bool   `json:"skip_supply_chain,omitempty"`
}
