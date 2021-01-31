package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmMasterCorpconfigUpdateRequest() *OapiSmartworkHrmMasterCorpconfigUpdateRequest {
	return &OapiSmartworkHrmMasterCorpconfigUpdateRequest{}
}

type OapiSmartworkHrmMasterCorpconfigUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmMasterCorpconfigUpdateResponse
	ScopeCode       string
	Status          int64
	TenantId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) SetScopeCode(scopeCode2 string) {
	this.ScopeCode = scopeCode2
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetScopeCode() string {
	return this.ScopeCode
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) SetTenantId(tenantId2 int64) {
	this.TenantId = tenantId2
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetTenantId() int64 {
	return this.TenantId
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.master.corpconfig.update"
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["scope_code"] = this.ScopeCode
	txtParams["status"] = this.Status
	txtParams["tenant_id"] = this.TenantId
	return txtParams
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ScopeCode, "scopeCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmMasterCorpconfigUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmMasterCorpconfigUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
