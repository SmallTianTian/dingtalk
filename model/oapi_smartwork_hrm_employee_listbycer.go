package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeListbycertRequest() *OapiSmartworkHrmEmployeeListbycertRequest {
	return &OapiSmartworkHrmEmployeeListbycertRequest{}
}

type OapiSmartworkHrmEmployeeListbycertRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeListbycertResponse
	Params          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) SetParams(params2 string) {
	this.Params = params2
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetParams() string {
	return this.Params
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.listbycert"
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["params"] = this.Params
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Params, 100, "params"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeListbycertRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type NameAndCertNumberQueryParam struct {
	CertNo   string `json:"cert_no,omitempty"`
	RealName string `json:"real_name,omitempty"`
}
type OapiSmartworkHrmEmployeeListbycertResponse struct {
	taobao.TaobaoResponse
	Result  []SimpleEmpPersonalInfoVo `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type SimpleEmpPersonalInfoVo struct {
	CertNo string `json:"cert_no,omitempty"`
	Userid string `json:"userid,omitempty"`
}
