package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAsrVoiceTranslateRequest() *OapiAsrVoiceTranslateRequest {
	return &OapiAsrVoiceTranslateRequest{}
}

type OapiAsrVoiceTranslateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAsrVoiceTranslateResponse
	MediaId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAsrVoiceTranslateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAsrVoiceTranslateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAsrVoiceTranslateRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiAsrVoiceTranslateRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiAsrVoiceTranslateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.asr.voice.translate"
}
func (this *OapiAsrVoiceTranslateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAsrVoiceTranslateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAsrVoiceTranslateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAsrVoiceTranslateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAsrVoiceTranslateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["media_id"] = this.MediaId
	return txtParams
}
func (this *OapiAsrVoiceTranslateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MediaId, "mediaId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAsrVoiceTranslateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAsrVoiceTranslateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAsrVoiceTranslateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
