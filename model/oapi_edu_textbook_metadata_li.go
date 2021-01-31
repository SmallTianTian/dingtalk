package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduTextbookMetadataListRequest() *OapiEduTextbookMetadataListRequest {
	return &OapiEduTextbookMetadataListRequest{}
}

type OapiEduTextbookMetadataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduTextbookMetadataListResponse
	Cursor          int64
	DataOrderType   int64
	Level           int64
	OpUserId        string
	ParentId        int64
	Size            int64
	SortType        int64
	SubjectCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduTextbookMetadataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduTextbookMetadataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduTextbookMetadataListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduTextbookMetadataListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduTextbookMetadataListRequest) SetDataOrderType(dataOrderType2 int64) {
	this.DataOrderType = dataOrderType2
}
func (this *OapiEduTextbookMetadataListRequest) GetDataOrderType() int64 {
	return this.DataOrderType
}
func (this *OapiEduTextbookMetadataListRequest) SetLevel(level2 int64) {
	this.Level = level2
}
func (this *OapiEduTextbookMetadataListRequest) GetLevel() int64 {
	return this.Level
}
func (this *OapiEduTextbookMetadataListRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiEduTextbookMetadataListRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiEduTextbookMetadataListRequest) SetParentId(parentId2 int64) {
	this.ParentId = parentId2
}
func (this *OapiEduTextbookMetadataListRequest) GetParentId() int64 {
	return this.ParentId
}
func (this *OapiEduTextbookMetadataListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduTextbookMetadataListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduTextbookMetadataListRequest) SetSortType(sortType2 int64) {
	this.SortType = sortType2
}
func (this *OapiEduTextbookMetadataListRequest) GetSortType() int64 {
	return this.SortType
}
func (this *OapiEduTextbookMetadataListRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduTextbookMetadataListRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduTextbookMetadataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.textbook.metadata.list"
}
func (this *OapiEduTextbookMetadataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduTextbookMetadataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduTextbookMetadataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduTextbookMetadataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduTextbookMetadataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["data_order_type"] = this.DataOrderType
	txtParams["level"] = this.Level
	txtParams["op_user_id"] = this.OpUserId
	txtParams["parent_id"] = this.ParentId
	txtParams["size"] = this.Size
	txtParams["sort_type"] = this.SortType
	txtParams["subject_code"] = this.SubjectCode
	return txtParams
}
func (this *OapiEduTextbookMetadataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduTextbookMetadataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduTextbookMetadataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduTextbookMetadataListResponse struct {
	taobao.TaobaoResponse
	Result  PageQueryResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type TextbookMetadataDTO struct {
	Id           int64  `json:"id,omitempty"`
	Level        int64  `json:"level,omitempty"`
	ParentId     int64  `json:"parent_id,omitempty"`
	Remark       string `json:"remark,omitempty"`
	Status       int64  `json:"status,omitempty"`
	SubjectCode  string `json:"subject_code,omitempty"`
	TextbookCode string `json:"textbook_code,omitempty"`
	TextbookName string `json:"textbook_name,omitempty"`
}
