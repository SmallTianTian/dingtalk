package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasUserSubaccountAddRequest() *OapiImpaasUserSubaccountAddRequest {
	return &OapiImpaasUserSubaccountAddRequest{}
}

type OapiImpaasUserSubaccountAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasUserSubaccountAddResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasUserSubaccountAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasUserSubaccountAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasUserSubaccountAddRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasUserSubaccountAddRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasUserSubaccountAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.user.subaccount.add"
}
func (this *OapiImpaasUserSubaccountAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasUserSubaccountAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasUserSubaccountAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasUserSubaccountAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasUserSubaccountAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasUserSubaccountAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasUserSubaccountAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasUserSubaccountAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AddSubAccountReq struct {
	AdminaccountId string `json:"adminaccount_id,omitempty"`
	Channel        string `json:"channel,omitempty"`
	SubaccountId   string `json:"subaccount_id,omitempty"`
}
type OapiImpaasUserSubaccountAddResponse struct {
	taobao.TaobaoResponse
	Result  AddSubAccountResp `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type AddSubAccountResp struct {
	ImOpenid string `json:"im_openid,omitempty"`
}
