package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseListRequest() *OapiEduCourseListRequest {
	return &OapiEduCourseListRequest{}
}

type OapiEduCourseListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseListResponse
	Cursor          int64
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduCourseListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduCourseListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduCourseListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduCourseListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.list"
}
func (this *OapiEduCourseListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiEduCourseListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseListResponse struct {
	taobao.TaobaoResponse
	Result  ListCourseResponse `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type CourseVO struct {
	BizKey        string `json:"biz_key,omitempty"`
	Code          string `json:"code,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	Introduce     string `json:"introduce,omitempty"`
	Name          string `json:"name,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	TeacherCorpid string `json:"teacher_corpid,omitempty"`
	TeacherUserid string `json:"teacher_userid,omitempty"`
}
type ListCourseResponse struct {
	HasMore    bool       `json:"has_more,omitempty"`
	List       []CourseVO `json:"list,omitempty"`
	NextCursor int64      `json:"next_cursor,omitempty"`
}
