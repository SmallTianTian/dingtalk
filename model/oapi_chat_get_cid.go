package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatGetCidRequest() *OapiChatGetCidRequest {
	return &OapiChatGetCidRequest{}
}

type OapiChatGetCidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatGetCidResponse
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatGetCidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatGetCidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatGetCidRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatGetCidRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatGetCidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.get.cid"
}
func (this *OapiChatGetCidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatGetCidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatGetCidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatGetCidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatGetCidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	return txtParams
}
func (this *OapiChatGetCidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatGetCidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatGetCidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatGetCidResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
