package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPlatformTranslateRequest() *OapiPlatformTranslateRequest {
	return &OapiPlatformTranslateRequest{}
}

type OapiPlatformTranslateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPlatformTranslateResponse
	Query           string
	SourceLanguage  string
	TargetLanguage  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPlatformTranslateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPlatformTranslateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPlatformTranslateRequest) SetQuery(query2 string) {
	this.Query = query2
}
func (this *OapiPlatformTranslateRequest) GetQuery() string {
	return this.Query
}
func (this *OapiPlatformTranslateRequest) SetSourceLanguage(sourceLanguage2 string) {
	this.SourceLanguage = sourceLanguage2
}
func (this *OapiPlatformTranslateRequest) GetSourceLanguage() string {
	return this.SourceLanguage
}
func (this *OapiPlatformTranslateRequest) SetTargetLanguage(targetLanguage2 string) {
	this.TargetLanguage = targetLanguage2
}
func (this *OapiPlatformTranslateRequest) GetTargetLanguage() string {
	return this.TargetLanguage
}
func (this *OapiPlatformTranslateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.platform.translate"
}
func (this *OapiPlatformTranslateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPlatformTranslateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPlatformTranslateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPlatformTranslateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPlatformTranslateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["query"] = this.Query
	txtParams["source_language"] = this.SourceLanguage
	txtParams["target_language"] = this.TargetLanguage
	return txtParams
}
func (this *OapiPlatformTranslateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Query, "query"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPlatformTranslateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPlatformTranslateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPlatformTranslateResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
