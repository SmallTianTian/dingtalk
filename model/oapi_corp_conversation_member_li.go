package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCorpConversationMemberListRequest() *OapiCorpConversationMemberListRequest {
	return &OapiCorpConversationMemberListRequest{}
}

type OapiCorpConversationMemberListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCorpConversationMemberListResponse
	ChatId          string
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCorpConversationMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCorpConversationMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCorpConversationMemberListRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiCorpConversationMemberListRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiCorpConversationMemberListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiCorpConversationMemberListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiCorpConversationMemberListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCorpConversationMemberListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCorpConversationMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.corp.conversation.member.list"
}
func (this *OapiCorpConversationMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCorpConversationMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCorpConversationMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCorpConversationMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCorpConversationMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiCorpConversationMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCorpConversationMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCorpConversationMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCorpConversationMemberListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenMemberModel struct {
	Userid string `json:"userid,omitempty"`
}
type OpenMemberListModel struct {
	MemberList []OpenMemberModel `json:"member_list,omitempty"`
}
