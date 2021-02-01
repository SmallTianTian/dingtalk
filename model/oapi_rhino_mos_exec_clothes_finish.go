package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecClothesFinishRequest() *OapiRhinoMosExecClothesFinishRequest {
	return &OapiRhinoMosExecClothesFinishRequest{}
}

type OapiRhinoMosExecClothesFinishRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesFinishResponse
	EntityIds       string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosExecClothesFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetEntityIds(entityIds2 string) {
	this.EntityIds = entityIds2
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetEntityIds() string {
	return this.EntityIds
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.finish"
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["entity_ids"] = this.EntityIds
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecClothesFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EntityIds, "entityIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecClothesFinishResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
