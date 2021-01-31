package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefsSectionfirstRequest() *OapiRhinoMosLayoutOperationdefsSectionfirstRequest {
	return &OapiRhinoMosLayoutOperationdefsSectionfirstRequest{}
}

type OapiRhinoMosLayoutOperationdefsSectionfirstRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosLayoutOperationdefsSectionfirstResponse
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

func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetNeedAssignInfo(needAssignInfo2 bool) {
	this.NeedAssignInfo = needAssignInfo2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetNeedAssignInfo() bool {
	return this.NeedAssignInfo
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetSectionCode(sectionCode2 string) {
	this.SectionCode = sectionCode2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetSectionCode() string {
	return this.SectionCode
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetTmpSave(tmpSave2 bool) {
	this.TmpSave = tmpSave2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTmpSave() bool {
	return this.TmpSave
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdefs.sectionfirst"
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.NeedAssignInfo, "needAssignInfo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefsSectionfirstRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefsSectionfirstResponse struct {
	taobao.TaobaoResponse
	Result []OperationDefDto `json:"result,omitempty"`
}
