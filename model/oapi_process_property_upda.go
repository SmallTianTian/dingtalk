package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessPropertyUpdateRequest() *OapiProcessPropertyUpdateRequest {
	return &OapiProcessPropertyUpdateRequest{}
}

type OapiProcessPropertyUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessPropertyUpdateResponse
	ComponentId     string
	ProcessCode     string
	Props           string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProcessPropertyUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessPropertyUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessPropertyUpdateRequest) SetComponentId(componentId2 string) {
	this.ComponentId = componentId2
}
func (this *OapiProcessPropertyUpdateRequest) GetComponentId() string {
	return this.ComponentId
}
func (this *OapiProcessPropertyUpdateRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessPropertyUpdateRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessPropertyUpdateRequest) SetProps(props2 string) {
	this.Props = props2
}
func (this *OapiProcessPropertyUpdateRequest) GetProps() string {
	return this.Props
}
func (this *OapiProcessPropertyUpdateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessPropertyUpdateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessPropertyUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.property.update"
}
func (this *OapiProcessPropertyUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessPropertyUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessPropertyUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessPropertyUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessPropertyUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["component_id"] = this.ComponentId
	txtParams["process_code"] = this.ProcessCode
	txtParams["props"] = this.Props
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessPropertyUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessPropertyUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessPropertyUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProcessTemplatePropsVo struct {
	PayEnable bool `json:"pay_enable,omitempty"`
}
type OapiProcessPropertyUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
