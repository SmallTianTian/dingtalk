package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripProjectModifyRequest() *OapiAlitripBtripProjectModifyRequest {
	return &OapiAlitripBtripProjectModifyRequest{}
}

type OapiAlitripBtripProjectModifyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripProjectModifyResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripProjectModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripProjectModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripProjectModifyRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripProjectModifyRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripProjectModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.project.modify"
}
func (this *OapiAlitripBtripProjectModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripProjectModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripProjectModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripProjectModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripProjectModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripProjectModifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripProjectModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripProjectModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripProjectModifyResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Module  bool   `json:"module,omitempty"`
	Success bool   `json:"success,omitempty"`
}
