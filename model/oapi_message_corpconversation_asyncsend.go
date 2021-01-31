package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationAsyncsendRequest() *OapiMessageCorpconversationAsyncsendRequest {
	return &OapiMessageCorpconversationAsyncsendRequest{}
}

type OapiMessageCorpconversationAsyncsendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationAsyncsendResponse
	AgentId         int64
	DeptIdList      string
	Msgcontent      string
	Msgtype         string
	ToAllUser       bool
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiMessageCorpconversationAsyncsendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetMsgcontent(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetMsgcontentString(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetMsgcontent() string {
	return this.Msgcontent
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetToAllUser(toAllUser2 bool) {
	this.ToAllUser = toAllUser2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetToAllUser() bool {
	return this.ToAllUser
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.asyncsend"
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationAsyncsendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["msgcontent"] = this.Msgcontent
	txtParams["msgtype"] = this.Msgtype
	txtParams["to_all_user"] = this.ToAllUser
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiMessageCorpconversationAsyncsendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationAsyncsendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationAsyncsendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	TaskId  int64  `json:"task_id,omitempty"`
}
