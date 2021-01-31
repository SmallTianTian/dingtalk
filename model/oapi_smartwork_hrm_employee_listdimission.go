package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeListdimissionRequest() *OapiSmartworkHrmEmployeeListdimissionRequest {
	return &OapiSmartworkHrmEmployeeListdimissionRequest{}
}

type OapiSmartworkHrmEmployeeListdimissionRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeListdimissionResponse
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.listdimission"
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.UseridList, 50, "useridList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeListdimissionRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeListdimissionResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  []EmpDimissionInfoVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type EmpDeptVO struct {
	DeptId   int64  `json:"dept_id,omitempty"`
	DeptPath string `json:"dept_path,omitempty"`
}
type EmpDimissionInfoVo struct {
	DeptList       []EmpDeptVO `json:"dept_list,omitempty"`
	HandoverUserid string      `json:"handover_userid,omitempty"`
	LastWorkDay    int64       `json:"last_work_day,omitempty"`
	MainDeptId     int64       `json:"main_dept_id,omitempty"`
	MainDeptName   string      `json:"main_dept_name,omitempty"`
	PreStatus      int64       `json:"pre_status,omitempty"`
	ReasonMemo     string      `json:"reason_memo,omitempty"`
	ReasonType     int64       `json:"reason_type,omitempty"`
	Status         int64       `json:"status,omitempty"`
	Userid         string      `json:"userid,omitempty"`
}
