package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduStudentListRequest() *OapiEduStudentListRequest {
	return &OapiEduStudentListRequest{}
}

type OapiEduStudentListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduStudentListResponse
	ClassId         int64
	PageNo          int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduStudentListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduStudentListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduStudentListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduStudentListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduStudentListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *OapiEduStudentListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *OapiEduStudentListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduStudentListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduStudentListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.student.list"
}
func (this *OapiEduStudentListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduStudentListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduStudentListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduStudentListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduStudentListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiEduStudentListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduStudentListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduStudentListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduStudentListResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
