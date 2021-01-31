package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceStatusUpdateRequest() *OapiCustomerserviceStatusUpdateRequest {
	return &OapiCustomerserviceStatusUpdateRequest{}
}

type OapiCustomerserviceStatusUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceStatusUpdateResponse
	StatusChange    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceStatusUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceStatusUpdateRequest) SetStatusChange(statusChange2 string) {
	this.StatusChange = statusChange2
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetStatusChange() string {
	return this.StatusChange
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.status.update"
}
func (this *OapiCustomerserviceStatusUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceStatusUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceStatusUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["status_change"] = this.StatusChange
	return txtParams
}
func (this *OapiCustomerserviceStatusUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceStatusUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ServiceStatusChangeDto struct {
	BuId         string `json:"bu_id,omitempty"`
	OriginStatus int64  `json:"origin_status,omitempty"`
	ServiceId    string `json:"service_id,omitempty"`
	Source       string `json:"source,omitempty"`
	TargetStatus int64  `json:"target_status,omitempty"`
}
type OapiCustomerserviceStatusUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
