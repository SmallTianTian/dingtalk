package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduAlumniGetRequest() *OapiEduAlumniGetRequest {
	return &OapiEduAlumniGetRequest{}
}

type OapiEduAlumniGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduAlumniGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduAlumniGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduAlumniGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduAlumniGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.alumni.get"
}
func (this *OapiEduAlumniGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduAlumniGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduAlumniGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduAlumniGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduAlumniGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiEduAlumniGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduAlumniGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduAlumniGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduAlumniGetResponse struct {
	taobao.TaobaoResponse
	Result  QueryAlumniResponse `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
type QueryAlumniResponse struct {
	CorpId string `json:"corp_id,omitempty"`
}
