package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeV2ListRequest() *OapiSmartworkHrmEmployeeV2ListRequest {
	return &OapiSmartworkHrmEmployeeV2ListRequest{}
}

type OapiSmartworkHrmEmployeeV2ListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeV2ListResponse
	Agentid         int64
	FieldFilterList string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) SetFieldFilterList(fieldFilterList2 string) {
	this.FieldFilterList = fieldFilterList2
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetFieldFilterList() string {
	return this.FieldFilterList
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.v2.list"
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["field_filter_list"] = this.FieldFilterList
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeV2ListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeV2ListResponse struct {
	taobao.TaobaoResponse
	Errcode int64              `json:"errcode,omitempty"`
	Errmsg  string             `json:"errmsg,omitempty"`
	Result  []EmpRosterFieldVo `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type FieldValueVo struct {
	ItemIndex int64  `json:"item_index,omitempty"`
	Label     string `json:"label,omitempty"`
	Value     string `json:"value,omitempty"`
}
type EmpFieldDataVo struct {
	FieldCode      string         `json:"field_code,omitempty"`
	FieldName      string         `json:"field_name,omitempty"`
	FieldValueList []FieldValueVo `json:"field_value_list,omitempty"`
	GroupId        string         `json:"group_id,omitempty"`
}
type EmpRosterFieldVo struct {
	CorpId        string           `json:"corp_id,omitempty"`
	FieldDataList []EmpFieldDataVo `json:"field_data_list,omitempty"`
	Unionid       string           `json:"unionid,omitempty"`
	Userid        string           `json:"userid,omitempty"`
}
