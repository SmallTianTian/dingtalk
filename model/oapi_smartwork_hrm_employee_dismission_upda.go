package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeDismissionUpdateRequest() *OapiSmartworkHrmEmployeeDismissionUpdateRequest {
	return &OapiSmartworkHrmEmployeeDismissionUpdateRequest{}
}

type OapiSmartworkHrmEmployeeDismissionUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeDismissionUpdateResponse
	Agentid         string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) SetAgentid(agentid2 string) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetAgentid() string {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.dismission.update"
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeDismissionUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EmpDismissionUpdateParam struct {
	DismissionMemo             string    `json:"dismission_memo,omitempty"`
	LastWorkDate               time.Time `json:"last_work_date,omitempty"`
	Partner                    bool      `json:"partner,omitempty"`
	TerminationReasonPassive   []string  `json:"terminationReasonPassive,omitempty"`
	TerminationReasonVoluntary []string  `json:"terminationReasonVoluntary,omitempty"`
	Userid                     string    `json:"userid,omitempty"`
}
type OapiSmartworkHrmEmployeeDismissionUpdateResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
	Suc    bool `json:"suc,omitempty"`
}
