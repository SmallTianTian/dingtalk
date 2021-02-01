package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceProjectQueryRequest() *OapiWorkspaceProjectQueryRequest {
	return &OapiWorkspaceProjectQueryRequest{}
}

type OapiWorkspaceProjectQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceProjectQueryResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceProjectQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceProjectQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceProjectQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.project.query"
}
func (this *OapiWorkspaceProjectQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceProjectQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceProjectQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceProjectQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceProjectQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiWorkspaceProjectQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceProjectQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceProjectQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceProjectQueryResponse struct {
	taobao.TaobaoResponse
	Result  OpenProjectDto `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
