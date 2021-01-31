package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServiceaccountListRequest() *OapiServiceaccountListRequest {
	return &OapiServiceaccountListRequest{}
}

type OapiServiceaccountListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceaccountListResponse
	PageSize        int64
	PageStart       int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceaccountListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceaccountListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceaccountListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiServiceaccountListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiServiceaccountListRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiServiceaccountListRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiServiceaccountListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.serviceaccount.list"
}
func (this *OapiServiceaccountListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceaccountListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceaccountListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceaccountListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceaccountListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["pageSize"] = this.PageSize
	txtParams["pageStart"] = this.PageStart
	return txtParams
}
func (this *OapiServiceaccountListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxValue(int(this.PageSize), 50, "pageSize"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServiceaccountListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceaccountListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceaccountListResponse struct {
	taobao.TaobaoResponse
	Errcode    int64          `json:"errcode,omitempty"`
	Errmsg     string         `json:"errmsg,omitempty"`
	ItemCount  int64          `json:"item_count,omitempty"`
	Items      []PublisherDTO `json:"items,omitempty"`
	TotalCount int64          `json:"total_count,omitempty"`
}
type PublisherDTO struct {
	AvatarMediaId  string `json:"avatar_media_id,omitempty"`
	Brief          string `json:"brief,omitempty"`
	Desc           string `json:"desc,omitempty"`
	Name           string `json:"name,omitempty"`
	PreviewMediaId string `json:"preview_media_id,omitempty"`
	Status         string `json:"status,omitempty"`
	Unionid        string `json:"unionid,omitempty"`
}
