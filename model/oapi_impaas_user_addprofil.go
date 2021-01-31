package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasUserAddprofileRequest() *OapiImpaasUserAddprofileRequest {
	return &OapiImpaasUserAddprofileRequest{}
}

type OapiImpaasUserAddprofileRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasUserAddprofileResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasUserAddprofileRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasUserAddprofileRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasUserAddprofileRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasUserAddprofileRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasUserAddprofileRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.user.addprofile"
}
func (this *OapiImpaasUserAddprofileRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasUserAddprofileRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasUserAddprofileRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasUserAddprofileRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasUserAddprofileRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasUserAddprofileRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasUserAddprofileRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasUserAddprofileRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AddProfileReq struct {
	AvatarMediaid string `json:"avatar_mediaid,omitempty"`
	Channel       string `json:"channel,omitempty"`
	Extension     string `json:"extension,omitempty"`
	Id            string `json:"id,omitempty"`
	Nick          string `json:"nick,omitempty"`
}
type OapiImpaasUserAddprofileResponse struct {
	taobao.TaobaoResponse
	Result  AddProfileResp `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type AddProfileResp struct {
	ImOpenid string `json:"im_openid,omitempty"`
}
