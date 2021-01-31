package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessDirlistGetRequest() *OapiProcessDirlistGetRequest {
	return &OapiProcessDirlistGetRequest{}
}

type OapiProcessDirlistGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessDirlistGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessDirlistGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessDirlistGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessDirlistGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessDirlistGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessDirlistGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.dirlist.get"
}
func (this *OapiProcessDirlistGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessDirlistGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessDirlistGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessDirlistGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessDirlistGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessDirlistGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessDirlistGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessDirlistGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BaseRequest struct {
	Agentid int64 `json:"agentid,omitempty"`
}
type OapiProcessDirlistGetResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
