package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceWorkerCheckOutRequest() *OapiRhinoMosSpaceWorkerCheckOutRequest {
	return &OapiRhinoMosSpaceWorkerCheckOutRequest{}
}

type OapiRhinoMosSpaceWorkerCheckOutRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceWorkerCheckOutResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.worker.check.out"
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceWorkerCheckOutRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosSpaceWorkerCheckOutResponse struct {
	taobao.TaobaoResponse
	Model bool `json:"model,omitempty"`
}
