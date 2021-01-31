package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasUserSubaccountDeleteRequest() *OapiImpaasUserSubaccountDeleteRequest {
	return &OapiImpaasUserSubaccountDeleteRequest{}
}

type OapiImpaasUserSubaccountDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasUserSubaccountDeleteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasUserSubaccountDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasUserSubaccountDeleteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.user.subaccount.delete"
}
func (this *OapiImpaasUserSubaccountDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasUserSubaccountDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasUserSubaccountDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasUserSubaccountDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasUserSubaccountDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DelSubAccountReq struct {
	AdminaccountId string `json:"adminaccount_id,omitempty"`
	Channel        string `json:"channel,omitempty"`
	SubaccountId   string `json:"subaccount_id,omitempty"`
}
type OapiImpaasUserSubaccountDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
