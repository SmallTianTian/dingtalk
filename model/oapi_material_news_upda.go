package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialNewsUpdateRequest() *OapiMaterialNewsUpdateRequest {
	return &OapiMaterialNewsUpdateRequest{}
}

type OapiMaterialNewsUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialNewsUpdateResponse
	Articles        string
	MediaId         string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialNewsUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialNewsUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialNewsUpdateRequest) SetArticles(articles2 string) {
	this.Articles = articles2
}
func (this *OapiMaterialNewsUpdateRequest) GetArticles() string {
	return this.Articles
}
func (this *OapiMaterialNewsUpdateRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiMaterialNewsUpdateRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiMaterialNewsUpdateRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialNewsUpdateRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialNewsUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.news.update"
}
func (this *OapiMaterialNewsUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialNewsUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialNewsUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialNewsUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialNewsUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["articles"] = this.Articles
	txtParams["media_id"] = this.MediaId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialNewsUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Articles, 8, "articles"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialNewsUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialNewsUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialNewsUpdateResponse struct {
	taobao.TaobaoResponse
}
