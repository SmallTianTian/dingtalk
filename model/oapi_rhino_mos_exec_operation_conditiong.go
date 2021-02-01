package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecOperationConditionGetRequest() *OapiRhinoMosExecOperationConditionGetRequest {
	return &OapiRhinoMosExecOperationConditionGetRequest{}
}

type OapiRhinoMosExecOperationConditionGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiRhinoMosExecOperationConditionGetResponse
	GetClothesConditionReq string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiRhinoMosExecOperationConditionGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) SetGetClothesConditionReq(getClothesConditionReq2 string) {
	this.GetClothesConditionReq = getClothesConditionReq2
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetGetClothesConditionReq() string {
	return this.GetClothesConditionReq
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.operation.condition.get"
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["get_clothes_condition_req"] = this.GetClothesConditionReq
	return txtParams
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecOperationConditionGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecOperationConditionGetResponse struct {
	taobao.TaobaoResponse

	ExternalMsgInfo string     `json:"external_msg_info,omitempty"`
	Model           PageResult `json:"model,omitempty"`
	Success         bool       `json:"success,omitempty"`
}
