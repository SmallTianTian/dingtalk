package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmMasterSaveRequest() *OapiSmartworkHrmMasterSaveRequest {
	return &OapiSmartworkHrmMasterSaveRequest{}
}

type OapiSmartworkHrmMasterSaveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmMasterSaveResponse
	BizData         string
	TenantId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmMasterSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmMasterSaveRequest) SetBizData(bizData2 string) {
	this.BizData = bizData2
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetBizData() string {
	return this.BizData
}
func (this *OapiSmartworkHrmMasterSaveRequest) SetTenantId(tenantId2 int64) {
	this.TenantId = tenantId2
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetTenantId() int64 {
	return this.TenantId
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.master.save"
}
func (this *OapiSmartworkHrmMasterSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmMasterSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmMasterSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_data"] = this.BizData
	txtParams["tenant_id"] = this.TenantId
	return txtParams
}
func (this *OapiSmartworkHrmMasterSaveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.BizData, 50, "bizData"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmMasterSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Bizdatavo struct {
	BizTime int64       `json:"biz_time,omitempty"`
	BizUk   string      `json:"biz_uk,omitempty"`
	Fields  []Fieldlist `json:"fields,omitempty"`
	Scope   ScopeVo     `json:"scope,omitempty"`
	Userid  string      `json:"userid,omitempty"`
}
type OapiSmartworkHrmMasterSaveResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  IntegrateBatchResultVO `json:"result,omitempty"`
}
type IntegrateBatchResultVO struct {
	AllSuccess bool         `json:"all_success,omitempty"`
	FailResult []Failresult `json:"fail_result,omitempty"`
}
