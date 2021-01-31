package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationAsyncsendbycodeRequest() *OapiMessageCorpconversationAsyncsendbycodeRequest {
	return &OapiMessageCorpconversationAsyncsendbycodeRequest{}
}

type OapiMessageCorpconversationAsyncsendbycodeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationAsyncsendbycodeResponse
	AgentId         int64
	Code            string
	DeptIdList      string
	Msgcontent      string
	Msgtype         string
	ToAllUser       bool
	TopHttpMethod   string
	TopResponseType string
	UserIdList      string
}

func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetCode() string {
	return this.Code
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetMsgcontent(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetMsgcontentString(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetMsgcontent() string {
	return this.Msgcontent
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetToAllUser(toAllUser2 bool) {
	this.ToAllUser = toAllUser2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetToAllUser() bool {
	return this.ToAllUser
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetUserIdList(userIdList2 string) {
	this.UserIdList = userIdList2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetUserIdList() string {
	return this.UserIdList
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.asyncsendbycode"
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams[taobao.ERROR_CODE] = this.Code
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["msgcontent"] = this.Msgcontent
	txtParams["msgtype"] = this.Msgtype
	txtParams["to_all_user"] = this.ToAllUser
	txtParams["user_id_list"] = this.UserIdList
	return txtParams
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationAsyncsendbycodeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationAsyncsendbycodeResponse struct {
	taobao.TaobaoResponse
	TaskId int64 `json:"task_id,omitempty"`
}
