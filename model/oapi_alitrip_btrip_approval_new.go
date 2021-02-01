package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripApprovalNewRequest() *OapiAlitripBtripApprovalNewRequest {
	return &OapiAlitripBtripApprovalNewRequest{}
}

type OapiAlitripBtripApprovalNewRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripApprovalNewResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripApprovalNewRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripApprovalNewRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripApprovalNewRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripApprovalNewRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripApprovalNewRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.approval.new"
}
func (this *OapiAlitripBtripApprovalNewRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripApprovalNewRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripApprovalNewRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripApprovalNewRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripApprovalNewRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripApprovalNewRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripApprovalNewRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripApprovalNewRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripApprovalNewResponse struct {
	taobao.TaobaoResponse
	Module  OpenApiNewApplyRs `json:"module,omitempty"`
	Success bool              `json:"success,omitempty"`
}
