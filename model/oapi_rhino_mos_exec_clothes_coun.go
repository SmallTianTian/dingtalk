package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesCountRequest() *OapiRhinoMosExecClothesCountRequest {
	return &OapiRhinoMosExecClothesCountRequest{}
}

type OapiRhinoMosExecClothesCountRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesCountResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesCountRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesCountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesCountRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosExecClothesCountRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosExecClothesCountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.count"
}
func (this *OapiRhinoMosExecClothesCountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesCountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesCountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesCountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesCountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosExecClothesCountRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesCountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesCountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ClothesSearchCondition struct {
	ClothesStatusList []string `json:"clothes_status_list,omitempty"`
	OrderId           int64    `json:"order_id,omitempty"`
	TenantId          string   `json:"tenant_id,omitempty"`
	Userid            string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesCountResponse struct {
	taobao.TaobaoResponse
	Model int64 `json:"model,omitempty"`
}
