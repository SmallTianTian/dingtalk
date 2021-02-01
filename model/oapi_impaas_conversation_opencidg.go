package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasConversationOpencidGetRequest() *OapiImpaasConversationOpencidGetRequest {
	return &OapiImpaasConversationOpencidGetRequest{}
}

type OapiImpaasConversationOpencidGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConversationOpencidGetResponse
	Model           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasConversationOpencidGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConversationOpencidGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConversationOpencidGetRequest) SetModel(model2 string) {
	this.Model = model2
}
func (this *OapiImpaasConversationOpencidGetRequest) GetModel() string {
	return this.Model
}
func (this *OapiImpaasConversationOpencidGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.conversation.opencid.get"
}
func (this *OapiImpaasConversationOpencidGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConversationOpencidGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConversationOpencidGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConversationOpencidGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConversationOpencidGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model"] = this.Model
	return txtParams
}
func (this *OapiImpaasConversationOpencidGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasConversationOpencidGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConversationOpencidGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CrossDomainBaseConversationModel struct {
	Cid              string `json:"cid,omitempty"`
	ConversationType int64  `json:"conversation_type,omitempty"`
}
type OapiImpaasConversationOpencidGetResponse struct {
	taobao.TaobaoResponse

	OpenConversationId string `json:"open_conversation_id,omitempty"`
}
