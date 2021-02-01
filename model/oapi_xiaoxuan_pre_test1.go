package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiXiaoxuanPreTest1Request() *OapiXiaoxuanPreTest1Request {
	return &OapiXiaoxuanPreTest1Request{}
}

type OapiXiaoxuanPreTest1Request struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiXiaoxuanPreTest1Response
	Name            string
	SystemData      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiXiaoxuanPreTest1Request) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiXiaoxuanPreTest1Request) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiXiaoxuanPreTest1Request) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiXiaoxuanPreTest1Request) GetName() string {
	return this.Name
}
func (this *OapiXiaoxuanPreTest1Request) SetSystemData(systemData2 string) {
	this.SystemData = systemData2
}
func (this *OapiXiaoxuanPreTest1Request) GetSystemData() string {
	return this.SystemData
}
func (this *OapiXiaoxuanPreTest1Request) GetApiMethodName() string {
	return "dingtalk.oapi.xiaoxuan.pre.test1"
}
func (this *OapiXiaoxuanPreTest1Request) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiXiaoxuanPreTest1Request) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiXiaoxuanPreTest1Request) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiXiaoxuanPreTest1Request) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiXiaoxuanPreTest1Request) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["name"] = this.Name
	txtParams["systemData"] = this.SystemData
	return txtParams
}
func (this *OapiXiaoxuanPreTest1Request) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiXiaoxuanPreTest1Request) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiXiaoxuanPreTest1Request) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiXiaoxuanPreTest1Response struct {
	taobao.TaobaoResponse

	Name       string `json:"name,omitempty"`
	ResultData string `json:"resultData,omitempty"`
}
