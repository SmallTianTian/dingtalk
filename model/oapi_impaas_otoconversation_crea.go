package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasOtoconversationCreateRequest() *OapiImpaasOtoconversationCreateRequest {
	return &OapiImpaasOtoconversationCreateRequest{}
}

type OapiImpaasOtoconversationCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasOtoconversationCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasOtoconversationCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasOtoconversationCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasOtoconversationCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasOtoconversationCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasOtoconversationCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.otoconversation.create"
}
func (this *OapiImpaasOtoconversationCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasOtoconversationCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasOtoconversationCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasOtoconversationCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasOtoconversationCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasOtoconversationCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasOtoconversationCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasOtoconversationCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CreateO2OSubConversationRequest struct {
	AccountInfoList []AccountInfo `json:"account_info_list,omitempty"`
	Channel         string        `json:"channel,omitempty"`
	EntranceIdList  []int64       `json:"entrance_id_list,omitempty"`
	Extension       string        `json:"extension,omitempty"`
	Uuid            string        `json:"uuid,omitempty"`
}
type OapiImpaasOtoconversationCreateResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
