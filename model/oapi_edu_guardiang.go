package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGuardianGetRequest() *OapiEduGuardianGetRequest {
	return &OapiEduGuardianGetRequest{}
}

type OapiEduGuardianGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGuardianGetResponse
	ClassId         int64
	GuardianUserid  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGuardianGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGuardianGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGuardianGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduGuardianGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduGuardianGetRequest) SetGuardianUserid(guardianUserid2 string) {
	this.GuardianUserid = guardianUserid2
}
func (this *OapiEduGuardianGetRequest) GetGuardianUserid() string {
	return this.GuardianUserid
}
func (this *OapiEduGuardianGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.guardian.get"
}
func (this *OapiEduGuardianGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGuardianGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGuardianGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGuardianGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGuardianGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["guardian_userid"] = this.GuardianUserid
	return txtParams
}
func (this *OapiEduGuardianGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGuardianGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGuardianGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGuardianGetResponse struct {
	taobao.TaobaoResponse
	Result  GuardianRespone `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type Relations struct {
	Nick          string `json:"nick,omitempty"`
	Relation      string `json:"relation,omitempty"`
	StudentUserid string `json:"student_userid,omitempty"`
}
type GuardianRespone struct {
	GuardianUserid string      `json:"guardian_userid,omitempty"`
	Nick           string      `json:"nick,omitempty"`
	Relation       string      `json:"relation,omitempty"`
	Relations      []Relations `json:"relations,omitempty"`
	StudentUserid  string      `json:"student_userid,omitempty"`
}
