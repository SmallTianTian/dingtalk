package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecClothesGetRequest() *OapiRhinoMosExecClothesGetRequest {
	return &OapiRhinoMosExecClothesGetRequest{}
}

type OapiRhinoMosExecClothesGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesGetResponse
	EntityIds       string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosExecClothesGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesGetRequest) SetEntityIds(entityIds2 string) {
	this.EntityIds = entityIds2
}
func (this *OapiRhinoMosExecClothesGetRequest) GetEntityIds() string {
	return this.EntityIds
}
func (this *OapiRhinoMosExecClothesGetRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecClothesGetRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecClothesGetRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecClothesGetRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecClothesGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecClothesGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecClothesGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.get"
}
func (this *OapiRhinoMosExecClothesGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["entity_ids"] = this.EntityIds
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecClothesGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EntityIds, "entityIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecClothesGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecClothesGetResponse struct {
	taobao.TaobaoResponse
	Errcode         int64        `json:"errcode,omitempty"`
	Errmsg          string       `json:"errmsg,omitempty"`
	ExternalMsgInfo string       `json:"external_msg_info,omitempty"`
	Model           []ClothesDto `json:"model,omitempty"`
	Success         bool         `json:"success,omitempty"`
}
