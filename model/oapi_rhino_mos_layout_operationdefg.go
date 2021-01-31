package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefGetRequest() *OapiRhinoMosLayoutOperationdefGetRequest {
	return &OapiRhinoMosLayoutOperationdefGetRequest{}
}

type OapiRhinoMosLayoutOperationdefGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiRhinoMosLayoutOperationdefGetResponse
	FlowVersion         int64
	NeedAssignInfo      bool
	OperationExternalId string
	OperationUid        int64
	OrderId             int64
	TenantId            string
	TmpSave             bool
	TopHttpMethod       string
	TopResponseType     string
	Userid              string
}

func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetNeedAssignInfo(needAssignInfo2 bool) {
	this.NeedAssignInfo = needAssignInfo2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetNeedAssignInfo() bool {
	return this.NeedAssignInfo
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetOperationExternalId(operationExternalId2 string) {
	this.OperationExternalId = operationExternalId2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetOperationExternalId() string {
	return this.OperationExternalId
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetOperationUid(operationUid2 int64) {
	this.OperationUid = operationUid2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetOperationUid() int64 {
	return this.OperationUid
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetTmpSave(tmpSave2 bool) {
	this.TmpSave = tmpSave2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTmpSave() bool {
	return this.TmpSave
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdef.get"
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["flow_version"] = this.FlowVersion
	txtParams["need_assign_info"] = this.NeedAssignInfo
	txtParams["operation_external_id"] = this.OperationExternalId
	txtParams["operation_uid"] = this.OperationUid
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams["tmp_save"] = this.TmpSave
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.NeedAssignInfo, "needAssignInfo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefGetResponse struct {
	taobao.TaobaoResponse
	Result OperationDefDto `json:"result,omitempty"`
}
type OperationExecutorAssignDto struct {
	ExecutorId string `json:"executor_id,omitempty"`
	OrderNum   int64  `json:"order_num,omitempty"`
}
type DeviceAssignDto struct {
	DeviceId           int64  `json:"device_id,omitempty"`
	DeviceModelId      string `json:"device_model_id,omitempty"`
	DeviceModelName    string `json:"device_model_name,omitempty"`
	DeviceModelVersion string `json:"device_model_version,omitempty"`
}
type WorkerAssignDto struct {
	Name   string `json:"name,omitempty"`
	WorkNo string `json:"work_no,omitempty"`
}
type WorkstationAssignDto struct {
	WorkstationCode string `json:"workstation_code,omitempty"`
}
type WorkUnitDto struct {
	DeviceAssigns      []DeviceAssignDto      `json:"device_assigns,omitempty"`
	WorkerAssigns      []WorkerAssignDto      `json:"worker_assigns,omitempty"`
	WorkstationAssigns []WorkstationAssignDto `json:"workstation_assigns,omitempty"`
}
type OperationDefDto struct {
	ApplicableSizeCode       string                       `json:"applicable_size_code,omitempty"`
	AutoSchedule             bool                         `json:"auto_schedule,omitempty"`
	BizCode                  string                       `json:"biz_code,omitempty"`
	BizSource                string                       `json:"biz_source,omitempty"`
	EnterCondition           string                       `json:"enter_condition,omitempty"`
	ExecSystem               string                       `json:"exec_system,omitempty"`
	FlowId                   int64                        `json:"flow_id,omitempty"`
	FlowVersion              int64                        `json:"flow_version,omitempty"`
	Name                     string                       `json:"name,omitempty"`
	NextOperationExternalIds []string                     `json:"next_operation_external_ids,omitempty"`
	NextOperationUids        []int64                      `json:"next_operation_uids,omitempty"`
	OperationExecutorAssigns []OperationExecutorAssignDto `json:"operation_executor_assigns,omitempty"`
	OperationExternalId      string                       `json:"operation_external_id,omitempty"`
	OperationType            string                       `json:"operation_type,omitempty"`
	OperationUid             int64                        `json:"operation_uid,omitempty"`
	ProcessTypeCode          string                       `json:"process_type_code,omitempty"`
	SectionCode              string                       `json:"section_code,omitempty"`
	SectionName              string                       `json:"section_name,omitempty"`
	Skip                     bool                         `json:"skip,omitempty"`
	StdCost                  string                       `json:"std_cost,omitempty"`
	WorkUnits                []WorkUnitDto                `json:"work_units,omitempty"`
}
