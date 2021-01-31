package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordCreateRequest() *OapiProcessWorkrecordCreateRequest {
	return &OapiProcessWorkrecordCreateRequest{}
}

type OapiProcessWorkrecordCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.create"
}
func (this *OapiProcessWorkrecordCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SaveFakeProcessInstanceRequest struct {
	Agentid             int64                  `json:"agentid,omitempty"`
	BizAction           string                 `json:"biz_action,omitempty"`
	CustomData          string                 `json:"custom_data,omitempty"`
	FormComponentValues []FormComponentValueVo `json:"form_component_values,omitempty"`
	MainInstanceId      string                 `json:"main_instance_id,omitempty"`
	OriginatorUserId    string                 `json:"originator_user_id,omitempty"`
	ProcessCode         string                 `json:"process_code,omitempty"`
	Remark              string                 `json:"remark,omitempty"`
	Title               string                 `json:"title,omitempty"`
	Url                 string                 `json:"url,omitempty"`
}
type OapiProcessWorkrecordCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64                           `json:"errcode,omitempty"`
	Errmsg  string                          `json:"errmsg,omitempty"`
	Result  SaveFaceProcessInstanceResponse `json:"result,omitempty"`
	Success bool                            `json:"success,omitempty"`
}
type SaveFaceProcessInstanceResponse struct {
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
}
