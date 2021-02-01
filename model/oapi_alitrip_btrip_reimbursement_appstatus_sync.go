package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripReimbursementAppstatusSyncRequest() *OapiAlitripBtripReimbursementAppstatusSyncRequest {
	return &OapiAlitripBtripReimbursementAppstatusSyncRequest{}
}

type OapiAlitripBtripReimbursementAppstatusSyncRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripReimbursementAppstatusSyncResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.reimbursement.appstatus.sync"
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripReimbursementAppstatusSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiUpdateAppStatusRq struct {
	Corpid    string `json:"corpid,omitempty"`
	Installed bool   `json:"installed,omitempty"`
	IsvCode   string `json:"isv_code,omitempty"`
	Version   string `json:"version,omitempty"`
}
type OapiAlitripBtripReimbursementAppstatusSyncResponse struct {
	taobao.TaobaoResponse
	Module int64 `json:"module,omitempty"`
}
