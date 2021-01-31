package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCspaceAuthUpdateRequest() *OapiCspaceAuthUpdateRequest {
	return &OapiCspaceAuthUpdateRequest{}
}

type OapiCspaceAuthUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceAuthUpdateResponse
	AgentId         int64
	Duration        int64
	IsvCode         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCspaceAuthUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCspaceAuthUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceAuthUpdateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiCspaceAuthUpdateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiCspaceAuthUpdateRequest) SetDuration(duration2 int64) {
	this.Duration = duration2
}
func (this *OapiCspaceAuthUpdateRequest) GetDuration() int64 {
	return this.Duration
}
func (this *OapiCspaceAuthUpdateRequest) SetIsvCode(isvCode2 string) {
	this.IsvCode = isvCode2
}
func (this *OapiCspaceAuthUpdateRequest) GetIsvCode() string {
	return this.IsvCode
}
func (this *OapiCspaceAuthUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.auth.update"
}
func (this *OapiCspaceAuthUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceAuthUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceAuthUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceAuthUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceAuthUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["duration"] = this.Duration
	txtParams["isv_code"] = this.IsvCode
	return txtParams
}
func (this *OapiCspaceAuthUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCspaceAuthUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceAuthUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceAuthUpdateResponse struct {
	taobao.TaobaoResponse
	Result  IsvAuthCodeResult `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
