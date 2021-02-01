package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatUpdatebanwordsRequest() *OapiChatUpdatebanwordsRequest {
	return &OapiChatUpdatebanwordsRequest{}
}

type OapiChatUpdatebanwordsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatUpdatebanwordsResponse
	BanWordsTime    int64
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
	Type            int64
	UseridList      string
}

func (this *OapiChatUpdatebanwordsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatUpdatebanwordsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatUpdatebanwordsRequest) SetBanWordsTime(banWordsTime2 int64) {
	this.BanWordsTime = banWordsTime2
}
func (this *OapiChatUpdatebanwordsRequest) GetBanWordsTime() int64 {
	return this.BanWordsTime
}
func (this *OapiChatUpdatebanwordsRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatUpdatebanwordsRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatUpdatebanwordsRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiChatUpdatebanwordsRequest) GetType() int64 {
	return this.Type
}
func (this *OapiChatUpdatebanwordsRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiChatUpdatebanwordsRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiChatUpdatebanwordsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.updatebanwords"
}
func (this *OapiChatUpdatebanwordsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatUpdatebanwordsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatUpdatebanwordsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatUpdatebanwordsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatUpdatebanwordsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ban_words_time"] = this.BanWordsTime
	txtParams["chatid"] = this.Chatid
	txtParams["type"] = this.Type
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiChatUpdatebanwordsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BanWordsTime, "banWordsTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatUpdatebanwordsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatUpdatebanwordsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatUpdatebanwordsResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
