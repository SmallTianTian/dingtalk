package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatControlgroupCreateRequest() *OapiImChatControlgroupCreateRequest {
	return &OapiImChatControlgroupCreateRequest{}
}

type OapiImChatControlgroupCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatControlgroupCreateResponse
	AuthorityType   int64
	GroupType       string
	GroupUniqId     string
	MemberUserids   string
	OwnerUserid     string
	Title           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatControlgroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatControlgroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatControlgroupCreateRequest) SetAuthorityType(authorityType2 int64) {
	this.AuthorityType = authorityType2
}
func (this *OapiImChatControlgroupCreateRequest) GetAuthorityType() int64 {
	return this.AuthorityType
}
func (this *OapiImChatControlgroupCreateRequest) SetGroupType(groupType2 string) {
	this.GroupType = groupType2
}
func (this *OapiImChatControlgroupCreateRequest) GetGroupType() string {
	return this.GroupType
}
func (this *OapiImChatControlgroupCreateRequest) SetGroupUniqId(groupUniqId2 string) {
	this.GroupUniqId = groupUniqId2
}
func (this *OapiImChatControlgroupCreateRequest) GetGroupUniqId() string {
	return this.GroupUniqId
}
func (this *OapiImChatControlgroupCreateRequest) SetMemberUserids(memberUserids2 string) {
	this.MemberUserids = memberUserids2
}
func (this *OapiImChatControlgroupCreateRequest) GetMemberUserids() string {
	return this.MemberUserids
}
func (this *OapiImChatControlgroupCreateRequest) SetOwnerUserid(ownerUserid2 string) {
	this.OwnerUserid = ownerUserid2
}
func (this *OapiImChatControlgroupCreateRequest) GetOwnerUserid() string {
	return this.OwnerUserid
}
func (this *OapiImChatControlgroupCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiImChatControlgroupCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiImChatControlgroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.controlgroup.create"
}
func (this *OapiImChatControlgroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatControlgroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatControlgroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatControlgroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatControlgroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["authority_type"] = this.AuthorityType
	txtParams["group_type"] = this.GroupType
	txtParams["group_uniq_id"] = this.GroupUniqId
	txtParams["member_userids"] = this.MemberUserids
	txtParams["owner_userid"] = this.OwnerUserid
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiImChatControlgroupCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MemberUserids, "memberUserids"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatControlgroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatControlgroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatControlgroupCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64                      `json:"errcode,omitempty"`
	Errmsg  string                     `json:"errmsg,omitempty"`
	Result  ControlGroupCreateResponse `json:"result,omitempty"`
	Success bool                       `json:"success,omitempty"`
}
type ControlGroupCreateResponse struct {
	ChatId             string `json:"chat_id,omitempty"`
	OpenConversationId string `json:"open_conversation_id,omitempty"`
}
