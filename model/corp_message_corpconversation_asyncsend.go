package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpMessageCorpconversationAsyncsendRequest() *CorpMessageCorpconversationAsyncsendRequest {
	return &CorpMessageCorpconversationAsyncsendRequest{}
}

type CorpMessageCorpconversationAsyncsendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpMessageCorpconversationAsyncsendResponse
	AgentId         int64
	DeptIdList      string
	Msgcontent      string
	Msgtype         string
	ToAllUser       bool
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *CorpMessageCorpconversationAsyncsendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetMsgcontent(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetMsgcontentString(msgcontent2 string) {
	this.Msgcontent = msgcontent2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetMsgcontent() string {
	return this.Msgcontent
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetToAllUser(toAllUser2 bool) {
	this.ToAllUser = toAllUser2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetToAllUser() bool {
	return this.ToAllUser
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetUseridList() string {
	return this.UseridList
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetApiMethodName() string {
	return "dingtalk.corp.message.corpconversation.asyncsend"
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpMessageCorpconversationAsyncsendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["msgcontent"] = this.Msgcontent
	txtParams["msgtype"] = this.Msgtype
	txtParams["to_all_user"] = this.ToAllUser
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *CorpMessageCorpconversationAsyncsendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpMessageCorpconversationAsyncsendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpMessageCorpconversationAsyncsendResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
