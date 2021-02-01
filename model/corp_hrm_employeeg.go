package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHrmEmployeeGetRequest() *CorpHrmEmployeeGetRequest {
	return &CorpHrmEmployeeGetRequest{}
}

type CorpHrmEmployeeGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHrmEmployeeGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpHrmEmployeeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHrmEmployeeGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpHrmEmployeeGetRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpHrmEmployeeGetRequest) GetApiMethodName() string {
	return "dingtalk.corp.hrm.employee.get"
}
func (this *CorpHrmEmployeeGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHrmEmployeeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHrmEmployeeGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHrmEmployeeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHrmEmployeeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHrmEmployeeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpHrmEmployeeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *CorpHrmEmployeeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHrmEmployeeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpHrmEmployeeGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type SysCustomFieldVpo struct {
	ConfirmEntryRequired   bool   `json:"confirm_entry_required,omitempty"`
	ContactClientFlag      bool   `json:"contact_client_flag,omitempty"`
	ContactFlag            bool   `json:"contact_flag,omitempty"`
	ContactSource          int64  `json:"contact_source,omitempty"`
	ContactSystemFlag      bool   `json:"contact_system_flag,omitempty"`
	CorpId                 string `json:"corp_id,omitempty"`
	Deleted                bool   `json:"deleted,omitempty"`
	EditableByEmp          bool   `json:"editable_by_emp,omitempty"`
	EditableByHr           bool   `json:"editable_by_hr,omitempty"`
	EmpProfileRequired     bool   `json:"emp_profile_required,omitempty"`
	FieldCode              string `json:"field_code,omitempty"`
	FieldName              string `json:"field_name,omitempty"`
	FieldType              string `json:"field_type,omitempty"`
	GroupId                string `json:"group_id,omitempty"`
	HiddenFromEmployeeFlag bool   `json:"hidden_from_employee_flag,omitempty"`
	Hint                   string `json:"hint,omitempty"`
	Key                    string `json:"key,omitempty"`
	NoWatermark            bool   `json:"no_watermark,omitempty"`
	OptionText             string `json:"option_text,omitempty"`
	Required               bool   `json:"required,omitempty"`
	SystemFlag             bool   `json:"system_flag,omitempty"`
	TypeName               string `json:"type_name,omitempty"`
	Value                  string `json:"value,omitempty"`
	VisibleByEmp           bool   `json:"visible_by_emp,omitempty"`
}
type FieldGroupVpo struct {
	FieldList []SysCustomFieldVpo `json:"field_list,omitempty"`
	GroupId   string              `json:"group_id,omitempty"`
	GroupName string              `json:"group_name,omitempty"`
}
