package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduTeacherCreateRequest() *OapiEduTeacherCreateRequest {
	return &OapiEduTeacherCreateRequest{}
}

type OapiEduTeacherCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduTeacherCreateResponse
	Adviser         int64
	BizId           string
	ClassId         int64
	Operator        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduTeacherCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduTeacherCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduTeacherCreateRequest) SetAdviser(adviser2 int64) {
	this.Adviser = adviser2
}
func (this *OapiEduTeacherCreateRequest) GetAdviser() int64 {
	return this.Adviser
}
func (this *OapiEduTeacherCreateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiEduTeacherCreateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiEduTeacherCreateRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduTeacherCreateRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduTeacherCreateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduTeacherCreateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduTeacherCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduTeacherCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduTeacherCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.teacher.create"
}
func (this *OapiEduTeacherCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduTeacherCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduTeacherCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduTeacherCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduTeacherCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["adviser"] = this.Adviser
	txtParams["biz_id"] = this.BizId
	txtParams["class_id"] = this.ClassId
	txtParams["operator"] = this.Operator
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduTeacherCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Adviser, "adviser"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduTeacherCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduTeacherCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduTeacherCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenEduUserCreateResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
