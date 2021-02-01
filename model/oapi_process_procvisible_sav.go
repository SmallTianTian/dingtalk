package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessProcvisibleSaveRequest() *OapiProcessProcvisibleSaveRequest {
	return &OapiProcessProcvisibleSaveRequest{}
}

type OapiProcessProcvisibleSaveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessProcvisibleSaveResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessProcvisibleSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessProcvisibleSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessProcvisibleSaveRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessProcvisibleSaveRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessProcvisibleSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.procvisible.save"
}
func (this *OapiProcessProcvisibleSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessProcvisibleSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessProcvisibleSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessProcvisibleSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessProcvisibleSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessProcvisibleSaveRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessProcvisibleSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessProcvisibleSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Visiblelist struct {
	VisibleType  int64  `json:"visible_type,omitempty"`
	VisibleValue string `json:"visible_value,omitempty"`
}
type UpdateProcessVisibleRequest struct {
	Agentid     int64         `json:"agentid,omitempty"`
	ProcessCode string        `json:"process_code,omitempty"`
	Userid      string        `json:"userid,omitempty"`
	VisibleList []Visiblelist `json:"visible_list,omitempty"`
}
type OapiProcessProcvisibleSaveResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
