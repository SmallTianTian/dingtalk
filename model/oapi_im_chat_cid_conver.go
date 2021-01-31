package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatCidConvertRequest() *OapiImChatCidConvertRequest {
	return &OapiImChatCidConvertRequest{}
}

type OapiImChatCidConvertRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatCidConvertResponse
	ChatId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatCidConvertRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatCidConvertRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatCidConvertRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatCidConvertRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatCidConvertRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.cid.convert"
}
func (this *OapiImChatCidConvertRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatCidConvertRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatCidConvertRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatCidConvertRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatCidConvertRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	return txtParams
}
func (this *OapiImChatCidConvertRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatCidConvertRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatCidConvertRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatCidConvertResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
