package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiOrgSetoaurlRequest() *OapiOrgSetoaurlRequest {
	return &OapiOrgSetoaurlRequest{}
}

type OapiOrgSetoaurlRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgSetoaurlResponse
	OaTitle         string
	OaUrl           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgSetoaurlRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgSetoaurlRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgSetoaurlRequest) SetOaTitle(oaTitle2 string) {
	this.OaTitle = oaTitle2
}
func (this *OapiOrgSetoaurlRequest) GetOaTitle() string {
	return this.OaTitle
}
func (this *OapiOrgSetoaurlRequest) SetOaUrl(oaUrl2 string) {
	this.OaUrl = oaUrl2
}
func (this *OapiOrgSetoaurlRequest) GetOaUrl() string {
	return this.OaUrl
}
func (this *OapiOrgSetoaurlRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.setoaurl"
}
func (this *OapiOrgSetoaurlRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgSetoaurlRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgSetoaurlRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgSetoaurlRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgSetoaurlRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["oa_title"] = this.OaTitle
	txtParams["oa_url"] = this.OaUrl
	return txtParams
}
func (this *OapiOrgSetoaurlRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OaUrl, "oaUrl"); err != nil {
		return err
	}
	return nil
}
func (this *OapiOrgSetoaurlRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgSetoaurlRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgSetoaurlResponse struct {
	taobao.TaobaoResponse
}
