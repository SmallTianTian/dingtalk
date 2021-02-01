package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserTokenGetRequest() *OapiUserTokenGetRequest {
	return &OapiUserTokenGetRequest{}
}

type OapiUserTokenGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserTokenGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiUserTokenGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserTokenGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserTokenGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiUserTokenGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiUserTokenGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.token.get"
}
func (this *OapiUserTokenGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserTokenGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserTokenGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserTokenGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserTokenGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiUserTokenGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserTokenGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserTokenGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserTokenGetResponse struct {
	taobao.TaobaoResponse
	Result UserToken4JsapiResponse `json:"result,omitempty"`
}
type UserToken4JsapiResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpireIn    int64  `json:"expire_in,omitempty"`
}
