package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduPeriodMetadataListRequest() *OapiEduPeriodMetadataListRequest {
	return &OapiEduPeriodMetadataListRequest{}
}

type OapiEduPeriodMetadataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduPeriodMetadataListResponse
	AreaCode        string
	Cursor          int64
	DataOrderType   int64
	Level           int64
	OperatorUserid  string
	ParentId        int64
	Size            int64
	SortType        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduPeriodMetadataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduPeriodMetadataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduPeriodMetadataListRequest) SetAreaCode(areaCode2 string) {
	this.AreaCode = areaCode2
}
func (this *OapiEduPeriodMetadataListRequest) GetAreaCode() string {
	return this.AreaCode
}
func (this *OapiEduPeriodMetadataListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduPeriodMetadataListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduPeriodMetadataListRequest) SetDataOrderType(dataOrderType2 int64) {
	this.DataOrderType = dataOrderType2
}
func (this *OapiEduPeriodMetadataListRequest) GetDataOrderType() int64 {
	return this.DataOrderType
}
func (this *OapiEduPeriodMetadataListRequest) SetLevel(level2 int64) {
	this.Level = level2
}
func (this *OapiEduPeriodMetadataListRequest) GetLevel() int64 {
	return this.Level
}
func (this *OapiEduPeriodMetadataListRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduPeriodMetadataListRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduPeriodMetadataListRequest) SetParentId(parentId2 int64) {
	this.ParentId = parentId2
}
func (this *OapiEduPeriodMetadataListRequest) GetParentId() int64 {
	return this.ParentId
}
func (this *OapiEduPeriodMetadataListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduPeriodMetadataListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduPeriodMetadataListRequest) SetSortType(sortType2 int64) {
	this.SortType = sortType2
}
func (this *OapiEduPeriodMetadataListRequest) GetSortType() int64 {
	return this.SortType
}
func (this *OapiEduPeriodMetadataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.period.metadata.list"
}
func (this *OapiEduPeriodMetadataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduPeriodMetadataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduPeriodMetadataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduPeriodMetadataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduPeriodMetadataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["area_code"] = this.AreaCode
	txtParams["cursor"] = this.Cursor
	txtParams["data_order_type"] = this.DataOrderType
	txtParams["level"] = this.Level
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["parent_id"] = this.ParentId
	txtParams["size"] = this.Size
	txtParams["sort_type"] = this.SortType
	return txtParams
}
func (this *OapiEduPeriodMetadataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AreaCode, "areaCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduPeriodMetadataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduPeriodMetadataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduPeriodMetadataListResponse struct {
	taobao.TaobaoResponse
	Result  PageQueryResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type PeriodMetadataDTO struct {
	AreaCode   string `json:"area_code,omitempty"`
	Id         int64  `json:"id,omitempty"`
	Level      int64  `json:"level,omitempty"`
	ParentId   int64  `json:"parent_id,omitempty"`
	PeriodCode string `json:"period_code,omitempty"`
	PeriodName string `json:"period_name,omitempty"`
}
