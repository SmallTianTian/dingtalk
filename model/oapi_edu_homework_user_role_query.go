package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkUserRoleQueryRequest() *OapiEduHomeworkUserRoleQueryRequest {
	return &OapiEduHomeworkUserRoleQueryRequest{}
}

type OapiEduHomeworkUserRoleQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkUserRoleQueryResponse
	BizCode         string
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduHomeworkUserRoleQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkUserRoleQueryRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkUserRoleQueryRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkUserRoleQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.user.role.query"
}
func (this *OapiEduHomeworkUserRoleQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkUserRoleQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkUserRoleQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduHomeworkUserRoleQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkUserRoleQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkUserRoleQueryResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
