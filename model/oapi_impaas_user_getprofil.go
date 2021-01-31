package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasUserGetprofileRequest() *OapiImpaasUserGetprofileRequest {
	return &OapiImpaasUserGetprofileRequest{}
}

type OapiImpaasUserGetprofileRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasUserGetprofileResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasUserGetprofileRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasUserGetprofileRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasUserGetprofileRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasUserGetprofileRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasUserGetprofileRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.user.getprofile"
}
func (this *OapiImpaasUserGetprofileRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasUserGetprofileRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasUserGetprofileRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasUserGetprofileRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasUserGetprofileRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasUserGetprofileRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasUserGetprofileRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasUserGetprofileRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetProfileReq struct {
	AccountInfo AccountInfo `json:"accountInfo,omitempty"`
}
type OapiImpaasUserGetprofileResponse struct {
	taobao.TaobaoResponse
	Result  GetProfileResp `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type GetProfileResp struct {
	AppUserid     string `json:"app_userid,omitempty"`
	AvatarMediaid string `json:"avatar_mediaid,omitempty"`
	Channel       string `json:"channel,omitempty"`
	Extension     string `json:"extension,omitempty"`
	ImOpenid      string `json:"im_openid,omitempty"`
	Nick          string `json:"nick,omitempty"`
	Status        int64  `json:"status,omitempty"`
}
