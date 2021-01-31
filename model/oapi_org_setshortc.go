package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiOrgSetshortcutRequest() *OapiOrgSetshortcutRequest {
	return &OapiOrgSetshortcutRequest{}
}

type OapiOrgSetshortcutRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgSetshortcutResponse
	AgentIds        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgSetshortcutRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgSetshortcutRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgSetshortcutRequest) SetAgentIds(agentIds2 string) {
	this.AgentIds = agentIds2
}
func (this *OapiOrgSetshortcutRequest) GetAgentIds() string {
	return this.AgentIds
}
func (this *OapiOrgSetshortcutRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.setshortcut"
}
func (this *OapiOrgSetshortcutRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgSetshortcutRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgSetshortcutRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgSetshortcutRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgSetshortcutRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentIds"] = this.AgentIds
	return txtParams
}
func (this *OapiOrgSetshortcutRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AgentIds, 20, "agentIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiOrgSetshortcutRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgSetshortcutRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgSetshortcutResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
