package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceWorkerCheckInRequest() *OapiRhinoMosSpaceWorkerCheckInRequest {
	return &OapiRhinoMosSpaceWorkerCheckInRequest{}
}

type OapiRhinoMosSpaceWorkerCheckInRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceWorkerCheckInResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.worker.check.in"
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceWorkerCheckInRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationWorkerCheckReq struct {
	ProdWorkstationCode string `json:"prod_workstation_code,omitempty"`
	TenantId            string `json:"tenant_id,omitempty"`
	Userid              string `json:"userid,omitempty"`
	WorkNo              string `json:"work_no,omitempty"`
}
type OapiRhinoMosSpaceWorkerCheckInResponse struct {
	taobao.TaobaoResponse
	Model bool `json:"model,omitempty"`
}
