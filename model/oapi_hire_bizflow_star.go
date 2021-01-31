package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHireBizflowStartRequest() *OapiHireBizflowStartRequest {
	return &OapiHireBizflowStartRequest{}
}

type OapiHireBizflowStartRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireBizflowStartResponse
	JobId           string
	OpUserid        string
	ResumeId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHireBizflowStartRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireBizflowStartRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireBizflowStartRequest) SetJobId(jobId2 string) {
	this.JobId = jobId2
}
func (this *OapiHireBizflowStartRequest) GetJobId() string {
	return this.JobId
}
func (this *OapiHireBizflowStartRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiHireBizflowStartRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiHireBizflowStartRequest) SetResumeId(resumeId2 string) {
	this.ResumeId = resumeId2
}
func (this *OapiHireBizflowStartRequest) GetResumeId() string {
	return this.ResumeId
}
func (this *OapiHireBizflowStartRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.bizflow.start"
}
func (this *OapiHireBizflowStartRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireBizflowStartRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireBizflowStartRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireBizflowStartRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireBizflowStartRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["job_id"] = this.JobId
	txtParams["op_userid"] = this.OpUserid
	txtParams["resume_id"] = this.ResumeId
	return txtParams
}
func (this *OapiHireBizflowStartRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.JobId, "jobId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHireBizflowStartRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireBizflowStartRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireBizflowStartResponse struct {
	taobao.TaobaoResponse
	Result TopStartBizFlowResult `json:"result,omitempty"`
}
type TopStartBizFlowResult struct {
	MobileJumpUrl string `json:"mobile_jump_url,omitempty"`
	PcJumpUrl     string `json:"pc_jump_url,omitempty"`
	PcRedirectUrl string `json:"pc_redirect_url,omitempty"`
}
