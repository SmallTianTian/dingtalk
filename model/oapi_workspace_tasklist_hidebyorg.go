package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTasklistHidebyorgRequest() *OapiWorkspaceTasklistHidebyorgRequest {
	return &OapiWorkspaceTasklistHidebyorgRequest{}
}

type OapiWorkspaceTasklistHidebyorgRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTasklistHidebyorgResponse
	MicroappAgentId int64
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiWorkspaceTasklistHidebyorgRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.tasklist.hidebyorg"
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTasklistHidebyorgRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTasklistHidebyorgResponse struct {
	taobao.TaobaoResponse
}
