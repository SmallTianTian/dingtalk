package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialArticlePublishRequest() *OapiMaterialArticlePublishRequest {
	return &OapiMaterialArticlePublishRequest{}
}

type OapiMaterialArticlePublishRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialArticlePublishResponse
	ArticleId       int64
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialArticlePublishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialArticlePublishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialArticlePublishRequest) SetArticleId(articleId2 int64) {
	this.ArticleId = articleId2
}
func (this *OapiMaterialArticlePublishRequest) GetArticleId() int64 {
	return this.ArticleId
}
func (this *OapiMaterialArticlePublishRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialArticlePublishRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialArticlePublishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.article.publish"
}
func (this *OapiMaterialArticlePublishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialArticlePublishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialArticlePublishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialArticlePublishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialArticlePublishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["article_id"] = this.ArticleId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialArticlePublishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ArticleId, "articleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialArticlePublishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialArticlePublishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialArticlePublishResponse struct {
	taobao.TaobaoResponse
	Url string `json:"url,omitempty"`
}
