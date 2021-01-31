package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiRhinoMosExecPerformCreateRequest() *OapiRhinoMosExecPerformCreateRequest {
	return &OapiRhinoMosExecPerformCreateRequest{}
}

type OapiRhinoMosExecPerformCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecPerformCreateResponse
	Operations      string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosExecPerformCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetOperations(operations2 string) {
	this.Operations = operations2
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetOperations() string {
	return this.Operations
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.create"
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operations"] = this.Operations
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecPerformCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Operations, 500, "operations"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OperationReq struct {
	Context          string    `json:"context,omitempty"`
	DeviceIds        []int64   `json:"device_ids,omitempty"`
	EntityId         int64     `json:"entity_id,omitempty"`
	EntityType       string    `json:"entity_type,omitempty"`
	FlowVersion      int64     `json:"flow_version,omitempty"`
	OperationType    string    `json:"operation_type,omitempty"`
	OperationUid     int64     `json:"operation_uid,omitempty"`
	OrderId          int64     `json:"order_id,omitempty"`
	PerformStatus    string    `json:"perform_status,omitempty"`
	PerformType      string    `json:"perform_type,omitempty"`
	Priority         int64     `json:"priority,omitempty"`
	ProcessCostTime  string    `json:"process_cost_time,omitempty"`
	ProcessEndTime   time.Time `json:"process_end_time,omitempty"`
	ProcessStartTime time.Time `json:"process_start_time,omitempty"`
	ProcessTypeCode  string    `json:"process_type_code,omitempty"`
	SectionCode      string    `json:"section_code,omitempty"`
	SourceId         string    `json:"source_id,omitempty"`
	SourceType       string    `json:"source_type,omitempty"`
	WorkNos          []string  `json:"work_nos,omitempty"`
	WorkstationCode  string    `json:"workstation_code,omitempty"`
}
type OapiRhinoMosExecPerformCreateResponse struct {
	taobao.TaobaoResponse
	Model   []OperationPerformDto `json:"model,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
