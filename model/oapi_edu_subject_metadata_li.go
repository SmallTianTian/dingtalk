package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubjectMetadataListRequest() *OapiEduSubjectMetadataListRequest {
	return &OapiEduSubjectMetadataListRequest{}
}

type OapiEduSubjectMetadataListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubjectMetadataListResponse
	AreaCode        string
	Cursor          int64
	DataOrderType   int64
	Level           int64
	OperatorUserid  string
	ParentId        int64
	PeriodCode      string
	Size            int64
	SortType        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubjectMetadataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubjectMetadataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubjectMetadataListRequest) SetAreaCode(areaCode2 string) {
	this.AreaCode = areaCode2
}
func (this *OapiEduSubjectMetadataListRequest) GetAreaCode() string {
	return this.AreaCode
}
func (this *OapiEduSubjectMetadataListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduSubjectMetadataListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduSubjectMetadataListRequest) SetDataOrderType(dataOrderType2 int64) {
	this.DataOrderType = dataOrderType2
}
func (this *OapiEduSubjectMetadataListRequest) GetDataOrderType() int64 {
	return this.DataOrderType
}
func (this *OapiEduSubjectMetadataListRequest) SetLevel(level2 int64) {
	this.Level = level2
}
func (this *OapiEduSubjectMetadataListRequest) GetLevel() int64 {
	return this.Level
}
func (this *OapiEduSubjectMetadataListRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduSubjectMetadataListRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduSubjectMetadataListRequest) SetParentId(parentId2 int64) {
	this.ParentId = parentId2
}
func (this *OapiEduSubjectMetadataListRequest) GetParentId() int64 {
	return this.ParentId
}
func (this *OapiEduSubjectMetadataListRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduSubjectMetadataListRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduSubjectMetadataListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduSubjectMetadataListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduSubjectMetadataListRequest) SetSortType(sortType2 int64) {
	this.SortType = sortType2
}
func (this *OapiEduSubjectMetadataListRequest) GetSortType() int64 {
	return this.SortType
}
func (this *OapiEduSubjectMetadataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.subject.metadata.list"
}
func (this *OapiEduSubjectMetadataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubjectMetadataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubjectMetadataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubjectMetadataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubjectMetadataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["area_code"] = this.AreaCode
	txtParams["cursor"] = this.Cursor
	txtParams["data_order_type"] = this.DataOrderType
	txtParams["level"] = this.Level
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["parent_id"] = this.ParentId
	txtParams["period_code"] = this.PeriodCode
	txtParams["size"] = this.Size
	txtParams["sort_type"] = this.SortType
	return txtParams
}
func (this *OapiEduSubjectMetadataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AreaCode, "areaCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubjectMetadataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubjectMetadataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduSubjectMetadataListResponse struct {
	taobao.TaobaoResponse
	Result  PageQueryResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type SubjectMetadataDTO struct {
	AreaCode    string `json:"area_code,omitempty"`
	Id          int64  `json:"id,omitempty"`
	Level       int64  `json:"level,omitempty"`
	ParentId    int64  `json:"parent_id,omitempty"`
	PeriodCode  string `json:"period_code,omitempty"`
	SubjectCode string `json:"subject_code,omitempty"`
	SubjectName string `json:"subject_name,omitempty"`
}
