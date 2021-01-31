package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeListRequest() *OapiSmartworkHrmEmployeeListRequest {
	return &OapiSmartworkHrmEmployeeListRequest{}
}

type OapiSmartworkHrmEmployeeListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeListResponse
	Agentid         int64
	FieldFilterList string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartworkHrmEmployeeListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeListRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeListRequest) SetFieldFilterList(fieldFilterList2 string) {
	this.FieldFilterList = fieldFilterList2
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetFieldFilterList() string {
	return this.FieldFilterList
}
func (this *OapiSmartworkHrmEmployeeListRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.list"
}
func (this *OapiSmartworkHrmEmployeeListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["field_filter_list"] = this.FieldFilterList
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.FieldFilterList, 100, "fieldFilterList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeListResponse struct {
	taobao.TaobaoResponse
	Result  []EmpFieldInfoVO `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type EmpFieldVO struct {
	FieldCode string `json:"field_code,omitempty"`
	FieldName string `json:"field_name,omitempty"`
	GroupId   string `json:"group_id,omitempty"`
	Label     string `json:"label,omitempty"`
	Value     string `json:"value,omitempty"`
}
type EmpFieldInfoVO struct {
	FieldList []EmpFieldVO `json:"field_list,omitempty"`
	Partner   bool         `json:"partner,omitempty"`
	Userid    string       `json:"userid,omitempty"`
}
