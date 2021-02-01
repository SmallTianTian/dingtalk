package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasConversationUpdateentranceidRequest() *OapiImpaasConversationUpdateentranceidRequest {
	return &OapiImpaasConversationUpdateentranceidRequest{}
}

type OapiImpaasConversationUpdateentranceidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConversationUpdateentranceidResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasConversationUpdateentranceidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConversationUpdateentranceidRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.conversation.updateentranceid"
}
func (this *OapiImpaasConversationUpdateentranceidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConversationUpdateentranceidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConversationUpdateentranceidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasConversationUpdateentranceidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConversationUpdateentranceidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AccountInfo struct {
	Channel string `json:"channel,omitempty"`
	Id      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
}
type UpdateEntranceIdRequest struct {
	Accounts   []AccountInfo `json:"accounts,omitempty"`
	Channel    string        `json:"channel,omitempty"`
	Cid        string        `json:"cid,omitempty"`
	EntranceId int64         `json:"entrance_id,omitempty"`
	Extension  string        `json:"extension,omitempty"`
	Uuid       string        `json:"uuid,omitempty"`
}
type OapiImpaasConversationUpdateentranceidResponse struct {
	taobao.TaobaoResponse
}
