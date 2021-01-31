package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduTeacherListRequest() *OapiEduTeacherListRequest {
	return &OapiEduTeacherListRequest{}
}

type OapiEduTeacherListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduTeacherListResponse
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduTeacherListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduTeacherListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduTeacherListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduTeacherListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduTeacherListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.teacher.list"
}
func (this *OapiEduTeacherListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduTeacherListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduTeacherListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduTeacherListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduTeacherListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	return txtParams
}
func (this *OapiEduTeacherListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduTeacherListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduTeacherListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduTeacherListResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  []TeacherRespone `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
