package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsMessageSystemaccountSendmessageRequest() *OapiAtsMessageSystemaccountSendmessageRequest {
	return &OapiAtsMessageSystemaccountSendmessageRequest{}
}

type OapiAtsMessageSystemaccountSendmessageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsMessageSystemaccountSendmessageResponse
	Content         string
	MessageBizCode  string
	Openid          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetContentString(content2 string) {
	this.Content = content2
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetContent() string {
	return this.Content
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetMessageBizCode(messageBizCode2 string) {
	this.MessageBizCode = messageBizCode2
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetMessageBizCode() string {
	return this.MessageBizCode
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetOpenid(openid2 string) {
	this.Openid = openid2
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetOpenid() string {
	return this.Openid
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.message.systemaccount.sendmessage"
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["message_biz_code"] = this.MessageBizCode
	txtParams["openid"] = this.Openid
	return txtParams
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Content, taobao.DATA_CONTENT); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsMessageSystemaccountSendmessageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsMessageSystemaccountSendmessageResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
