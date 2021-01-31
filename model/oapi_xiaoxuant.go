package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiXiaoxuanTestRequest() *OapiXiaoxuanTestRequest {
	return &OapiXiaoxuanTestRequest{}
}

type OapiXiaoxuanTestRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiXiaoxuanTestResponse
	NormalData      string
	SystemData      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiXiaoxuanTestRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiXiaoxuanTestRequest) SetNormalData(normalData2 string) {
	this.NormalData = normalData2
}
func (this *OapiXiaoxuanTestRequest) GetNormalData() string {
	return this.NormalData
}
func (this *OapiXiaoxuanTestRequest) SetSystemData(systemData2 string) {
	this.SystemData = systemData2
}
func (this *OapiXiaoxuanTestRequest) GetSystemData() string {
	return this.SystemData
}
func (this *OapiXiaoxuanTestRequest) GetApiMethodName() string {
	return "dingtalk.oapi.xiaoxuan.test"
}
func (this *OapiXiaoxuanTestRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *OapiXiaoxuanTestRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiXiaoxuanTestRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiXiaoxuanTestRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiXiaoxuanTestRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiXiaoxuanTestRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["normal_data"] = this.NormalData
	txtParams["system_data"] = this.SystemData
	return txtParams
}
func (this *OapiXiaoxuanTestRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiXiaoxuanTestRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiXiaoxuanTestRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiXiaoxuanTestResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
