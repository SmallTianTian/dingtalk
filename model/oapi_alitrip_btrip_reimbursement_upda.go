package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripReimbursementUpdateRequest() *OapiAlitripBtripReimbursementUpdateRequest {
	return &OapiAlitripBtripReimbursementUpdateRequest{}
}

type OapiAlitripBtripReimbursementUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripReimbursementUpdateResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripReimbursementUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.reimbursement.update"
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripReimbursementUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Auditlist struct {
	Note        string    `json:"note,omitempty"`
	OperateTime time.Time `json:"operate_time,omitempty"`
	Status      int64     `json:"status,omitempty"`
	Userid      string    `json:"userid,omitempty"`
}
type OpenApiUpdateReimbursementRq struct {
	ApplyFlowNo      int64       `json:"apply_flow_no,omitempty"`
	AuditList        []Auditlist `json:"audit_list,omitempty"`
	Corpid           string      `json:"corpid,omitempty"`
	FlowNo           int64       `json:"flow_no,omitempty"`
	OperateTime      time.Time   `json:"operate_time,omitempty"`
	OrderIds         string      `json:"order_ids,omitempty"`
	Status           int64       `json:"status,omitempty"`
	ThirdpartyFlowId string      `json:"thirdparty_flow_id,omitempty"`
}
type OapiAlitripBtripReimbursementUpdateResponse struct {
	taobao.TaobaoResponse
	Module  bool `json:"module,omitempty"`
	Success bool `json:"success,omitempty"`
}
