package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeUpdateRequest() *OapiSmartworkHrmEmployeeUpdateRequest {
	return &OapiSmartworkHrmEmployeeUpdateRequest{}
}

type OapiSmartworkHrmEmployeeUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeUpdateResponse
	Agentid         int64
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.update"
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EmpListFieldVO struct {
	Section []EmpFieldVo `json:"section,omitempty"`
}
type OapiSmartworkHrmEmployeeUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
