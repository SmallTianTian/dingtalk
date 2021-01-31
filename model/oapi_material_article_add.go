package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialArticleAddRequest() *OapiMaterialArticleAddRequest {
	return &OapiMaterialArticleAddRequest{}
}

type OapiMaterialArticleAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialArticleAddResponse
	Article         string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialArticleAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialArticleAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialArticleAddRequest) SetArticle(article2 string) {
	this.Article = article2
}
func (this *OapiMaterialArticleAddRequest) GetArticle() string {
	return this.Article
}
func (this *OapiMaterialArticleAddRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialArticleAddRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialArticleAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.article.add"
}
func (this *OapiMaterialArticleAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialArticleAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialArticleAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialArticleAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialArticleAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["article"] = this.Article
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialArticleAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Unionid, "unionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialArticleAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialArticleAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ArticleCreateDTO struct {
	Content      string `json:"content,omitempty"`
	Digest       string `json:"digest,omitempty"`
	ThumbMediaId string `json:"thumb_media_id,omitempty"`
	Title        string `json:"title,omitempty"`
	Uuid         string `json:"uuid,omitempty"`
}
type OapiMaterialArticleAddResponse struct {
	taobao.TaobaoResponse
	ArticleId int64  `json:"article_id,omitempty"`
	Errcode   int64  `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	Success   bool   `json:"success,omitempty"`
}
