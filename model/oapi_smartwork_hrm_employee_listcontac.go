package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeListcontactRequest() *OapiSmartworkHrmEmployeeListcontactRequest {
	return &OapiSmartworkHrmEmployeeListcontactRequest{}
}

type OapiSmartworkHrmEmployeeListcontactRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeListcontactResponse
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.listcontact"
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UseridList, "useridList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeListcontactRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeListcontactResponse struct {
	taobao.TaobaoResponse
	Result  []EmpFieldInfoVo `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type EmpFieldVo struct {
	FieldCode string `json:"field_code,omitempty"`
	FieldName string `json:"field_name,omitempty"`
	GroupId   string `json:"group_id,omitempty"`
	Label     string `json:"label,omitempty"`
	Value     string `json:"value,omitempty"`
}
type EmpFieldInfoVo struct {
	FieldList []EmpFieldVo `json:"field_list,omitempty"`
	Userid    string       `json:"userid,omitempty"`
}
