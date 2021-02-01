package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsGetPersistentCodeRequest() *OapiSnsGetPersistentCodeRequest {
	return &OapiSnsGetPersistentCodeRequest{}
}

type OapiSnsGetPersistentCodeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsGetPersistentCodeResponse
	TmpAuthCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsGetPersistentCodeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsGetPersistentCodeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsGetPersistentCodeRequest) SetTmpAuthCode(tmpAuthCode2 string) {
	this.TmpAuthCode = tmpAuthCode2
}
func (this *OapiSnsGetPersistentCodeRequest) GetTmpAuthCode() string {
	return this.TmpAuthCode
}
func (this *OapiSnsGetPersistentCodeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.get_persistent_code"
}
func (this *OapiSnsGetPersistentCodeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsGetPersistentCodeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsGetPersistentCodeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsGetPersistentCodeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsGetPersistentCodeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["tmp_auth_code"] = this.TmpAuthCode
	return txtParams
}
func (this *OapiSnsGetPersistentCodeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsGetPersistentCodeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsGetPersistentCodeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsGetPersistentCodeResponse struct {
	taobao.TaobaoResponse

	Openid         string `json:"openid,omitempty"`
	PersistentCode string `json:"persistent_code,omitempty"`
	Unionid        string `json:"unionid,omitempty"`
}
