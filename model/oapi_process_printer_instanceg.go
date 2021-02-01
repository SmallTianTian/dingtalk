package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessPrinterInstanceGetRequest() *OapiProcessPrinterInstanceGetRequest {
	return &OapiProcessPrinterInstanceGetRequest{}
}

type OapiProcessPrinterInstanceGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessPrinterInstanceGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessPrinterInstanceGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessPrinterInstanceGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessPrinterInstanceGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessPrinterInstanceGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessPrinterInstanceGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.printer.instance.get"
}
func (this *OapiProcessPrinterInstanceGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessPrinterInstanceGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessPrinterInstanceGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessPrinterInstanceGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessPrinterInstanceGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessPrinterInstanceGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessPrinterInstanceGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessPrinterInstanceGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type InstancePrintDataRequest struct {
	Agentid     int64  `json:"agentid,omitempty"`
	InstanceId  string `json:"instance_id,omitempty"`
	ProcessCode string `json:"process_code,omitempty"`
}
type OapiProcessPrinterInstanceGetResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
