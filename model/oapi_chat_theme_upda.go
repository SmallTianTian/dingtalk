package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatThemeUpdateRequest() *OapiChatThemeUpdateRequest {
	return &OapiChatThemeUpdateRequest{}
}

type OapiChatThemeUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatThemeUpdateResponse
	Chatid          string
	Mediaid         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatThemeUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatThemeUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatThemeUpdateRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatThemeUpdateRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatThemeUpdateRequest) SetMediaid(mediaid2 string) {
	this.Mediaid = mediaid2
}
func (this *OapiChatThemeUpdateRequest) GetMediaid() string {
	return this.Mediaid
}
func (this *OapiChatThemeUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.theme.update"
}
func (this *OapiChatThemeUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatThemeUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatThemeUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatThemeUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatThemeUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["mediaid"] = this.Mediaid
	return txtParams
}
func (this *OapiChatThemeUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatThemeUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatThemeUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatThemeUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
