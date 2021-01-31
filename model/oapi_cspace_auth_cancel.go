package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCspaceAuthCancelRequest() *OapiCspaceAuthCancelRequest {
	return &OapiCspaceAuthCancelRequest{}
}

type OapiCspaceAuthCancelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceAuthCancelResponse
	AgentId         int64
	IsvCode         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCspaceAuthCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCspaceAuthCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceAuthCancelRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiCspaceAuthCancelRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiCspaceAuthCancelRequest) SetIsvCode(isvCode2 string) {
	this.IsvCode = isvCode2
}
func (this *OapiCspaceAuthCancelRequest) GetIsvCode() string {
	return this.IsvCode
}
func (this *OapiCspaceAuthCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.auth.cancel"
}
func (this *OapiCspaceAuthCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceAuthCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceAuthCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceAuthCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceAuthCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["isv_code"] = this.IsvCode
	return txtParams
}
func (this *OapiCspaceAuthCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCspaceAuthCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceAuthCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceAuthCancelResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
