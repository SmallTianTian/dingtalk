package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappDelwithuseridRequest() *OapiMicroappDelwithuseridRequest {
	return &OapiMicroappDelwithuseridRequest{}
}

type OapiMicroappDelwithuseridRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappDelwithuseridResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiMicroappDelwithuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappDelwithuseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappDelwithuseridRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappDelwithuseridRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappDelwithuseridRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiMicroappDelwithuseridRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiMicroappDelwithuseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.delwithuserid"
}
func (this *OapiMicroappDelwithuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappDelwithuseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappDelwithuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappDelwithuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappDelwithuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiMicroappDelwithuseridRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappDelwithuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappDelwithuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappDelwithuseridResponse struct {
	taobao.TaobaoResponse
}
