package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationAsyncsendV2Request() *OapiMessageCorpconversationAsyncsendV2Request {
	return &OapiMessageCorpconversationAsyncsendV2Request{}
}

type OapiMessageCorpconversationAsyncsendV2Request struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationAsyncsendV2Response
	AgentId         int64
	DeptIdList      string
	Msg             string
	ToAllUser       bool
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetMsg(msg2 string) {
	this.Msg = msg2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetMsg() string {
	return this.Msg
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetToAllUser(toAllUser2 bool) {
	this.ToAllUser = toAllUser2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetToAllUser() bool {
	return this.ToAllUser
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetUseridList() string {
	return this.UseridList
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.asyncsend_v2"
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["msg"] = this.Msg
	txtParams["to_all_user"] = this.ToAllUser
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationAsyncsendV2Request) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type StatusBar struct {
	StatusBg    string `json:"status_bg,omitempty"`
	StatusValue string `json:"status_value,omitempty"`
}
type OA struct {
	Body         Body      `json:"body,omitempty"`
	Head         Head      `json:"head,omitempty"`
	MessageUrl   string    `json:"message_url,omitempty"`
	PcMessageUrl string    `json:"pc_message_url,omitempty"`
	StatusBar    StatusBar `json:"status_bar,omitempty"`
}
type OapiMessageCorpconversationAsyncsendV2Response struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	TaskId  int64  `json:"task_id,omitempty"`
}
