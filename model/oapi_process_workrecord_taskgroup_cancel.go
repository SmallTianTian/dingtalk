package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordTaskgroupCancelRequest() *OapiProcessWorkrecordTaskgroupCancelRequest {
	return &OapiProcessWorkrecordTaskgroupCancelRequest{}
}

type OapiProcessWorkrecordTaskgroupCancelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordTaskgroupCancelResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.taskgroup.cancel"
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordTaskgroupCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessWorkrecordTaskgroupCancelResponse struct {
	taobao.TaobaoResponse
}
