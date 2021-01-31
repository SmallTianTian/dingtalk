package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceDeviceQuerylistRequest() *OapiSmartdeviceDeviceQuerylistRequest {
	return &OapiSmartdeviceDeviceQuerylistRequest{}
}

type OapiSmartdeviceDeviceQuerylistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDeviceQuerylistResponse
	PageQueryVo     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDeviceQuerylistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) SetPageQueryVo(pageQueryVo2 string) {
	this.PageQueryVo = pageQueryVo2
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetPageQueryVo() string {
	return this.PageQueryVo
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.device.querylist"
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["page_query_vo"] = this.PageQueryVo
	return txtParams
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDeviceQuerylistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PageQueryVo struct {
	Cursor int64  `json:"cursor,omitempty"`
	Pk     string `json:"pk,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiSmartdeviceDeviceQuerylistResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
