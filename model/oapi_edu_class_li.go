package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduClassListRequest() *OapiEduClassListRequest {
	return &OapiEduClassListRequest{}
}

type OapiEduClassListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassListResponse
	GradeId         int64
	PageNo          int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduClassListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassListRequest) SetGradeId(gradeId2 int64) {
	this.GradeId = gradeId2
}
func (this *OapiEduClassListRequest) GetGradeId() int64 {
	return this.GradeId
}
func (this *OapiEduClassListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *OapiEduClassListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *OapiEduClassListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduClassListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduClassListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.list"
}
func (this *OapiEduClassListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["grade_id"] = this.GradeId
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiEduClassListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GradeId, "gradeId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduClassListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduClassListResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
