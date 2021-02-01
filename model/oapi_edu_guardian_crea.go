package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGuardianCreateRequest() *OapiEduGuardianCreateRequest {
	return &OapiEduGuardianCreateRequest{}
}

type OapiEduGuardianCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGuardianCreateResponse
	BizId           string
	ClassId         int64
	Mobile          string
	Operator        string
	Relation        string
	StuId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGuardianCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGuardianCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGuardianCreateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiEduGuardianCreateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiEduGuardianCreateRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduGuardianCreateRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduGuardianCreateRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiEduGuardianCreateRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiEduGuardianCreateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduGuardianCreateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduGuardianCreateRequest) SetRelation(relation2 string) {
	this.Relation = relation2
}
func (this *OapiEduGuardianCreateRequest) GetRelation() string {
	return this.Relation
}
func (this *OapiEduGuardianCreateRequest) SetStuId(stuId2 string) {
	this.StuId = stuId2
}
func (this *OapiEduGuardianCreateRequest) GetStuId() string {
	return this.StuId
}
func (this *OapiEduGuardianCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.guardian.create"
}
func (this *OapiEduGuardianCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGuardianCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGuardianCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGuardianCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGuardianCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["class_id"] = this.ClassId
	txtParams["mobile"] = this.Mobile
	txtParams["operator"] = this.Operator
	txtParams["relation"] = this.Relation
	txtParams["stu_id"] = this.StuId
	return txtParams
}
func (this *OapiEduGuardianCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGuardianCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGuardianCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGuardianCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenEduUserCreateResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OpenEduUserCreateResponse struct {
	BizId  string `json:"biz_id,omitempty"`
	Userid string `json:"userid,omitempty"`
}
