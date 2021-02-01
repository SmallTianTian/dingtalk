package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesGroupbyoperationCountRequest() *OapiRhinoMosExecClothesGroupbyoperationCountRequest {
	return &OapiRhinoMosExecClothesGroupbyoperationCountRequest{}
}

type OapiRhinoMosExecClothesGroupbyoperationCountRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesGroupbyoperationCountResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.groupbyoperation.count"
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesGroupbyoperationCountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ClothesGroupByOperationCondition struct {
	ActiveCondition   string   `json:"active_condition,omitempty"`
	ClothesStatusList []string `json:"clothes_status_list,omitempty"`
	OrderId           int64    `json:"order_id,omitempty"`
	PerformStatusList []string `json:"perform_status_list,omitempty"`
	TenantId          string   `json:"tenant_id,omitempty"`
	Userid            string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesGroupbyoperationCountResponse struct {
	taobao.TaobaoResponse
	Model []CountGroupByIdDto `json:"model,omitempty"`
}
type CountGroupByIdDto struct {
	Count int64 `json:"count,omitempty"`
	Id    int64 `json:"id,omitempty"`
}
