package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpMessageCorpconversationAsyncsendbycodeRequest() *CorpMessageCorpconversationAsyncsendbycodeRequest {
	return &CorpMessageCorpconversationAsyncsendbycodeRequest{}
}

type CorpMessageCorpconversationAsyncsendbycodeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpMessageCorpconversationAsyncsendbycodeResponse
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

func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetCode() string {
	return this.Code
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetMsgcontent(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetMsgcontentString(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetMsgcontent() string {
	return this.Msgcontent
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetToAllUser(toAllUser2 bool) {
	this.ToAllUser = toAllUser2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetToAllUser() bool {
	return this.ToAllUser
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetUserIdList(userIdList2 string) {
	this.UserIdList = userIdList2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetUserIdList() string {
	return this.UserIdList
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetApiMethodName() string {
	return "dingtalk.corp.message.corpconversation.asyncsendbycode"
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetTextParams() map[string]interface{} {
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
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpMessageCorpconversationAsyncsendbycodeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpMessageCorpconversationAsyncsendbycodeResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
