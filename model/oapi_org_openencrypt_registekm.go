package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgOpenencryptRegistekmsRequest() *OapiOrgOpenencryptRegistekmsRequest {
	return &OapiOrgOpenencryptRegistekmsRequest{}
}

type OapiOrgOpenencryptRegistekmsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgOpenencryptRegistekmsResponse
	TopHttpMethod   string
	TopKmsMeta      string
	TopResponseType string
}

func (this *OapiOrgOpenencryptRegistekmsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgOpenencryptRegistekmsRequest) SetTopKmsMeta(topKmsMeta2 string) {
	this.TopKmsMeta = topKmsMeta2
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetTopKmsMeta() string {
	return this.TopKmsMeta
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.openencrypt.registekms"
}
func (this *OapiOrgOpenencryptRegistekmsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgOpenencryptRegistekmsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgOpenencryptRegistekmsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_kms_meta"] = this.TopKmsMeta
	return txtParams
}
func (this *OapiOrgOpenencryptRegistekmsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgOpenencryptRegistekmsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgOpenencryptRegistekmsResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
