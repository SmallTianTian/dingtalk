package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatChatidTransformqrcodeGetRequest() *OapiChatChatidTransformqrcodeGetRequest {
	return &OapiChatChatidTransformqrcodeGetRequest{}
}

type OapiChatChatidTransformqrcodeGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatChatidTransformqrcodeGetResponse
	GroupUrl        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatChatidTransformqrcodeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatChatidTransformqrcodeGetRequest) SetGroupUrl(groupUrl2 string) {
	this.GroupUrl = groupUrl2
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetGroupUrl() string {
	return this.GroupUrl
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.chatid.transformqrcode.get"
}
func (this *OapiChatChatidTransformqrcodeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatChatidTransformqrcodeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatChatidTransformqrcodeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_url"] = this.GroupUrl
	return txtParams
}
func (this *OapiChatChatidTransformqrcodeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupUrl, "groupUrl"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatChatidTransformqrcodeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatChatidTransformqrcodeGetResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
