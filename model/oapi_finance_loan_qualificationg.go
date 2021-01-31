package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanQualificationGetRequest() *OapiFinanceLoanQualificationGetRequest {
	return &OapiFinanceLoanQualificationGetRequest{}
}

type OapiFinanceLoanQualificationGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanQualificationGetResponse
	Code            string
	OpenChannelType string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiFinanceLoanQualificationGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanQualificationGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanQualificationGetRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiFinanceLoanQualificationGetRequest) GetCode() string {
	return this.Code
}
func (this *OapiFinanceLoanQualificationGetRequest) SetOpenChannelType(openChannelType2 string) {
	this.OpenChannelType = openChannelType2
}
func (this *OapiFinanceLoanQualificationGetRequest) GetOpenChannelType() string {
	return this.OpenChannelType
}
func (this *OapiFinanceLoanQualificationGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.qualification.get"
}
func (this *OapiFinanceLoanQualificationGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanQualificationGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanQualificationGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanQualificationGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanQualificationGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.ERROR_CODE] = this.Code
	txtParams["open_channel_type"] = this.OpenChannelType
	return txtParams
}
func (this *OapiFinanceLoanQualificationGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Code, taobao.ERROR_CODE); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanQualificationGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanQualificationGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanQualificationGetResponse struct {
	taobao.TaobaoResponse
	Result  CustomerInfo `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
type CustomerInfo struct {
	Enable string `json:"enable,omitempty"`
	Url    string `json:"url,omitempty"`
}
