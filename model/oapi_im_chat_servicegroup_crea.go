package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupCreateRequest() *OapiImChatServicegroupCreateRequest {
	return &OapiImChatServicegroupCreateRequest{}
}

type OapiImChatServicegroupCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatServicegroupCreateResponse
	GroupUniqId     string
	OrgInnerGroup   bool
	OwnerUserid     string
	Title           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatServicegroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupCreateRequest) SetGroupUniqId(groupUniqId2 string) {
	this.GroupUniqId = groupUniqId2
}
func (this *OapiImChatServicegroupCreateRequest) GetGroupUniqId() string {
	return this.GroupUniqId
}
func (this *OapiImChatServicegroupCreateRequest) SetOrgInnerGroup(orgInnerGroup2 bool) {
	this.OrgInnerGroup = orgInnerGroup2
}
func (this *OapiImChatServicegroupCreateRequest) GetOrgInnerGroup() bool {
	return this.OrgInnerGroup
}
func (this *OapiImChatServicegroupCreateRequest) SetOwnerUserid(ownerUserid2 string) {
	this.OwnerUserid = ownerUserid2
}
func (this *OapiImChatServicegroupCreateRequest) GetOwnerUserid() string {
	return this.OwnerUserid
}
func (this *OapiImChatServicegroupCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiImChatServicegroupCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiImChatServicegroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.create"
}
func (this *OapiImChatServicegroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_uniq_id"] = this.GroupUniqId
	txtParams["org_inner_group"] = this.OrgInnerGroup
	txtParams["owner_userid"] = this.OwnerUserid
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiImChatServicegroupCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OwnerUserid, "ownerUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupCreateResponse struct {
	taobao.TaobaoResponse
	Result  ServiceGroupCreateResponse `json:"result,omitempty"`
	Success bool                       `json:"success,omitempty"`
}
type ServiceGroupCreateResponse struct {
	ChatId             string `json:"chat_id,omitempty"`
	OpenConversationId string `json:"open_conversation_id,omitempty"`
	Url                string `json:"url,omitempty"`
}
