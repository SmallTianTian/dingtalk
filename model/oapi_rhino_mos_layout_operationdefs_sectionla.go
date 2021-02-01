package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefsSectionlastRequest() *OapiRhinoMosLayoutOperationdefsSectionlastRequest {
	return &OapiRhinoMosLayoutOperationdefsSectionlastRequest{}
}

type OapiRhinoMosLayoutOperationdefsSectionlastRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosLayoutOperationdefsSectionlastResponse
	FlowVersion     int64
	NeedAssignInfo  bool
	OrderId         int64
	SectionCode     string
	TenantId        string
	TmpSave         bool
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetNeedAssignInfo(needAssignInfo2 bool) {
	this.NeedAssignInfo = needAssignInfo2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetNeedAssignInfo() bool {
	return this.NeedAssignInfo
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetSectionCode(sectionCode2 string) {
	this.SectionCode = sectionCode2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetSectionCode() string {
	return this.SectionCode
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetTmpSave(tmpSave2 bool) {
	this.TmpSave = tmpSave2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTmpSave() bool {
	return this.TmpSave
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdefs.sectionlast"
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["flow_version"] = this.FlowVersion
	txtParams["need_assign_info"] = this.NeedAssignInfo
	txtParams["order_id"] = this.OrderId
	txtParams["section_code"] = this.SectionCode
	txtParams["tenant_id"] = this.TenantId
	txtParams["tmp_save"] = this.TmpSave
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.NeedAssignInfo, "needAssignInfo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefsSectionlastRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefsSectionlastResponse struct {
	taobao.TaobaoResponse
	Result []OperationDefDto `json:"result,omitempty"`
}
