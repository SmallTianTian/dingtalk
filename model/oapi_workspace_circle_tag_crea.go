package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceCircleTagCreateRequest() *OapiWorkspaceCircleTagCreateRequest {
	return &OapiWorkspaceCircleTagCreateRequest{}
}

type OapiWorkspaceCircleTagCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCircleTagCreateResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceCircleTagCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCircleTagCreateRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetReq() string {
	return this.Req
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.circle.tag.create"
}
func (this *OapiWorkspaceCircleTagCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCircleTagCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCircleTagCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiWorkspaceCircleTagCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCircleTagCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FvTagCreateOpenDto struct {
	Name   string `json:"name,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OapiWorkspaceCircleTagCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  FvPostTagOpenDto `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
