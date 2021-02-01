package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecPerformBatchCreateRequest() *OapiRhinoMosExecPerformBatchCreateRequest {
	return &OapiRhinoMosExecPerformBatchCreateRequest{}
}

type OapiRhinoMosExecPerformBatchCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                    OapiRhinoMosExecPerformBatchCreateResponse
	BatchCreateOperationReq string
	TopHttpMethod           string
	TopResponseType         string
}

func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) SetBatchCreateOperationReq(batchCreateOperationReq2 string) {
	this.BatchCreateOperationReq = batchCreateOperationReq2
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetBatchCreateOperationReq() string {
	return this.BatchCreateOperationReq
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.batch.create"
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["batch_create_operation_req"] = this.BatchCreateOperationReq
	return txtParams
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformBatchCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PerformOperationReq struct {
	Context          string    `json:"context,omitempty"`
	DeviceIds        []int64   `json:"device_ids,omitempty"`
	FlowVersion      int64     `json:"flow_version,omitempty"`
	OperationType    string    `json:"operation_type,omitempty"`
	OperationUid     int64     `json:"operation_uid,omitempty"`
	PerformStatus    string    `json:"perform_status,omitempty"`
	Priority         int64     `json:"priority,omitempty"`
	ProcessCostTime  string    `json:"process_cost_time,omitempty"`
	ProcessEndTime   time.Time `json:"process_end_time,omitempty"`
	ProcessStartTime time.Time `json:"process_start_time,omitempty"`
	ProcessTypeCode  string    `json:"process_type_code,omitempty"`
	SectionCode      string    `json:"section_code,omitempty"`
	WorkNos          []string  `json:"work_nos,omitempty"`
	WorkstationCode  string    `json:"workstation_code,omitempty"`
}
type BatchCreateOperationWithEntityReq struct {
	EntityCondition      EntityCondition       `json:"entity_condition,omitempty"`
	OrderId              int64                 `json:"order_id,omitempty"`
	PerformOperationReqs []PerformOperationReq `json:"perform_operation_reqs,omitempty"`
	Source               Source                `json:"source,omitempty"`
	TenantId             string                `json:"tenant_id,omitempty"`
	Userid               string                `json:"userid,omitempty"`
}
type OapiRhinoMosExecPerformBatchCreateResponse struct {
	taobao.TaobaoResponse

	ExternalMsgInfo string                `json:"external_msg_info,omitempty"`
	Model           []OperationPerformDto `json:"model,omitempty"`
}
type OperationPerformDto struct {
	Active              string    `json:"active,omitempty"`
	BatchId             int64     `json:"batch_id,omitempty"`
	Context             string    `json:"context,omitempty"`
	CreateType          string    `json:"create_type,omitempty"`
	DeviceIds           []int64   `json:"device_ids,omitempty"`
	EntityId            int64     `json:"entity_id,omitempty"`
	EntityType          string    `json:"entity_type,omitempty"`
	FlowVersion         int64     `json:"flow_version,omitempty"`
	Id                  int64     `json:"id,omitempty"`
	OperationType       string    `json:"operation_type,omitempty"`
	OperationUid        int64     `json:"operation_uid,omitempty"`
	OrderId             int64     `json:"order_id,omitempty"`
	PerformStatus       string    `json:"perform_status,omitempty"`
	PerformType         string    `json:"perform_type,omitempty"`
	Priority            int64     `json:"priority,omitempty"`
	ProcessTypeCode     string    `json:"process_type_code,omitempty"`
	ProcessingEndTime   time.Time `json:"processing_end_time,omitempty"`
	ProcessingStartTime time.Time `json:"processing_start_time,omitempty"`
	SectionCode         string    `json:"section_code,omitempty"`
	SourceId            string    `json:"source_id,omitempty"`
	SourceType          string    `json:"source_type,omitempty"`
	TenantId            string    `json:"tenant_id,omitempty"`
	WorkNos             []string  `json:"work_nos,omitempty"`
	WorkstationCode     string    `json:"workstation_code,omitempty"`
}
