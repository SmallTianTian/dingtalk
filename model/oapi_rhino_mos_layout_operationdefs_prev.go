package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefsPrevRequest() *OapiRhinoMosLayoutOperationdefsPrevRequest {
	return &OapiRhinoMosLayoutOperationdefsPrevRequest{}
}

type OapiRhinoMosLayoutOperationdefsPrevRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiRhinoMosLayoutOperationdefsPrevResponse
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

func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetNeedAssignInfo(needAssignInfo2 bool) {
	this.NeedAssignInfo = needAssignInfo2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetNeedAssignInfo() bool {
	return this.NeedAssignInfo
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetOperationExternalId(operationExternalId2 string) {
	this.OperationExternalId = operationExternalId2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetOperationExternalId() string {
	return this.OperationExternalId
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetOperationUid(operationUid2 int64) {
	this.OperationUid = operationUid2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetOperationUid() int64 {
	return this.OperationUid
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetTmpSave(tmpSave2 bool) {
	this.TmpSave = tmpSave2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTmpSave() bool {
	return this.TmpSave
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdefs.prev"
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.NeedAssignInfo, "needAssignInfo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefsPrevRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefsPrevResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  []OperationDefDto `json:"result,omitempty"`
}
