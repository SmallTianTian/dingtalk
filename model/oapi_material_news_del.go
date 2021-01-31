package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialNewsDeleteRequest() *OapiMaterialNewsDeleteRequest {
	return &OapiMaterialNewsDeleteRequest{}
}

type OapiMaterialNewsDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialNewsDeleteResponse
	MediaId         string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialNewsDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialNewsDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialNewsDeleteRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiMaterialNewsDeleteRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiMaterialNewsDeleteRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialNewsDeleteRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialNewsDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.news.delete"
}
func (this *OapiMaterialNewsDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialNewsDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialNewsDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialNewsDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialNewsDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["media_id"] = this.MediaId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialNewsDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MediaId, "mediaId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialNewsDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialNewsDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialNewsDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
