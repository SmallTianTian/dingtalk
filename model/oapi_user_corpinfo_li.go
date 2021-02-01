package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserCorpinfoListRequest() *OapiUserCorpinfoListRequest {
	return &OapiUserCorpinfoListRequest{}
}

type OapiUserCorpinfoListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserCorpinfoListResponse
	CorpName        string
	Mobile          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserCorpinfoListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserCorpinfoListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserCorpinfoListRequest) SetCorpName(corpName2 string) {
	this.CorpName = corpName2
}
func (this *OapiUserCorpinfoListRequest) GetCorpName() string {
	return this.CorpName
}
func (this *OapiUserCorpinfoListRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiUserCorpinfoListRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiUserCorpinfoListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.corpinfo.list"
}
func (this *OapiUserCorpinfoListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserCorpinfoListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserCorpinfoListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserCorpinfoListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserCorpinfoListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_name"] = this.CorpName
	txtParams["mobile"] = this.Mobile
	return txtParams
}
func (this *OapiUserCorpinfoListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpName, "corpName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserCorpinfoListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserCorpinfoListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserCorpinfoListResponse struct {
	taobao.TaobaoResponse
	CorpList []CorpInfoVo `json:"corp_list,omitempty"`
}
type CorpInfoVo struct {
	AuthStatus int64  `json:"auth_status,omitempty"`
	CorpName   string `json:"corp_name,omitempty"`
	Corpid     string `json:"corpid,omitempty"`
}
