package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiTrainingGroupinfoGetRequest() *OapiTrainingGroupinfoGetRequest {
	return &OapiTrainingGroupinfoGetRequest{}
}

type OapiTrainingGroupinfoGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTrainingGroupinfoGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTrainingGroupinfoGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTrainingGroupinfoGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTrainingGroupinfoGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiTrainingGroupinfoGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiTrainingGroupinfoGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.training.groupinfo.get"
}
func (this *OapiTrainingGroupinfoGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTrainingGroupinfoGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTrainingGroupinfoGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTrainingGroupinfoGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTrainingGroupinfoGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiTrainingGroupinfoGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiTrainingGroupinfoGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTrainingGroupinfoGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type RequestVO struct {
	CorpId string   `json:"corp_id,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}
type OapiTrainingGroupinfoGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64     `json:"errcode,omitempty"`
	Errmsg  string    `json:"errmsg,omitempty"`
	Result  GroupInfo `json:"result,omitempty"`
}
