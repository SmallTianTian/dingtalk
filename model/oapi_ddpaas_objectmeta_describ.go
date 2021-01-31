package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDdpaasObjectmetaDescribeRequest() *OapiDdpaasObjectmetaDescribeRequest {
	return &OapiDdpaasObjectmetaDescribeRequest{}
}

type OapiDdpaasObjectmetaDescribeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDdpaasObjectmetaDescribeResponse
	AppUuid         string
	FormCode        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDdpaasObjectmetaDescribeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDdpaasObjectmetaDescribeRequest) SetAppUuid(appUuid2 string) {
	this.AppUuid = appUuid2
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetAppUuid() string {
	return this.AppUuid
}
func (this *OapiDdpaasObjectmetaDescribeRequest) SetFormCode(formCode2 string) {
	this.FormCode = formCode2
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetFormCode() string {
	return this.FormCode
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ddpaas.objectmeta.describe"
}
func (this *OapiDdpaasObjectmetaDescribeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDdpaasObjectmetaDescribeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDdpaasObjectmetaDescribeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_uuid"] = this.AppUuid
	txtParams["form_code"] = this.FormCode
	return txtParams
}
func (this *OapiDdpaasObjectmetaDescribeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppUuid, "appUuid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDdpaasObjectmetaDescribeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDdpaasObjectmetaDescribeResponse struct {
	taobao.TaobaoResponse
	Result DObject `json:"result,omitempty"`
}
type Selectoptions struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
type Referencefields struct {
	Format        string          `json:taobao.FORMA,omitemptyT`
	Label         string          `json:"label,omitempty"`
	Name          string          `json:"name,omitempty"`
	Nillable      bool            `json:"nillable,omitempty"`
	SelectOptions []Selectoptions `json:"select_options,omitempty"`
	Type          string          `json:"type,omitempty"`
	Unit          string          `json:"unit,omitempty"`
}
type Rollupsummaryfields struct {
	Aggregator string `json:"aggregator,omitempty"`
	Name       string `json:"name,omitempty"`
}
