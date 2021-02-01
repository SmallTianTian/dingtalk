package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsPluginDataDeleteRequest() *OapiAtsPluginDataDeleteRequest {
	return &OapiAtsPluginDataDeleteRequest{}
}

type OapiAtsPluginDataDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsPluginDataDeleteResponse
	BizCode         string
	Header          string
	OutId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsPluginDataDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsPluginDataDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsPluginDataDeleteRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsPluginDataDeleteRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsPluginDataDeleteRequest) SetHeader(header2 string) {
	this.Header = header2
}
func (this *OapiAtsPluginDataDeleteRequest) GetHeader() string {
	return this.Header
}
func (this *OapiAtsPluginDataDeleteRequest) SetOutId(outId2 string) {
	this.OutId = outId2
}
func (this *OapiAtsPluginDataDeleteRequest) GetOutId() string {
	return this.OutId
}
func (this *OapiAtsPluginDataDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.plugin.data.delete"
}
func (this *OapiAtsPluginDataDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsPluginDataDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsPluginDataDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsPluginDataDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsPluginDataDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["header"] = this.Header
	txtParams["out_id"] = this.OutId
	return txtParams
}
func (this *OapiAtsPluginDataDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsPluginDataDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsPluginDataDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PushHeaderVO struct {
	PluginCode string `json:"plugin_code,omitempty"`
	SchemaId   string `json:"schema_id,omitempty"`
}
type OapiAtsPluginDataDeleteResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
