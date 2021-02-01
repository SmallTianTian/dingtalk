package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacV2DatavVideoconfGetRequest() *OapiKacV2DatavVideoconfGetRequest {
	return &OapiKacV2DatavVideoconfGetRequest{}
}

type OapiKacV2DatavVideoconfGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacV2DatavVideoconfGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacV2DatavVideoconfGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacV2DatavVideoconfGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.v2.datav.videoconf.get"
}
func (this *OapiKacV2DatavVideoconfGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacV2DatavVideoconfGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacV2DatavVideoconfGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacV2DatavVideoconfGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacV2DatavVideoconfGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacV2DatavVideoconfGetResponse struct {
	taobao.TaobaoResponse
	Result McsSummaryResponse `json:"result,omitempty"`
}
