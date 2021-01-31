package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAiMtTranslateRequest() *OapiAiMtTranslateRequest {
	return &OapiAiMtTranslateRequest{}
}

type OapiAiMtTranslateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAiMtTranslateResponse
	Query           string
	SourceLanguage  string
	TargetLanguage  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAiMtTranslateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAiMtTranslateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAiMtTranslateRequest) SetQuery(query2 string) {
	this.Query = query2
}
func (this *OapiAiMtTranslateRequest) GetQuery() string {
	return this.Query
}
func (this *OapiAiMtTranslateRequest) SetSourceLanguage(sourceLanguage2 string) {
	this.SourceLanguage = sourceLanguage2
}
func (this *OapiAiMtTranslateRequest) GetSourceLanguage() string {
	return this.SourceLanguage
}
func (this *OapiAiMtTranslateRequest) SetTargetLanguage(targetLanguage2 string) {
	this.TargetLanguage = targetLanguage2
}
func (this *OapiAiMtTranslateRequest) GetTargetLanguage() string {
	return this.TargetLanguage
}
func (this *OapiAiMtTranslateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ai.mt.translate"
}
func (this *OapiAiMtTranslateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAiMtTranslateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAiMtTranslateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAiMtTranslateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAiMtTranslateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["query"] = this.Query
	txtParams["source_language"] = this.SourceLanguage
	txtParams["target_language"] = this.TargetLanguage
	return txtParams
}
func (this *OapiAiMtTranslateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Query, "query"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAiMtTranslateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAiMtTranslateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAiMtTranslateResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
