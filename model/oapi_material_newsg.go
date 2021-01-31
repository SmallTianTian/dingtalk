package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialNewsGetRequest() *OapiMaterialNewsGetRequest {
	return &OapiMaterialNewsGetRequest{}
}

type OapiMaterialNewsGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialNewsGetResponse
	MediaId         string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialNewsGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialNewsGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialNewsGetRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiMaterialNewsGetRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiMaterialNewsGetRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialNewsGetRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialNewsGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.news.get"
}
func (this *OapiMaterialNewsGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialNewsGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialNewsGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialNewsGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialNewsGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["media_id"] = this.MediaId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialNewsGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MediaId, "mediaId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialNewsGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialNewsGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialNewsGetResponse struct {
	taobao.TaobaoResponse
	Articles   []ArticleDTO `json:"articles,omitempty"`
	Errcode    int64        `json:"errcode,omitempty"`
	Errmsg     string       `json:"errmsg,omitempty"`
	MediaId    string       `json:"media_id,omitempty"`
	UpdateTime int64        `json:"update_time,omitempty"`
}
