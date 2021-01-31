package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskCleanRequest() *OapiWorkspaceTaskCleanRequest {
	return &OapiWorkspaceTaskCleanRequest{}
}

type OapiWorkspaceTaskCleanRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskCleanResponse
	Agentid         int64
	CorpId          string
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskCleanRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskCleanRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskCleanRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWorkspaceTaskCleanRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWorkspaceTaskCleanRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiWorkspaceTaskCleanRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiWorkspaceTaskCleanRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTaskCleanRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTaskCleanRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.clean"
}
func (this *OapiWorkspaceTaskCleanRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskCleanRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskCleanRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskCleanRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskCleanRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["corp_id"] = this.CorpId
	txtParams["operator_userid"] = this.OperatorUserid
	return txtParams
}
func (this *OapiWorkspaceTaskCleanRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskCleanRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskCleanRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskCleanResponse struct {
	taobao.TaobaoResponse
}
