package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesBatchUnscrapRequest() *OapiRhinoMosExecClothesBatchUnscrapRequest {
	return &OapiRhinoMosExecClothesBatchUnscrapRequest{}
}

type OapiRhinoMosExecClothesBatchUnscrapRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiRhinoMosExecClothesBatchUnscrapResponse
	BatchClothesPerformReq string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) SetBatchClothesPerformReq(batchClothesPerformReq2 string) {
	this.BatchClothesPerformReq = batchClothesPerformReq2
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetBatchClothesPerformReq() string {
	return this.BatchClothesPerformReq
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.batch.unscrap"
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["batch_clothes_perform_req"] = this.BatchClothesPerformReq
	return txtParams
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesBatchUnscrapRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BatchClothesPerformBaseReq struct {
	EntityIds     []int64 `json:"entity_ids,omitempty"`
	ExtProperties string  `json:"ext_properties,omitempty"`
	OrderId       int64   `json:"order_id,omitempty"`
	TenantId      string  `json:"tenant_id,omitempty"`
	Userid        string  `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesBatchUnscrapResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Model   bool   `json:"model,omitempty"`
	Success bool   `json:"success,omitempty"`
}
