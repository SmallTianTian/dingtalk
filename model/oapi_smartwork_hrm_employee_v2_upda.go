package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeV2UpdateRequest() *OapiSmartworkHrmEmployeeV2UpdateRequest {
	return &OapiSmartworkHrmEmployeeV2UpdateRequest{}
}

type OapiSmartworkHrmEmployeeV2UpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeV2UpdateResponse
	Agentid         int64
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.v2.update"
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeV2UpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EmpListFieldVo struct {
	OldIndex int64        `json:"old_index,omitempty"`
	Section  []EmpFieldVo `json:"section,omitempty"`
}
type EmpGroupFieldVo struct {
	GroupId  string           `json:"group_id,omitempty"`
	Sections []EmpListFieldVo `json:"sections,omitempty"`
}
type EmpUpdateByCustomParam struct {
	Groups []EmpGroupFieldVo `json:"groups,omitempty"`
	Userid string            `json:"userid,omitempty"`
}
type OapiSmartworkHrmEmployeeV2UpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
