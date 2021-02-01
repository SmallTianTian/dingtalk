package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasUserGetlogintokenRequest() *OapiImpaasUserGetlogintokenRequest {
	return &OapiImpaasUserGetlogintokenRequest{}
}

type OapiImpaasUserGetlogintokenRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasUserGetlogintokenResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasUserGetlogintokenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasUserGetlogintokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasUserGetlogintokenRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasUserGetlogintokenRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasUserGetlogintokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.user.getlogintoken"
}
func (this *OapiImpaasUserGetlogintokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasUserGetlogintokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasUserGetlogintokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasUserGetlogintokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasUserGetlogintokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasUserGetlogintokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasUserGetlogintokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasUserGetlogintokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetLoginTokenReq struct {
	Channel string `json:"channel,omitempty"`
	Id      string `json:"id,omitempty"`
}
type OapiImpaasUserGetlogintokenResponse struct {
	taobao.TaobaoResponse
	Result  GetLoginTokenResp `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type GetLoginTokenResp struct {
	Expire     int64  `json:"expire,omitempty"`
	LoginToken string `json:"login_token,omitempty"`
}
