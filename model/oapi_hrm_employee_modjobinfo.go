package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHrmEmployeeModjobinfoRequest() *OapiHrmEmployeeModjobinfoRequest {
	return &OapiHrmEmployeeModjobinfoRequest{}
}

type OapiHrmEmployeeModjobinfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHrmEmployeeModjobinfoResponse
	HrmApiJobModel  string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHrmEmployeeModjobinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeModjobinfoRequest) SetHrmApiJobModel(hrmApiJobModel2 string) {
	this.HrmApiJobModel = hrmApiJobModel2
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetHrmApiJobModel() string {
	return this.HrmApiJobModel
}
func (this *OapiHrmEmployeeModjobinfoRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.modjobinfo"
}
func (this *OapiHrmEmployeeModjobinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeModjobinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeModjobinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["hrm_api_job_model"] = this.HrmApiJobModel
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiHrmEmployeeModjobinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeModjobinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeModjobinfoResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
