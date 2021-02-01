package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatQrcodeGetRequest() *OapiChatQrcodeGetRequest {
	return &OapiChatQrcodeGetRequest{}
}

type OapiChatQrcodeGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatQrcodeGetResponse
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiChatQrcodeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatQrcodeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatQrcodeGetRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatQrcodeGetRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatQrcodeGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiChatQrcodeGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiChatQrcodeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.qrcode.get"
}
func (this *OapiChatQrcodeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatQrcodeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatQrcodeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatQrcodeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatQrcodeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiChatQrcodeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatQrcodeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatQrcodeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatQrcodeGetResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
