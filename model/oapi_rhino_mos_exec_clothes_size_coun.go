package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesSizeCountRequest() *OapiRhinoMosExecClothesSizeCountRequest {
	return &OapiRhinoMosExecClothesSizeCountRequest{}
}

type OapiRhinoMosExecClothesSizeCountRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesSizeCountResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesSizeCountRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.size.count"
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesSizeCountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ClothesGroupBySizeConditionReq struct {
	ClothesStatusList []string `json:"clothes_status_list,omitempty"`
	OrderId           int64    `json:"order_id,omitempty"`
	TenantId          string   `json:"tenant_id,omitempty"`
	Userid            string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesSizeCountResponse struct {
	taobao.TaobaoResponse
	Model   []ClothesCountGroupBySizeDto `json:"model,omitempty"`
	Success bool                         `json:"success,omitempty"`
}
type ClothesCountGroupBySizeDto struct {
	Count    int64  `json:"count,omitempty"`
	SizeCode string `json:"size_code,omitempty"`
}
