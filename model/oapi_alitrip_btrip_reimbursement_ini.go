package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripReimbursementInitRequest() *OapiAlitripBtripReimbursementInitRequest {
	return &OapiAlitripBtripReimbursementInitRequest{}
}

type OapiAlitripBtripReimbursementInitRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripReimbursementInitResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripReimbursementInitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripReimbursementInitRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.reimbursement.init"
}
func (this *OapiAlitripBtripReimbursementInitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripReimbursementInitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripReimbursementInitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripReimbursementInitRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripReimbursementInitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ApproverNode struct {
	Note        string    `json:"note,omitempty"`
	OperateTime time.Time `json:"operate_time,omitempty"`
	Status      int64     `json:"status,omitempty"`
	Userid      string    `json:"userid,omitempty"`
}
type OpenApiNewReimbursementRq struct {
	ApplyFlowNo      int64          `json:"apply_flow_no,omitempty"`
	AuditList        []ApproverNode `json:"audit_list,omitempty"`
	Corpid           string         `json:"corpid,omitempty"`
	DepartId         string         `json:"depart_id,omitempty"`
	DepartName       string         `json:"depart_name,omitempty"`
	DetailUrl        string         `json:"detail_url,omitempty"`
	IsvCode          string         `json:"isv_code,omitempty"`
	Operator         OpenUserInfo   `json:"operator,omitempty"`
	OrderIds         string         `json:"order_ids,omitempty"`
	PayAmount        int64          `json:"pay_amount,omitempty"`
	Status           int64          `json:"status,omitempty"`
	ThirdpartyFlowId string         `json:"thirdparty_flow_id,omitempty"`
	Title            string         `json:"title,omitempty"`
}
type OapiAlitripBtripReimbursementInitResponse struct {
	taobao.TaobaoResponse
	Errcode int64                     `json:"errcode,omitempty"`
	Errmsg  string                    `json:"errmsg,omitempty"`
	Module  OpenApiNewReimbursementRs `json:"module,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OpenApiNewReimbursementRs struct {
	FlowNo           int64  `json:"flow_no,omitempty"`
	ThirdpartyFlowId string `json:"thirdparty_flow_id,omitempty"`
}
