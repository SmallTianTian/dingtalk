package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingmiGroupGetRequest() *OapiDingmiGroupGetRequest {
	return &OapiDingmiGroupGetRequest{}
}

type OapiDingmiGroupGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiGroupGetResponse
	ConversationId  string
	Date            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingmiGroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiGroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiGroupGetRequest) SetConversationId(conversationId2 string) {
	this.ConversationId = conversationId2
}
func (this *OapiDingmiGroupGetRequest) GetConversationId() string {
	return this.ConversationId
}
func (this *OapiDingmiGroupGetRequest) SetDate(date2 string) {
	this.Date = date2
}
func (this *OapiDingmiGroupGetRequest) GetDate() string {
	return this.Date
}
func (this *OapiDingmiGroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.group.get"
}
func (this *OapiDingmiGroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiGroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiGroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiGroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiGroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conversation_id"] = this.ConversationId
	txtParams["date"] = this.Date
	return txtParams
}
func (this *OapiDingmiGroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConversationId, "conversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingmiGroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiGroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingmiGroupGetResponse struct {
	taobao.TaobaoResponse
	Result  GroupChatDataResult `json:"result,omitempty"`
	Success string              `json:"success,omitempty"`
}
type GroupMemberChatData struct {
	ChatCnt       string `json:"chat_cnt,omitempty"`
	CustomFlag    string `json:"custom_flag,omitempty"`
	SenderStaffid string `json:"sender_staffid,omitempty"`
}
type GroupChatDataResult struct {
	FromCache bool                  `json:"from_cache,omitempty"`
	IsBlock   bool                  `json:"is_block,omitempty"`
	Rawset    []GroupMemberChatData `json:"rawset,omitempty"`
	Runtime   int64                 `json:"runtime,omitempty"`
}
