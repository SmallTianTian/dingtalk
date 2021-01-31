package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubjectListRequest() *OapiEduSubjectListRequest {
	return &OapiEduSubjectListRequest{}
}

type OapiEduSubjectListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubjectListResponse
	Cursor          int64
	DataOrderType   int64
	OperatorUserid  string
	PeriodCode      string
	Size            int64
	SortType        int64
	SubjectCodeList string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubjectListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubjectListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubjectListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduSubjectListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduSubjectListRequest) SetDataOrderType(dataOrderType2 int64) {
	this.DataOrderType = dataOrderType2
}
func (this *OapiEduSubjectListRequest) GetDataOrderType() int64 {
	return this.DataOrderType
}
func (this *OapiEduSubjectListRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduSubjectListRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduSubjectListRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduSubjectListRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduSubjectListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduSubjectListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduSubjectListRequest) SetSortType(sortType2 int64) {
	this.SortType = sortType2
}
func (this *OapiEduSubjectListRequest) GetSortType() int64 {
	return this.SortType
}
func (this *OapiEduSubjectListRequest) SetSubjectCodeList(subjectCodeList2 string) {
	this.SubjectCodeList = subjectCodeList2
}
func (this *OapiEduSubjectListRequest) GetSubjectCodeList() string {
	return this.SubjectCodeList
}
func (this *OapiEduSubjectListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.subject.list"
}
func (this *OapiEduSubjectListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubjectListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubjectListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubjectListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubjectListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["data_order_type"] = this.DataOrderType
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["period_code"] = this.PeriodCode
	txtParams["size"] = this.Size
	txtParams["sort_type"] = this.SortType
	txtParams["subject_code_list"] = this.SubjectCodeList
	return txtParams
}
func (this *OapiEduSubjectListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubjectListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubjectListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduSubjectListResponse struct {
	taobao.TaobaoResponse
	Result  PageQueryResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
