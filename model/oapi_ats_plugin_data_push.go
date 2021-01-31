package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsPluginDataPushRequest() *OapiAtsPluginDataPushRequest {
	return &OapiAtsPluginDataPushRequest{}
}

type OapiAtsPluginDataPushRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsPluginDataPushResponse
	BizCode         string
	Content         string
	Header          string
	OutId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsPluginDataPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsPluginDataPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsPluginDataPushRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsPluginDataPushRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsPluginDataPushRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiAtsPluginDataPushRequest) SetContentString(content2 string) {
	this.Content = content2
}
func (this *OapiAtsPluginDataPushRequest) GetContent() string {
	return this.Content
}
func (this *OapiAtsPluginDataPushRequest) SetHeader(header2 string) {
	this.Header = header2
}
func (this *OapiAtsPluginDataPushRequest) GetHeader() string {
	return this.Header
}
func (this *OapiAtsPluginDataPushRequest) SetOutId(outId2 string) {
	this.OutId = outId2
}
func (this *OapiAtsPluginDataPushRequest) GetOutId() string {
	return this.OutId
}
func (this *OapiAtsPluginDataPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.plugin.data.push"
}
func (this *OapiAtsPluginDataPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsPluginDataPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsPluginDataPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsPluginDataPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsPluginDataPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["header"] = this.Header
	txtParams["out_id"] = this.OutId
	return txtParams
}
func (this *OapiAtsPluginDataPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsPluginDataPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsPluginDataPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PushHeaderVo struct {
	PluginCode string `json:"plugin_code,omitempty"`
	SchemaId   string `json:"schema_id,omitempty"`
}
type OapiAtsPluginDataPushResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
