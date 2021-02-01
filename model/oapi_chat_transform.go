package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatTransformRequest() *OapiChatTransformRequest {
	return &OapiChatTransformRequest{}
}

type OapiChatTransformRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiChatTransformResponse
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiChatTransformRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatTransformRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatTransformRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiChatTransformRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiChatTransformRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.transform"
}
func (this *OapiChatTransformRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatTransformRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatTransformRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatTransformRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatTransformRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiChatTransformRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatTransformRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatTransformRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatTransformResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
