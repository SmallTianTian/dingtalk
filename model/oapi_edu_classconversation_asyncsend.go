package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduClassconversationAsyncsendRequest() *OapiEduClassconversationAsyncsendRequest {
	return &OapiEduClassconversationAsyncsendRequest{}
}

type OapiEduClassconversationAsyncsendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassconversationAsyncsendResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduClassconversationAsyncsendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassconversationAsyncsendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassconversationAsyncsendRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiEduClassconversationAsyncsendRequest) GetRequest() string {
	return this.Request
}
func (this *OapiEduClassconversationAsyncsendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.classconversation.asyncsend"
}
func (this *OapiEduClassconversationAsyncsendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassconversationAsyncsendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassconversationAsyncsendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassconversationAsyncsendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassconversationAsyncsendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiEduClassconversationAsyncsendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduClassconversationAsyncsendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassconversationAsyncsendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopSendConversationMsgRequest struct {
	Attributes      string   `json:"attributes,omitempty"`
	ClassId         int64    `json:"class_id,omitempty"`
	Nonce           string   `json:"nonce,omitempty"`
	ReceiverUserIds []string `json:"receiver_user_ids,omitempty"`
	TemplateId      int64    `json:"template_id,omitempty"`
}
type OapiEduClassconversationAsyncsendResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type TopSendConversationMsgResponse struct {
	TaskId string `json:"task_id,omitempty"`
}
