package model

import (
	"os"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMediaUploadRequest() *OapiMediaUploadRequest {
	return &OapiMediaUploadRequest{}
}

type OapiMediaUploadRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMediaUploadResponse
	Media           os.File
	TopResponseType string
	Type            string
}

func (this *OapiMediaUploadRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMediaUploadRequest) SetMedia(media2 os.File) {
	this.Media = media2
}
func (this *OapiMediaUploadRequest) GetMedia() os.File {
	return this.Media
}
func (this *OapiMediaUploadRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiMediaUploadRequest) GetType() string {
	return this.Type
}
func (this *OapiMediaUploadRequest) GetApiMethodName() string {
	return "dingtalk.oapi.media.upload"
}
func (this *OapiMediaUploadRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMediaUploadRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMediaUploadRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMediaUploadRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiMediaUploadRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Media, "media"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMediaUploadRequest) GetFileParams() map[string]os.File {
	params := make(map[string]os.File)
	params["media"] = this.Media
	return params
}
func (this *OapiMediaUploadRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMediaUploadRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMediaUploadResponse struct {
	taobao.TaobaoResponse
	CreatedAt int64  `json:"created_at,omitempty"`
	Errcode   int64  `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	MediaId   string `json:"media_id,omitempty"`
	Type      string `json:"type,omitempty"`
}
