package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeAttachmentUpdateRequest() *OapiSmartworkHrmEmployeeAttachmentUpdateRequest {
	return &OapiSmartworkHrmEmployeeAttachmentUpdateRequest{}
}

type OapiSmartworkHrmEmployeeAttachmentUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeAttachmentUpdateResponse
	Agentid         string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) SetAgentid(agentid2 string) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetAgentid() string {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.attachment.update"
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeAttachmentUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EmpAttachmentUpdateParam struct {
	FieldCode  string `json:"field_code,omitempty"`
	FileSuffix string `json:"file_suffix,omitempty"`
	MediaId    string `json:"media_id,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type OapiSmartworkHrmEmployeeAttachmentUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
