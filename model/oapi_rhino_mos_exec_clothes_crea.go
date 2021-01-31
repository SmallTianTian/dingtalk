package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiRhinoMosExecClothesCreateRequest() *OapiRhinoMosExecClothesCreateRequest {
	return &OapiRhinoMosExecClothesCreateRequest{}
}

type OapiRhinoMosExecClothesCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiRhinoMosExecClothesCreateResponse
	AdditionalOperations string
	AutoStart            bool
	BizType              string
	Clothes              string
	EntityType           string
	OrderId              int64
	Source               string
	TenantId             string
	TopHttpMethod        string
	TopResponseType      string
	Userid               string
}

func (this *OapiRhinoMosExecClothesCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetAdditionalOperations(additionalOperations2 string) {
	this.AdditionalOperations = additionalOperations2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetAdditionalOperations() string {
	return this.AdditionalOperations
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetAutoStart(autoStart2 bool) {
	this.AutoStart = autoStart2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetAutoStart() bool {
	return this.AutoStart
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetBizType(bizType2 string) {
	this.BizType = bizType2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetBizType() string {
	return this.BizType
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetClothes(clothes2 string) {
	this.Clothes = clothes2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetClothes() string {
	return this.Clothes
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetEntityType(entityType2 string) {
	this.EntityType = entityType2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetEntityType() string {
	return this.EntityType
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetSource(source2 string) {
	this.Source = source2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetSource() string {
	return this.Source
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.create"
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["additional_operations"] = this.AdditionalOperations
	txtParams["auto_start"] = this.AutoStart
	txtParams["biz_type"] = this.BizType
	txtParams["clothes"] = this.Clothes
	txtParams["entity_type"] = this.EntityType
	txtParams["order_id"] = this.OrderId
	txtParams["source"] = this.Source
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecClothesCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.AdditionalOperations, 500, "additionalOperations"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AddtionalOperation struct {
	Context          string    `json:"context,omitempty"`
	DeviceIds        []int64   `json:"device_ids,omitempty"`
	FlowVersion      int64     `json:"flow_version,omitempty"`
	OperationType    string    `json:"operation_type,omitempty"`
	OperationUid     int64     `json:"operation_uid,omitempty"`
	PerformStatus    string    `json:"perform_status,omitempty"`
	Priority         int64     `json:"priority,omitempty"`
	ProcessEndTime   time.Time `json:"process_end_time,omitempty"`
	ProcessStartTime time.Time `json:"process_start_time,omitempty"`
	ProcessTypeCode  string    `json:"process_type_code,omitempty"`
	SectionCode      string    `json:"section_code,omitempty"`
	WorkNos          []string  `json:"work_nos,omitempty"`
	WorkstationCode  string    `json:"workstation_code,omitempty"`
}
type ClothesInfoReq struct {
	ColorCode  string    `json:"color_code,omitempty"`
	ColorName  string    `json:"color_name,omitempty"`
	Count      int64     `json:"count,omitempty"`
	FinishTime time.Time `json:"finish_time,omitempty"`
	SizeCode   string    `json:"size_code,omitempty"`
	SizeName   string    `json:"size_name,omitempty"`
	StartTime  time.Time `json:"start_time,omitempty"`
}
type OapiRhinoMosExecClothesCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64   `json:"errcode,omitempty"`
	Errmsg  string  `json:"errmsg,omitempty"`
	Model   []int64 `json:"model,omitempty"`
	Success bool    `json:"success,omitempty"`
}
