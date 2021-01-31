package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTasklistHiddenCancelRequest() *OapiWorkspaceTasklistHiddenCancelRequest {
	return &OapiWorkspaceTasklistHiddenCancelRequest{}
}

type OapiWorkspaceTasklistHiddenCancelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTasklistHiddenCancelResponse
	Agentid         int64
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.tasklist.hidden.cancel"
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTasklistHiddenCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTasklistHiddenCancelResponse struct {
	taobao.TaobaoResponse
}
