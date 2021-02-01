package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceCircleGroupCreateRequest() *OapiWorkspaceCircleGroupCreateRequest {
	return &OapiWorkspaceCircleGroupCreateRequest{}
}

type OapiWorkspaceCircleGroupCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCircleGroupCreateResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceCircleGroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCircleGroupCreateRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetReq() string {
	return this.Req
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.circle.group.create"
}
func (this *OapiWorkspaceCircleGroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCircleGroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCircleGroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiWorkspaceCircleGroupCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCircleGroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCreateGroupRequestDto struct {
	Name   string `json:"name,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OapiWorkspaceCircleGroupCreateResponse struct {
	taobao.TaobaoResponse
	Result OpenCreateGroupResponseDto `json:"result,omitempty"`
}
type OpenCreateGroupResponseDto struct {
	ConversationId string `json:"conversation_id,omitempty"`
}
