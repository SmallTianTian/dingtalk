package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgpaasOrgInfoGetRequest() *OapiOrgpaasOrgInfoGetRequest {
	return &OapiOrgpaasOrgInfoGetRequest{}
}

type OapiOrgpaasOrgInfoGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgpaasOrgInfoGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgpaasOrgInfoGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgpaasOrgInfoGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgpaasOrgInfoGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.orgpaas.org.info.get"
}
func (this *OapiOrgpaasOrgInfoGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgpaasOrgInfoGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgpaasOrgInfoGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgpaasOrgInfoGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgpaasOrgInfoGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiOrgpaasOrgInfoGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgpaasOrgInfoGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgpaasOrgInfoGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgpaasOrgInfoGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  GetOrgInfoResp `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type GetOrgInfoResp struct {
	Extension string `json:"extension,omitempty"`
	OrgName   string `json:"org_name,omitempty"`
}
