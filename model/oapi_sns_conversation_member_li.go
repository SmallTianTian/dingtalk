package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSnsConversationMemberListRequest() *OapiSnsConversationMemberListRequest {
	return &OapiSnsConversationMemberListRequest{}
}

type OapiSnsConversationMemberListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiSnsConversationMemberListResponse
	Cursor             string
	OpenConversationId string
	Size               int64
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiSnsConversationMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsConversationMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsConversationMemberListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiSnsConversationMemberListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiSnsConversationMemberListRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiSnsConversationMemberListRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiSnsConversationMemberListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSnsConversationMemberListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSnsConversationMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.conversation.member.list"
}
func (this *OapiSnsConversationMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsConversationMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsConversationMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsConversationMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsConversationMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSnsConversationMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSnsConversationMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsConversationMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsConversationMemberListResponse struct {
	taobao.TaobaoResponse
	Result  SnsOpenGroupMemberResponse `json:"result,omitempty"`
	Success bool                       `json:"success,omitempty"`
}
type SnsGroupMemberModel struct {
	GroupNickName string `json:"group_nick_name,omitempty"`
	NickName      string `json:"nick_name,omitempty"`
	PortraitUrl   string `json:"portrait_url,omitempty"`
	Role          int64  `json:"role,omitempty"`
	Unionid       string `json:"unionid,omitempty"`
}
type SnsOpenGroupMemberResponse struct {
	HasMore    bool                  `json:"has_more,omitempty"`
	Members    []SnsGroupMemberModel `json:"members,omitempty"`
	NextCursor string                `json:"next_cursor,omitempty"`
}
