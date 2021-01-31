package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceWorkerCheckInListRequest() *OapiRhinoMosSpaceWorkerCheckInListRequest {
	return &OapiRhinoMosSpaceWorkerCheckInListRequest{}
}

type OapiRhinoMosSpaceWorkerCheckInListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceWorkerCheckInListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.worker.check.in.list"
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceWorkerCheckInListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationWorkerListReq struct {
	ProdWorkstationCode string `json:"prod_workstation_code,omitempty"`
	TenantId            string `json:"tenant_id,omitempty"`
	Userid              string `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceWorkerCheckInListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                           `json:"errcode,omitempty"`
	Errmsg  string                          `json:"errmsg,omitempty"`
	Model   []ProdWorkstationWorkerCheckDto `json:"model,omitempty"`
}
type ProdWorkstationWorkerCheckDto struct {
	CheckInTime         time.Time `json:"check_in_time,omitempty"`
	CheckOutTime        time.Time `json:"check_out_time,omitempty"`
	CheckStatus         string    `json:"check_status,omitempty"`
	ProdWorkstationCode string    `json:"prod_workstation_code,omitempty"`
	TenantId            string    `json:"tenant_id,omitempty"`
	WorkNo              string    `json:"work_no,omitempty"`
}
