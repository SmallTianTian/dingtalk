package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmMasterCheckRequest() *OapiSmartworkHrmMasterCheckRequest {
	return &OapiSmartworkHrmMasterCheckRequest{}
}

type OapiSmartworkHrmMasterCheckRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmMasterCheckResponse
	BizUk           string
	Tenantid        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmMasterCheckRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmMasterCheckRequest) SetBizUk(bizUk2 string) {
	this.BizUk = bizUk2
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetBizUk() string {
	return this.BizUk
}
func (this *OapiSmartworkHrmMasterCheckRequest) SetTenantid(tenantid2 int64) {
	this.Tenantid = tenantid2
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetTenantid() int64 {
	return this.Tenantid
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.master.check"
}
func (this *OapiSmartworkHrmMasterCheckRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmMasterCheckRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmMasterCheckRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_uk"] = this.BizUk
	txtParams["tenantid"] = this.Tenantid
	return txtParams
}
func (this *OapiSmartworkHrmMasterCheckRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizUk, "bizUk"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmMasterCheckRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmMasterCheckResponse struct {
	taobao.TaobaoResponse
	Result IntegrateBatchResultVo `json:"result,omitempty"`
}
type Failresult struct {
	BizUk     string `json:"biz_uk,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
	Success   bool   `json:"success,omitempty"`
}
type IntegrateBatchResultVo struct {
	AllSuccess bool         `json:"all_success,omitempty"`
	FailResult []Failresult `json:"fail_result,omitempty"`
}
