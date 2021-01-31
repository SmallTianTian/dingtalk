package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGroupMsgSendRequest() *OapiEduGroupMsgSendRequest {
	return &OapiEduGroupMsgSendRequest{}
}

type OapiEduGroupMsgSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiEduGroupMsgSendResponse
	AgentId           int64
	BizId             string
	ClassId           int64
	ImageUrl          string
	ReceiveUseridList string
	Replace           string
	TemplateCode      string
	TopHttpMethod     string
	TopResponseType   string
	Userid            string
}

func (this *OapiEduGroupMsgSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGroupMsgSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGroupMsgSendRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiEduGroupMsgSendRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiEduGroupMsgSendRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiEduGroupMsgSendRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiEduGroupMsgSendRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduGroupMsgSendRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduGroupMsgSendRequest) SetImageUrl(imageUrl2 string) {
	this.ImageUrl = imageUrl2
}
func (this *OapiEduGroupMsgSendRequest) GetImageUrl() string {
	return this.ImageUrl
}
func (this *OapiEduGroupMsgSendRequest) SetReceiveUseridList(receiveUseridList2 string) {
	this.ReceiveUseridList = receiveUseridList2
}
func (this *OapiEduGroupMsgSendRequest) GetReceiveUseridList() string {
	return this.ReceiveUseridList
}
func (this *OapiEduGroupMsgSendRequest) SetReplace(replace2 string) {
	this.Replace = replace2
}
func (this *OapiEduGroupMsgSendRequest) GetReplace() string {
	return this.Replace
}
func (this *OapiEduGroupMsgSendRequest) SetTemplateCode(templateCode2 string) {
	this.TemplateCode = templateCode2
}
func (this *OapiEduGroupMsgSendRequest) GetTemplateCode() string {
	return this.TemplateCode
}
func (this *OapiEduGroupMsgSendRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduGroupMsgSendRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduGroupMsgSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.group.msg.send"
}
func (this *OapiEduGroupMsgSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGroupMsgSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGroupMsgSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGroupMsgSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGroupMsgSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["biz_id"] = this.BizId
	txtParams["class_id"] = this.ClassId
	txtParams["image_url"] = this.ImageUrl
	txtParams["receive_userid_list"] = this.ReceiveUseridList
	txtParams["replace"] = this.Replace
	txtParams["template_code"] = this.TemplateCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduGroupMsgSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGroupMsgSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGroupMsgSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGroupMsgSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
