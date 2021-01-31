package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupMemberQueryRequest() *OapiImChatServicegroupMemberQueryRequest {
	return &OapiImChatServicegroupMemberQueryRequest{}
}

type OapiImChatServicegroupMemberQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatServicegroupMemberQueryResponse
	ChatId          string
	IncludeOwner    int64
	PageNum         int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatServicegroupMemberQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetIncludeOwner(includeOwner2 int64) {
	this.IncludeOwner = includeOwner2
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetIncludeOwner() int64 {
	return this.IncludeOwner
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetPageNum(pageNum2 int64) {
	this.PageNum = pageNum2
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetPageNum() int64 {
	return this.PageNum
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.member.query"
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupMemberQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	txtParams["include_owner"] = this.IncludeOwner
	txtParams["page_num"] = this.PageNum
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiImChatServicegroupMemberQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupMemberQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupMemberQueryResponse struct {
	taobao.TaobaoResponse
	Result  ServiceGroupMemberQueryResponse `json:"result,omitempty"`
	Success bool                            `json:"success,omitempty"`
}
type Members struct {
	DingtalkId    string `json:"dingtalk_id,omitempty"`
	GroupNickName string `json:"group_nick_name,omitempty"`
	NickName      string `json:"nick_name,omitempty"`
	Role          string `json:"role,omitempty"`
	Userid        string `json:"userid,omitempty"`
}
type ServiceGroupMemberQueryResponse struct {
	Members    []Members `json:"members,omitempty"`
	TotalCount int64     `json:"total_count,omitempty"`
}
