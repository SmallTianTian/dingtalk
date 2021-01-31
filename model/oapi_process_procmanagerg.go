package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessProcmanagerGetRequest() *OapiProcessProcmanagerGetRequest {
	return &OapiProcessProcmanagerGetRequest{}
}

type OapiProcessProcmanagerGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessProcmanagerGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessProcmanagerGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessProcmanagerGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessProcmanagerGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessProcmanagerGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessProcmanagerGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.procmanager.get"
}
func (this *OapiProcessProcmanagerGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessProcmanagerGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessProcmanagerGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessProcmanagerGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessProcmanagerGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessProcmanagerGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessProcmanagerGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessProcmanagerGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BaseProcessRequest struct {
	Agentid     int64  `json:"agentid,omitempty"`
	ProcessCode string `json:"process_code,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
type OapiProcessProcmanagerGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Result  []string `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
