package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordForwardCreateRequest() *OapiProcessWorkrecordForwardCreateRequest {
	return &OapiProcessWorkrecordForwardCreateRequest{}
}

type OapiProcessWorkrecordForwardCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordForwardCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordForwardCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordForwardCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.forward.create"
}
func (this *OapiProcessWorkrecordForwardCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordForwardCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordForwardCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordForwardCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordForwardCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AddForwardRequest struct {
	Agentid           int64    `json:"agentid,omitempty"`
	ProcessInstanceId string   `json:"process_instance_id,omitempty"`
	UseridList        []string `json:"userid_list,omitempty"`
}
type OapiProcessWorkrecordForwardCreateResponse struct {
	taobao.TaobaoResponse
}
