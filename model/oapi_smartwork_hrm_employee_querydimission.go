package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeQuerydimissionRequest() *OapiSmartworkHrmEmployeeQuerydimissionRequest {
	return &OapiSmartworkHrmEmployeeQuerydimissionRequest{}
}

type OapiSmartworkHrmEmployeeQuerydimissionRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeQuerydimissionResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.querydimission"
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeQuerydimissionRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeQuerydimissionResponse struct {
	taobao.TaobaoResponse
	Result  Paginator `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
type Paginator struct {
	DataList   []string `json:"data_list,omitempty"`
	NextCursor int64    `json:"next_cursor,omitempty"`
}
