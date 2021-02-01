package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduClassStudentidGetRequest() *OapiEduClassStudentidGetRequest {
	return &OapiEduClassStudentidGetRequest{}
}

type OapiEduClassStudentidGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassStudentidGetResponse
	AppId           int64
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduClassStudentidGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassStudentidGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassStudentidGetRequest) SetAppId(appId2 int64) {
	this.AppId = appId2
}
func (this *OapiEduClassStudentidGetRequest) GetAppId() int64 {
	return this.AppId
}
func (this *OapiEduClassStudentidGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduClassStudentidGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduClassStudentidGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduClassStudentidGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduClassStudentidGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.studentid.get"
}
func (this *OapiEduClassStudentidGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassStudentidGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassStudentidGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassStudentidGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassStudentidGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["class_id"] = this.ClassId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduClassStudentidGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppId, "appId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduClassStudentidGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassStudentidGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduClassStudentidGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenEduSelectStudentIdResponse `json:"result,omitempty"`
	Success bool                           `json:"success,omitempty"`
}
type OpenEduSelectStudentIdResponse struct {
	ClassId    int64    `json:"class_id,omitempty"`
	StudentIds []string `json:"student_ids,omitempty"`
}
