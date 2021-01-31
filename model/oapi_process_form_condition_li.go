package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessFormConditionListRequest() *OapiProcessFormConditionListRequest {
	return &OapiProcessFormConditionListRequest{}
}

type OapiProcessFormConditionListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessFormConditionListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessFormConditionListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessFormConditionListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessFormConditionListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessFormConditionListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessFormConditionListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.form.condition.list"
}
func (this *OapiProcessFormConditionListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessFormConditionListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessFormConditionListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessFormConditionListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessFormConditionListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessFormConditionListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessFormConditionListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessFormConditionListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type QureyProcessRequest struct {
	Agentid     int64  `json:"agentid,omitempty"`
	ProcessCode string `json:"process_code,omitempty"`
}
type OapiProcessFormConditionListResponse struct {
	taobao.TaobaoResponse
	List []SimpleFormComponentVo `json:"list,omitempty"`
}
type SimpleFormComponentVo struct {
	Id    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
}
