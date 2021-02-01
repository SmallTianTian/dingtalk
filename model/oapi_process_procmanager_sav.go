package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessProcmanagerSaveRequest() *OapiProcessProcmanagerSaveRequest {
	return &OapiProcessProcmanagerSaveRequest{}
}

type OapiProcessProcmanagerSaveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessProcmanagerSaveResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessProcmanagerSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessProcmanagerSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessProcmanagerSaveRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessProcmanagerSaveRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessProcmanagerSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.procmanager.save"
}
func (this *OapiProcessProcmanagerSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessProcmanagerSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessProcmanagerSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessProcmanagerSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessProcmanagerSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessProcmanagerSaveRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessProcmanagerSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessProcmanagerSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UpdateProcessManagersRequest struct {
	Agentid     int64    `json:"agentid,omitempty"`
	ManagerList []string `json:"manager_list,omitempty"`
	ProcessCode string   `json:"process_code,omitempty"`
	Userid      string   `json:"userid,omitempty"`
}
type OapiProcessProcmanagerSaveResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
