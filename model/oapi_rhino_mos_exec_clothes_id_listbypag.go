package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesIdListbypageRequest() *OapiRhinoMosExecClothesIdListbypageRequest {
	return &OapiRhinoMosExecClothesIdListbypageRequest{}
}

type OapiRhinoMosExecClothesIdListbypageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesIdListbypageResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.id.listbypage"
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesIdListbypageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PageQueryClothesReq struct {
	BizTypes   []string `json:"biz_types,omitempty"`
	OrderId    int64    `json:"order_id,omitempty"`
	Page       Page     `json:"page,omitempty"`
	SizeCode   string   `json:"size_code,omitempty"`
	Source     Source   `json:"source,omitempty"`
	StatusList []string `json:"status_list,omitempty"`
	TenantId   string   `json:"tenant_id,omitempty"`
	Userid     string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesIdListbypageResponse struct {
	taobao.TaobaoResponse
	Errcode int64   `json:"errcode,omitempty"`
	Errmsg  string  `json:"errmsg,omitempty"`
	Model   []int64 `json:"model,omitempty"`
	Success bool    `json:"success,omitempty"`
}
