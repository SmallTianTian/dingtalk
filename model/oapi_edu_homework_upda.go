package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkUpdateRequest() *OapiEduHomeworkUpdateRequest {
	return &OapiEduHomeworkUpdateRequest{}
}

type OapiEduHomeworkUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkUpdateResponse
	BizCode         string
	HwId            int64
	Identifier      string
	Status          string
	TeacherUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduHomeworkUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkUpdateRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkUpdateRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkUpdateRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkUpdateRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkUpdateRequest) SetIdentifier(identifier2 string) {
	this.Identifier = identifier2
}
func (this *OapiEduHomeworkUpdateRequest) GetIdentifier() string {
	return this.Identifier
}
func (this *OapiEduHomeworkUpdateRequest) SetStatus(status2 string) {
	this.Status = status2
}
func (this *OapiEduHomeworkUpdateRequest) GetStatus() string {
	return this.Status
}
func (this *OapiEduHomeworkUpdateRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduHomeworkUpdateRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduHomeworkUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.update"
}
func (this *OapiEduHomeworkUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["hw_id"] = this.HwId
	txtParams["identifier"] = this.Identifier
	txtParams["status"] = this.Status
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduHomeworkUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  int64  `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
