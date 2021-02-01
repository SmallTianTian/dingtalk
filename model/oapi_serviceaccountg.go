package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServiceaccountGetRequest() *OapiServiceaccountGetRequest {
	return &OapiServiceaccountGetRequest{}
}

type OapiServiceaccountGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceaccountGetResponse
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiServiceaccountGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceaccountGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceaccountGetRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiServiceaccountGetRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiServiceaccountGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.serviceaccount.get"
}
func (this *OapiServiceaccountGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceaccountGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceaccountGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceaccountGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceaccountGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiServiceaccountGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Unionid, "unionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServiceaccountGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceaccountGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceaccountGetResponse struct {
	taobao.TaobaoResponse

	ServiceAccount ServiceAccountDTO `json:"service_account,omitempty"`
}
type ServiceAccountDTO struct {
	AvatarMediaId  string `json:"avatar_media_id,omitempty"`
	Brief          string `json:"brief,omitempty"`
	Desc           string `json:"desc,omitempty"`
	Name           string `json:"name,omitempty"`
	PreviewMediaId string `json:"preview_media_id,omitempty"`
	Status         string `json:"status,omitempty"`
	Unionid        string `json:"unionid,omitempty"`
}
