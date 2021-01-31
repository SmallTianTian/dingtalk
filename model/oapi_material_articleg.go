package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialArticleGetRequest() *OapiMaterialArticleGetRequest {
	return &OapiMaterialArticleGetRequest{}
}

type OapiMaterialArticleGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialArticleGetResponse
	ArticleId       int64
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialArticleGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialArticleGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialArticleGetRequest) SetArticleId(articleId2 int64) {
	this.ArticleId = articleId2
}
func (this *OapiMaterialArticleGetRequest) GetArticleId() int64 {
	return this.ArticleId
}
func (this *OapiMaterialArticleGetRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialArticleGetRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialArticleGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.article.get"
}
func (this *OapiMaterialArticleGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialArticleGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialArticleGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialArticleGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialArticleGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["article_id"] = this.ArticleId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialArticleGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ArticleId, "articleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialArticleGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialArticleGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialArticleGetResponse struct {
	taobao.TaobaoResponse
	ArticleId     int64  `json:"article_id,omitempty"`
	Content       string `json:"content,omitempty"`
	CreateTime    int64  `json:"create_time,omitempty"`
	Digest        string `json:"digest,omitempty"`
	Errcode       int64  `json:"errcode,omitempty"`
	Errmsg        string `json:"errmsg,omitempty"`
	PublishStatus int64  `json:Securitytaobao.PUBLISH_STATU,omitemptyS`
	PublishTime   int64  `json:"publish_time,omitempty"`
	ThumbMediaId  string `json:"thumb_media_id,omitempty"`
	Title         string `json:"title,omitempty"`
	UpdateTime    int64  `json:"update_time,omitempty"`
	Url           string `json:"url,omitempty"`
}
