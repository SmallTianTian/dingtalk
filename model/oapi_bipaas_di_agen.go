package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiBipaasDiAgentRequest() *OapiBipaasDiAgentRequest {
	return &OapiBipaasDiAgentRequest{}
}

type OapiBipaasDiAgentRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBipaasDiAgentResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBipaasDiAgentRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBipaasDiAgentRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBipaasDiAgentRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiBipaasDiAgentRequest) GetRequest() string {
	return this.Request
}
func (this *OapiBipaasDiAgentRequest) GetApiMethodName() string {
	return "dingtalk.oapi.bipaas.di.agent"
}
func (this *OapiBipaasDiAgentRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBipaasDiAgentRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBipaasDiAgentRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBipaasDiAgentRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBipaasDiAgentRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiBipaasDiAgentRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiBipaasDiAgentRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBipaasDiAgentRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BaseAgentRequest struct {
	Params string `json:"params,omitempty"`
	Path   string `json:"path,omitempty"`
}
type OapiBipaasDiAgentResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
