package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappAddwithuseridRequest() *OapiMicroappAddwithuseridRequest {
	return &OapiMicroappAddwithuseridRequest{}
}

type OapiMicroappAddwithuseridRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappAddwithuseridResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiMicroappAddwithuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappAddwithuseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappAddwithuseridRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappAddwithuseridRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappAddwithuseridRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiMicroappAddwithuseridRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiMicroappAddwithuseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.addwithuserid"
}
func (this *OapiMicroappAddwithuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappAddwithuseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappAddwithuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappAddwithuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappAddwithuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiMicroappAddwithuseridRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappAddwithuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappAddwithuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappAddwithuseridResponse struct {
	taobao.TaobaoResponse
}
