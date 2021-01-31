package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiSmartworkHrmEmployeeOnjoblistQueryRequest() *OapiSmartworkHrmEmployeeOnjoblistQueryRequest {
	return &OapiSmartworkHrmEmployeeOnjoblistQueryRequest{}
}

type OapiSmartworkHrmEmployeeOnjoblistQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeOnjoblistQueryResponse
	Cursor          int64
	SearchParam     string
	Size            int64
	StatusList      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetSearchParam(searchParam2 string) {
	this.SearchParam = searchParam2
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetSearchParam() string {
	return this.SearchParam
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetStatusList(statusList2 string) {
	this.StatusList = statusList2
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetStatusList() string {
	return this.StatusList
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.onjoblist.query"
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["search_param"] = this.SearchParam
	txtParams["size"] = this.Size
	txtParams["status_list"] = this.StatusList
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeOnjoblistQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EmpSearchParamVo struct {
	EntryEndDate   time.Time `json:"entry_end_date,omitempty"`
	EntryStartDate time.Time `json:"entry_start_date,omitempty"`
}
type OapiSmartworkHrmEmployeeOnjoblistQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
