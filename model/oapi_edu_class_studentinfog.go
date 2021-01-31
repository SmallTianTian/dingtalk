package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduClassStudentinfoGetRequest() *OapiEduClassStudentinfoGetRequest {
	return &OapiEduClassStudentinfoGetRequest{}
}

type OapiEduClassStudentinfoGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassStudentinfoGetResponse
	AppId           int64
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduClassStudentinfoGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassStudentinfoGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassStudentinfoGetRequest) SetAppId(appId2 int64) {
	this.AppId = appId2
}
func (this *OapiEduClassStudentinfoGetRequest) GetAppId() int64 {
	return this.AppId
}
func (this *OapiEduClassStudentinfoGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduClassStudentinfoGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduClassStudentinfoGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduClassStudentinfoGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduClassStudentinfoGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.studentinfo.get"
}
func (this *OapiEduClassStudentinfoGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassStudentinfoGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassStudentinfoGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassStudentinfoGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassStudentinfoGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["class_id"] = this.ClassId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduClassStudentinfoGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduClassStudentinfoGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassStudentinfoGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduClassStudentinfoGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenStudentSelectDto `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type OpenPatriarchSelectDto struct {
	Avatar       string `json:"avatar,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
	Name         string `json:"name,omitempty"`
	Relation     string `json:"relation,omitempty"`
	RelationName string `json:"relation_name,omitempty"`
	Userid       string `json:"userid,omitempty"`
}
type OpenStudentSelectDto struct {
	Avatar     string                   `json:"avatar,omitempty"`
	ClassId    int64                    `json:"class_id,omitempty"`
	Guardians  []OpenPatriarchSelectDto `json:"guardians,omitempty"`
	Name       string                   `json:"name,omitempty"`
	StudentNum string                   `json:"student_num,omitempty"`
	Userid     string                   `json:"userid,omitempty"`
}
