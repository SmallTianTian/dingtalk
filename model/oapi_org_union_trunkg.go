package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgUnionTrunkGetRequest() *OapiOrgUnionTrunkGetRequest {
	return &OapiOrgUnionTrunkGetRequest{}
}

type OapiOrgUnionTrunkGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgUnionTrunkGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgUnionTrunkGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgUnionTrunkGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgUnionTrunkGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.union.trunk.get"
}
func (this *OapiOrgUnionTrunkGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgUnionTrunkGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgUnionTrunkGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgUnionTrunkGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgUnionTrunkGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiOrgUnionTrunkGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgUnionTrunkGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgUnionTrunkGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgUnionTrunkGetResponse struct {
	taobao.TaobaoResponse
	Result  []OpenOrgUnion `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
