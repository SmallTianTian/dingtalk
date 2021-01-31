package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMpdevAccesskeyGetRequest() *OapiMpdevAccesskeyGetRequest {
	return &OapiMpdevAccesskeyGetRequest{}
}

type OapiMpdevAccesskeyGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMpdevAccesskeyGetResponse
	MiniappId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMpdevAccesskeyGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMpdevAccesskeyGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMpdevAccesskeyGetRequest) SetMiniappId(miniappId2 string) {
	this.MiniappId = miniappId2
}
func (this *OapiMpdevAccesskeyGetRequest) GetMiniappId() string {
	return this.MiniappId
}
func (this *OapiMpdevAccesskeyGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.mpdev.accesskey.get"
}
func (this *OapiMpdevAccesskeyGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMpdevAccesskeyGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMpdevAccesskeyGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMpdevAccesskeyGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMpdevAccesskeyGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["miniapp_id"] = this.MiniappId
	return txtParams
}
func (this *OapiMpdevAccesskeyGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MiniappId, "miniappId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMpdevAccesskeyGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMpdevAccesskeyGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMpdevAccesskeyGetResponse struct {
	taobao.TaobaoResponse
	Result  OssTokenVo `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type OssTokenVo struct {
	AccessKeySecret string `json:"access_key_secret,omitempty"`
	Accessid        string `json:"accessid,omitempty"`
	Expiration      string `json:"expiration,omitempty"`
	Name            string `json:"name,omitempty"`
	SecurityToken   string `json:"security_token,omitempty"`
}
