package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatBanwordsQueryRequest() *OapiChatBanwordsQueryRequest {
	return &OapiChatBanwordsQueryRequest{}
}

type OapiChatBanwordsQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatBanwordsQueryResponse
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiChatBanwordsQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatBanwordsQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatBanwordsQueryRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatBanwordsQueryRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatBanwordsQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiChatBanwordsQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiChatBanwordsQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.banwords.query"
}
func (this *OapiChatBanwordsQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatBanwordsQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatBanwordsQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatBanwordsQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatBanwordsQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiChatBanwordsQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatBanwordsQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatBanwordsQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatBanwordsQueryResponse struct {
	taobao.TaobaoResponse
	Result  OpenBanWordModel `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type UserBanWordModel struct {
	BanWordsStatus bool  `json:"ban_words_status,omitempty"`
	EndTime        int64 `json:"end_time,omitempty"`
	StartTime      int64 `json:"start_time,omitempty"`
}
type OpenBanWordModel struct {
	AllBanWords  bool             `json:"all_ban_words,omitempty"`
	UserBanWords UserBanWordModel `json:"user_ban_words,omitempty"`
}
