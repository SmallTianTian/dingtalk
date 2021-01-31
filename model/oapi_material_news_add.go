package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialNewsAddRequest() *OapiMaterialNewsAddRequest {
	return &OapiMaterialNewsAddRequest{}
}

type OapiMaterialNewsAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialNewsAddResponse
	Articles        string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialNewsAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialNewsAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialNewsAddRequest) SetArticles(articles2 string) {
	this.Articles = articles2
}
func (this *OapiMaterialNewsAddRequest) GetArticles() string {
	return this.Articles
}
func (this *OapiMaterialNewsAddRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialNewsAddRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialNewsAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.news.add"
}
func (this *OapiMaterialNewsAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialNewsAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialNewsAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialNewsAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialNewsAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["articles"] = this.Articles
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialNewsAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Articles, 8, "articles"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialNewsAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialNewsAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialNewsAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}
