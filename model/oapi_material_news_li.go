package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMaterialNewsListRequest() *OapiMaterialNewsListRequest {
	return &OapiMaterialNewsListRequest{}
}

type OapiMaterialNewsListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMaterialNewsListResponse
	PageSize        int64
	PageStart       int64
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMaterialNewsListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMaterialNewsListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMaterialNewsListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiMaterialNewsListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiMaterialNewsListRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiMaterialNewsListRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiMaterialNewsListRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMaterialNewsListRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMaterialNewsListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.material.news.list"
}
func (this *OapiMaterialNewsListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMaterialNewsListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMaterialNewsListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMaterialNewsListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMaterialNewsListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMaterialNewsListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxValue(int(this.PageSize), 20, "pageSize"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMaterialNewsListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMaterialNewsListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMaterialNewsListResponse struct {
	taobao.TaobaoResponse

	ItemCount  int64         `json:"item_count,omitempty"`
	Items      []NewsCardDTO `json:"items,omitempty"`
	TotalCount int64         `json:"total_count,omitempty"`
}
type NewsCardDTO struct {
	Articles   []ArticleDTO `json:"articles,omitempty"`
	MediaId    string       `json:"media_id,omitempty"`
	UpdateTime int64        `json:"update_time,omitempty"`
}
