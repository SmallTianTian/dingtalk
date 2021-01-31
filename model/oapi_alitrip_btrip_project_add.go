package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripProjectAddRequest() *OapiAlitripBtripProjectAddRequest {
	return &OapiAlitripBtripProjectAddRequest{}
}

type OapiAlitripBtripProjectAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripProjectAddResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripProjectAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripProjectAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripProjectAddRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripProjectAddRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripProjectAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.project.add"
}
func (this *OapiAlitripBtripProjectAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripProjectAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripProjectAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripProjectAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripProjectAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripProjectAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripProjectAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripProjectAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenProjectRs struct {
	Code                  string `json:taobao.ERROR_COD,omitemptyE`
	Corpid                string `json:"corpid,omitempty"`
	ProjectName           string `json:"project_name,omitempty"`
	ThirdPartCostCenterId string `json:"third_part_cost_center_id,omitempty"`
	ThirdPartId           string `json:"third_part_id,omitempty"`
	ThirdPartInvoiceId    string `json:"third_part_invoice_id,omitempty"`
}
type OapiAlitripBtripProjectAddResponse struct {
	taobao.TaobaoResponse
	Module  string `json:"module,omitempty"`
	Success bool   `json:"success,omitempty"`
}
