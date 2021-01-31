package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseDetaildataListRequest() *OapiEduCourseDetaildataListRequest {
	return &OapiEduCourseDetaildataListRequest{}
}

type OapiEduCourseDetaildataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseDetaildataListResponse
	CategoryCode    string
	CourseCode      string
	Cursor          int64
	FactorCodes     string
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	UserCropid      string
	UserIds         string
}

func (this *OapiEduCourseDetaildataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseDetaildataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseDetaildataListRequest) SetCategoryCode(categoryCode2 string) {
	this.CategoryCode = categoryCode2
}
func (this *OapiEduCourseDetaildataListRequest) GetCategoryCode() string {
	return this.CategoryCode
}
func (this *OapiEduCourseDetaildataListRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseDetaildataListRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseDetaildataListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduCourseDetaildataListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduCourseDetaildataListRequest) SetFactorCodes(factorCodes2 string) {
	this.FactorCodes = factorCodes2
}
func (this *OapiEduCourseDetaildataListRequest) GetFactorCodes() string {
	return this.FactorCodes
}
func (this *OapiEduCourseDetaildataListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseDetaildataListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseDetaildataListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduCourseDetaildataListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduCourseDetaildataListRequest) SetUserCropid(userCropid2 string) {
	this.UserCropid = userCropid2
}
func (this *OapiEduCourseDetaildataListRequest) GetUserCropid() string {
	return this.UserCropid
}
func (this *OapiEduCourseDetaildataListRequest) SetUserIds(userIds2 string) {
	this.UserIds = userIds2
}
func (this *OapiEduCourseDetaildataListRequest) GetUserIds() string {
	return this.UserIds
}
func (this *OapiEduCourseDetaildataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.detaildata.list"
}
func (this *OapiEduCourseDetaildataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseDetaildataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseDetaildataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseDetaildataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseDetaildataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["category_code"] = this.CategoryCode
	txtParams["course_code"] = this.CourseCode
	txtParams["cursor"] = this.Cursor
	txtParams["factor_codes"] = this.FactorCodes
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	txtParams["user_cropid"] = this.UserCropid
	txtParams["user_ids"] = this.UserIds
	return txtParams
}
func (this *OapiEduCourseDetaildataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CategoryCode, "categoryCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseDetaildataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseDetaildataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseDetaildataListResponse struct {
	taobao.TaobaoResponse
	Result  PageQueryResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type CourseDetailDataDTO struct {
	CategoryBizKey string `json:"category_biz_key,omitempty"`
	CategoryCode   string `json:"category_code,omitempty"`
	CourseCode     string `json:"course_code,omitempty"`
	FactorCode     string `json:"factor_code,omitempty"`
	UserCropid     string `json:"user_cropid,omitempty"`
	Userid         string `json:"userid,omitempty"`
	Value          string `json:"value,omitempty"`
}
type PageQueryResponse struct {
	HasMore    bool                  `json:"has_more,omitempty"`
	List       []CourseDetailDataDTO `json:"list,omitempty"`
	NextCursor int64                 `json:"next_cursor,omitempty"`
}
