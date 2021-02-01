package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecClothesScrapRequest() *OapiRhinoMosExecClothesScrapRequest {
	return &OapiRhinoMosExecClothesScrapRequest{}
}

type OapiRhinoMosExecClothesScrapRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesScrapResponse
	EntityIds       string
	ExtProperties   string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosExecClothesScrapRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetEntityIds(entityIds2 string) {
	this.EntityIds = entityIds2
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetEntityIds() string {
	return this.EntityIds
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetExtProperties(extProperties2 string) {
	this.ExtProperties = extProperties2
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetExtPropertiesString(extProperties2 string) {
	this.ExtProperties = extProperties2
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetExtProperties() string {
	return this.ExtProperties
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.scrap"
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesScrapRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["entity_ids"] = this.EntityIds
	txtParams["ext_properties"] = this.ExtProperties
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecClothesScrapRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EntityIds, "entityIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesScrapRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecClothesScrapResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
