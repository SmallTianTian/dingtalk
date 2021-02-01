package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialArticleUpdateRequest() *OapiMaterialArticleUpdateRequest {
	return &OapiMaterialArticleUpdateRequest{}
}

type OapiMaterialArticleUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialArticleUpdateResponse
	Article         string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialArticleUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialArticleUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialArticleUpdateRequest) SetArticle(article2 string) {
	this.Article = article2
}
func (this *OapiMaterialArticleUpdateRequest) GetArticle() string {
	return this.Article
}
func (this *OapiMaterialArticleUpdateRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialArticleUpdateRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialArticleUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.article.update"
}
func (this *OapiMaterialArticleUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialArticleUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialArticleUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialArticleUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialArticleUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["article"] = this.Article
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialArticleUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Unionid, "unionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialArticleUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialArticleUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialArticleUpdateResponse struct {
	taobao.TaobaoResponse
}
