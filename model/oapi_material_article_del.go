package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialArticleDeleteRequest() *OapiMaterialArticleDeleteRequest {
	return &OapiMaterialArticleDeleteRequest{}
}

type OapiMaterialArticleDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialArticleDeleteResponse
	ArticleId       int64
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialArticleDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialArticleDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialArticleDeleteRequest) SetArticleId(articleId2 int64) {
	this.ArticleId = articleId2
}
func (this *OapiMaterialArticleDeleteRequest) GetArticleId() int64 {
	return this.ArticleId
}
func (this *OapiMaterialArticleDeleteRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialArticleDeleteRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialArticleDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.article.delete"
}
func (this *OapiMaterialArticleDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialArticleDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialArticleDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialArticleDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialArticleDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["article_id"] = this.ArticleId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialArticleDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ArticleId, "articleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialArticleDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialArticleDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialArticleDeleteResponse struct {
	taobao.TaobaoResponse
}
