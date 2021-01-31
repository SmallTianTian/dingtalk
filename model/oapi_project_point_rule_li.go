package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProjectPointRuleListRequest() *OapiProjectPointRuleListRequest {
	return &OapiProjectPointRuleListRequest{}
}

type OapiProjectPointRuleListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProjectPointRuleListResponse
	TenantId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProjectPointRuleListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProjectPointRuleListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProjectPointRuleListRequest) SetTenantId(tenantId2 int64) {
	this.TenantId = tenantId2
}
func (this *OapiProjectPointRuleListRequest) GetTenantId() int64 {
	return this.TenantId
}
func (this *OapiProjectPointRuleListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.project.point.rule.list"
}
func (this *OapiProjectPointRuleListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProjectPointRuleListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProjectPointRuleListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProjectPointRuleListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProjectPointRuleListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["tenant_id"] = this.TenantId
	return txtParams
}
func (this *OapiProjectPointRuleListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TenantId, "tenantId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProjectPointRuleListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProjectPointRuleListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProjectPointRuleListResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
