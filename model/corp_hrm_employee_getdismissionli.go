package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHrmEmployeeGetdismissionlistRequest() *CorpHrmEmployeeGetdismissionlistRequest {
	return &CorpHrmEmployeeGetdismissionlistRequest{}
}

type CorpHrmEmployeeGetdismissionlistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHrmEmployeeGetdismissionlistResponse
	Current         int64
	OpUserid        string
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpHrmEmployeeGetdismissionlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) SetCurrent(current2 int64) {
	this.Current = current2
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetCurrent() int64 {
	return this.Current
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetApiMethodName() string {
	return "dingtalk.corp.hrm.employee.getdismissionlist"
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current"] = this.Current
	txtParams["op_userid"] = this.OpUserid
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Current, "current"); err != nil {
		return err
	}
	return nil
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHrmEmployeeGetdismissionlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpHrmEmployeeGetdismissionlistResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type HrmApiDismissionModel struct {
	ConfirmJoinTime  time.Time `json:"confirm_join_time,omitempty"`
	DeptName         string    `json:"dept_name,omitempty"`
	DismissionMemo   string    `json:"dismission_memo,omitempty"`
	DismissionReason string    `json:"dismission_reason,omitempty"`
	Email            string    `json:"email,omitempty"`
	EmployeeStatus   string    `json:"employee_status,omitempty"`
	EmployeeType     string    `json:"employee_type,omitempty"`
	LastWorkDate     time.Time `json:"last_work_date,omitempty"`
	Name             string    `json:"name,omitempty"`
	Position         string    `json:"position,omitempty"`
	Userid           string    `json:"userid,omitempty"`
}
type HrmApiPage struct {
	Current   int64                   `json:"current,omitempty"`
	DataList  []HrmApiDismissionModel `json:"data_list,omitempty"`
	PageSize  int64                   `json:"page_size,omitempty"`
	Total     int64                   `json:"total,omitempty"`
	TotalPage int64                   `json:"total_page,omitempty"`
}
