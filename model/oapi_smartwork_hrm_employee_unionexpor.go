package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartworkHrmEmployeeUnionexportRequest() *OapiSmartworkHrmEmployeeUnionexportRequest {
	return &OapiSmartworkHrmEmployeeUnionexportRequest{}
}

type OapiSmartworkHrmEmployeeUnionexportRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeUnionexportResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.unionexport"
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeUnionexportRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UnionExportParam struct {
	BizUniqueId        string `json:"biz_unique_id,omitempty"`
	ExpireStrategy     int64  `json:"expire_strategy,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	MediaId            string `json:"media_id,omitempty"`
	PermissionStrategy int64  `json:"permission_strategy,omitempty"`
	Userid             string `json:"userid,omitempty"`
}
type OapiSmartworkHrmEmployeeUnionexportResponse struct {
	taobao.TaobaoResponse
	Errcode   int64       `json:"errcode,omitempty"`
	Errmsg    string      `json:"errmsg,omitempty"`
	IsSuccess bool        `json:"is_success,omitempty"`
	Result    ExportRsult `json:"result,omitempty"`
}
type ExportRsult struct {
	ExportId string `json:"export_id,omitempty"`
}
