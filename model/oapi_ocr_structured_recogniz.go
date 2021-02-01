package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiOcrStructuredRecognizeRequest() *OapiOcrStructuredRecognizeRequest {
	return &OapiOcrStructuredRecognizeRequest{}
}

type OapiOcrStructuredRecognizeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOcrStructuredRecognizeResponse
	ImageUrl        string
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *OapiOcrStructuredRecognizeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOcrStructuredRecognizeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOcrStructuredRecognizeRequest) SetImageUrl(imageUrl2 string) {
	this.ImageUrl = imageUrl2
}
func (this *OapiOcrStructuredRecognizeRequest) GetImageUrl() string {
	return this.ImageUrl
}
func (this *OapiOcrStructuredRecognizeRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiOcrStructuredRecognizeRequest) GetType() string {
	return this.Type
}
func (this *OapiOcrStructuredRecognizeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ocr.structured.recognize"
}
func (this *OapiOcrStructuredRecognizeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOcrStructuredRecognizeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOcrStructuredRecognizeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOcrStructuredRecognizeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOcrStructuredRecognizeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["image_url"] = this.ImageUrl
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiOcrStructuredRecognizeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ImageUrl, "imageUrl"); err != nil {
		return err
	}
	return nil
}
func (this *OapiOcrStructuredRecognizeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOcrStructuredRecognizeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOcrStructuredRecognizeResponse struct {
	taobao.TaobaoResponse
	Result OcrStructuredResult `json:"result,omitempty"`
}
type OcrStructuredResult struct {
	Angle          int64  `json:"angle,omitempty"`
	Data           string `json:"data,omitempty"`
	Height         int64  `json:"height,omitempty"`
	OriginalHeight int64  `json:"original_height,omitempty"`
	OriginalWidth  int64  `json:"original_width,omitempty"`
	Width          int64  `json:"width,omitempty"`
}
