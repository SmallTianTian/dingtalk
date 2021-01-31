package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmMasterdataSaveRequest() *OapiSmartworkHrmMasterdataSaveRequest {
	return &OapiSmartworkHrmMasterdataSaveRequest{}
}

type OapiSmartworkHrmMasterdataSaveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmMasterdataSaveResponse
	BizDataFields   string
	OuterId         string
	ScopeCode       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiSmartworkHrmMasterdataSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetBizDataFields(bizDataFields2 string) {
	this.BizDataFields = bizDataFields2
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetBizDataFields() string {
	return this.BizDataFields
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetOuterId(outerId2 string) {
	this.OuterId = outerId2
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetOuterId() string {
	return this.OuterId
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetScopeCode(scopeCode2 string) {
	this.ScopeCode = scopeCode2
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetScopeCode() string {
	return this.ScopeCode
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.masterdata.save"
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_data_fields"] = this.BizDataFields
	txtParams["outer_id"] = this.OuterId
	txtParams["scope_code"] = this.ScopeCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.BizDataFields, 999, "bizDataFields"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmMasterdataSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BizDataFieldVo struct {
	Name     string `json:"name,omitempty"`
	ValueStr string `json:"value_str,omitempty"`
}
type OapiSmartworkHrmMasterdataSaveResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
