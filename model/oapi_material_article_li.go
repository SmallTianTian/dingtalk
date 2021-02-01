package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialArticleListRequest() *OapiMaterialArticleListRequest {
	return &OapiMaterialArticleListRequest{}
}

type OapiMaterialArticleListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialArticleListResponse
	PageSize        int64
	PageStart       int64
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialArticleListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialArticleListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialArticleListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiMaterialArticleListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiMaterialArticleListRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiMaterialArticleListRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiMaterialArticleListRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialArticleListRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialArticleListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.article.list"
}
func (this *OapiMaterialArticleListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialArticleListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialArticleListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialArticleListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialArticleListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialArticleListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PageSize, "pageSize"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialArticleListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialArticleListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialArticleListResponse struct {
	taobao.TaobaoResponse

	ItemCount  int64        `json:"item_count,omitempty"`
	Items      []ArticleDTO `json:"items,omitempty"`
	TotalCount int64        `json:"total_count,omitempty"`
}
type ArticleDTO struct {
	ArticleId     int64  `json:"article_id,omitempty"`
	Content       string `json:"content,omitempty"`
	CreateTime    int64  `json:"create_time,omitempty"`
	Digest        string `json:"digest,omitempty"`
	PublishStatus int64  `json:"publish_status,omitempty"`
	PublishTime   int64  `json:"publish_time,omitempty"`
	ThumbMediaId  string `json:"thumb_media_id,omitempty"`
	Title         string `json:"title,omitempty"`
	UpdateTime    int64  `json:"update_time,omitempty"`
	Url           string `json:"url,omitempty"`
}
