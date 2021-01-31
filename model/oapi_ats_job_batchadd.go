package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsJobBatchaddRequest() *OapiAtsJobBatchaddRequest {
	return &OapiAtsJobBatchaddRequest{}
}

type OapiAtsJobBatchaddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsJobBatchaddResponse
	BizCode         string
	Jobs            string
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsJobBatchaddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsJobBatchaddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsJobBatchaddRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsJobBatchaddRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsJobBatchaddRequest) SetJobs(jobs2 string) {
	this.Jobs = jobs2
}
func (this *OapiAtsJobBatchaddRequest) GetJobs() string {
	return this.Jobs
}
func (this *OapiAtsJobBatchaddRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAtsJobBatchaddRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAtsJobBatchaddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.job.batchadd"
}
func (this *OapiAtsJobBatchaddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsJobBatchaddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsJobBatchaddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsJobBatchaddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsJobBatchaddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["jobs"] = this.Jobs
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAtsJobBatchaddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsJobBatchaddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsJobBatchaddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type JobExtDataVo struct {
	HeadCount int64 `json:"head_count,omitempty"`
}
type JobAddressVo struct {
	Detail string `json:"detail,omitempty"`
	Name   string `json:"name,omitempty"`
}
type AtsAddJobParam struct {
	Address       JobAddressVo `json:"address,omitempty"`
	Campus        bool         `json:"campus,omitempty"`
	City          string       `json:"city,omitempty"`
	CreatorUserId string       `json:"creator_user_id,omitempty"`
	Description   string       `json:"description,omitempty"`
	District      string       `json:"district,omitempty"`
	ExtData       JobExtDataVo `json:"ext_data,omitempty"`
	JobNature     string       `json:"job_nature,omitempty"`
	MaxSalary     int64        `json:"max_salary,omitempty"`
	MinSalary     int64        `json:"min_salary,omitempty"`
	Name          string       `json:"name,omitempty"`
	Province      string       `json:"province,omitempty"`
	RequiredEdu   int64        `json:"required_edu,omitempty"`
}
type OapiAtsJobBatchaddResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type JobSimpleVo struct {
	JobCode string `json:"job_code,omitempty"`
	JobId   string `json:"job_id,omitempty"`
	Name    string `json:"name,omitempty"`
}
type BatchResultItemVO struct {
	ErrorCode string      `json:"error_code,omitempty"`
	ErrorMsg  string      `json:"error_msg,omitempty"`
	Index     int64       `json:"index,omitempty"`
	Item      JobSimpleVo `json:"item,omitempty"`
	Success   bool        `json:"success,omitempty"`
}
type BatchResultVo struct {
	FailedCount  int64               `json:"failed_count,omitempty"`
	Result       []BatchResultItemVO `json:"result,omitempty"`
	SuccessCount int64               `json:"success_count,omitempty"`
	TotalCount   int64               `json:"total_count,omitempty"`
}
