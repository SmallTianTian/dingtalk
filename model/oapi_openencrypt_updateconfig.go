package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOpenencryptUpdateconfigRequest() *OapiOpenencryptUpdateconfigRequest {
	return &OapiOpenencryptUpdateconfigRequest{}
}

type OapiOpenencryptUpdateconfigRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiOpenencryptUpdateconfigResponse
	TopHttpMethod        string
	TopResourceKmsConfig string
	TopResponseType      string
}

func (this *OapiOpenencryptUpdateconfigRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOpenencryptUpdateconfigRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOpenencryptUpdateconfigRequest) SetTopResourceKmsConfig(topResourceKmsConfig2 string) {
	this.TopResourceKmsConfig = topResourceKmsConfig2
}
func (this *OapiOpenencryptUpdateconfigRequest) GetTopResourceKmsConfig() string {
	return this.TopResourceKmsConfig
}
func (this *OapiOpenencryptUpdateconfigRequest) GetApiMethodName() string {
	return "dingtalk.oapi.openencrypt.updateconfig"
}
func (this *OapiOpenencryptUpdateconfigRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOpenencryptUpdateconfigRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOpenencryptUpdateconfigRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOpenencryptUpdateconfigRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOpenencryptUpdateconfigRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_resource_kms_config"] = this.TopResourceKmsConfig
	return txtParams
}
func (this *OapiOpenencryptUpdateconfigRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOpenencryptUpdateconfigRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOpenencryptUpdateconfigRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopResourceKmsConfig struct {
	Agentid         int64  `json:"agentid,omitempty"`
	Extension       string `json:"extension,omitempty"`
	Requestid       string `json:"requestid,omitempty"`
	Resource        string `json:"resource,omitempty"`
	RotatePeriodDay int64  `json:"rotate_period_day,omitempty"`
}
type OapiOpenencryptUpdateconfigResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
