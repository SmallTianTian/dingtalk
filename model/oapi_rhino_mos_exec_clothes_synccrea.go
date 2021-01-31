package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecClothesSynccreateRequest() *OapiRhinoMosExecClothesSynccreateRequest {
	return &OapiRhinoMosExecClothesSynccreateRequest{}
}

type OapiRhinoMosExecClothesSynccreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiRhinoMosExecClothesSynccreateResponse
	AdditionalOperations string
	AutoStart            bool
	BizType              string
	Clothes              string
	EntityType           string
	OrderId              int64
	Source               string
	TenantId             string
	TopHttpMethod        string
	TopResponseType      string
	Userid               string
}

func (this *OapiRhinoMosExecClothesSynccreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetAdditionalOperations(additionalOperations2 string) {
	this.AdditionalOperations = additionalOperations2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetAdditionalOperations() string {
	return this.AdditionalOperations
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetAutoStart(autoStart2 bool) {
	this.AutoStart = autoStart2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetAutoStart() bool {
	return this.AutoStart
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetBizType(bizType2 string) {
	this.BizType = bizType2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetBizType() string {
	return this.BizType
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetClothes(clothes2 string) {
	this.Clothes = clothes2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetClothes() string {
	return this.Clothes
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetEntityType(entityType2 string) {
	this.EntityType = entityType2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetEntityType() string {
	return this.EntityType
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetSource(source2 string) {
	this.Source = source2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetSource() string {
	return this.Source
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.synccreate"
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["additional_operations"] = this.AdditionalOperations
	txtParams["auto_start"] = this.AutoStart
	txtParams["biz_type"] = this.BizType
	txtParams["clothes"] = this.Clothes
	txtParams["entity_type"] = this.EntityType
	txtParams["order_id"] = this.OrderId
	txtParams["source"] = this.Source
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.AdditionalOperations, 500, "additionalOperations"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesSynccreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecClothesSynccreateResponse struct {
	taobao.TaobaoResponse
	Model   []int64 `json:"model,omitempty"`
	Success bool    `json:"success,omitempty"`
}
