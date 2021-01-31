package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoCooperationCarrierGetRequest() *OapiRhinoCooperationCarrierGetRequest {
	return &OapiRhinoCooperationCarrierGetRequest{}
}

type OapiRhinoCooperationCarrierGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoCooperationCarrierGetResponse
	CarrierId       int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoCooperationCarrierGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoCooperationCarrierGetRequest) SetCarrierId(carrierId2 int64) {
	this.CarrierId = carrierId2
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetCarrierId() int64 {
	return this.CarrierId
}
func (this *OapiRhinoCooperationCarrierGetRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoCooperationCarrierGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.cooperation.carrier.get"
}
func (this *OapiRhinoCooperationCarrierGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoCooperationCarrierGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoCooperationCarrierGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["carrier_id"] = this.CarrierId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoCooperationCarrierGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoCooperationCarrierGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoCooperationCarrierGetResponse struct {
	taobao.TaobaoResponse
	Errcode         int64      `json:"errcode,omitempty"`
	Errmsg          string     `json:"errmsg,omitempty"`
	ExternalMsgInfo string     `json:"external_msg_info,omitempty"`
	Model           CarrierDto `json:"model,omitempty"`
	Success         bool       `json:"success,omitempty"`
}
type CarrierDto struct {
	CarrierCode     string `json:"carrier_code,omitempty"`
	CarrierId       int64  `json:"carrier_id,omitempty"`
	CarrierName     string `json:"carrier_name,omitempty"`
	CarrierType     int64  `json:"carrier_type,omitempty"`
	CarrierTypeName string `json:"carrier_type_name,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	Status          int64  `json:"status,omitempty"`
	TenantId        string `json:"tenant_id,omitempty"`
}
