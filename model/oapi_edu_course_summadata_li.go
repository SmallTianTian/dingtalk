package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseSummadataListRequest() *OapiEduCourseSummadataListRequest {
	return &OapiEduCourseSummadataListRequest{}
}

type OapiEduCourseSummadataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseSummadataListResponse
	CategoryCodes   string
	CourseCode      string
	Cursor          int64
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseSummadataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseSummadataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseSummadataListRequest) SetCategoryCodes(categoryCodes2 string) {
	this.CategoryCodes = categoryCodes2
}
func (this *OapiEduCourseSummadataListRequest) GetCategoryCodes() string {
	return this.CategoryCodes
}
func (this *OapiEduCourseSummadataListRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseSummadataListRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseSummadataListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduCourseSummadataListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduCourseSummadataListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseSummadataListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseSummadataListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduCourseSummadataListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduCourseSummadataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.summadata.list"
}
func (this *OapiEduCourseSummadataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseSummadataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseSummadataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseSummadataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseSummadataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["category_codes"] = this.CategoryCodes
	txtParams["course_code"] = this.CourseCode
	txtParams["cursor"] = this.Cursor
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiEduCourseSummadataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CategoryCodes, "categoryCodes"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseSummadataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseSummadataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseSummadataListResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  PageQueryResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type CourseSummaryDataDTO struct {
	CategoryBizKey string `json:"category_biz_key,omitempty"`
	CategoryCode   string `json:"category_code,omitempty"`
	CourseCode     string `json:"course_code,omitempty"`
	Data           string `json:"data,omitempty"`
}
