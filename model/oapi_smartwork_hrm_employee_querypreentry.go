package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeQuerypreentryRequest() *OapiSmartworkHrmEmployeeQuerypreentryRequest {
	return &OapiSmartworkHrmEmployeeQuerypreentryRequest{}
}

type OapiSmartworkHrmEmployeeQuerypreentryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeQuerypreentryResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.querypreentry"
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeQuerypreentryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeQuerypreentryResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
