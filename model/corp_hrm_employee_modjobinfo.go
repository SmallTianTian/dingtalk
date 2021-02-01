package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHrmEmployeeModjobinfoRequest() *CorpHrmEmployeeModjobinfoRequest {
	return &CorpHrmEmployeeModjobinfoRequest{}
}

type CorpHrmEmployeeModjobinfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHrmEmployeeModjobinfoResponse
	HrmApiJobModel  string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpHrmEmployeeModjobinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHrmEmployeeModjobinfoRequest) SetHrmApiJobModel(hrmApiJobModel2 string) {
	this.HrmApiJobModel = hrmApiJobModel2
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetHrmApiJobModel() string {
	return this.HrmApiJobModel
}
func (this *CorpHrmEmployeeModjobinfoRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetApiMethodName() string {
	return "dingtalk.corp.hrm.employee.modjobinfo"
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHrmEmployeeModjobinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHrmEmployeeModjobinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHrmEmployeeModjobinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["hrm_api_job_model"] = this.HrmApiJobModel
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *CorpHrmEmployeeModjobinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHrmEmployeeModjobinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type HrmApiJobModel struct {
	BirthTime           time.Time `json:"birth_time,omitempty"`
	ConfirmJoinTime     time.Time `json:"confirm_join_time,omitempty"`
	EmployeeStatus      int64     `json:"employee_status,omitempty"`
	EmployeeType        int64     `json:"employee_type,omitempty"`
	JoinWorkingTime     time.Time `json:"join_working_time,omitempty"`
	ProbationPeriodType int64     `json:"probation_period_type,omitempty"`
	RegularTime         time.Time `json:"regular_time,omitempty"`
	Userid              string    `json:"userid,omitempty"`
}
type CorpHrmEmployeeModjobinfoResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
