package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHrmEmployeeDelresumerecordRequest() *OapiHrmEmployeeDelresumerecordRequest {
	return &OapiHrmEmployeeDelresumerecordRequest{}
}

type OapiHrmEmployeeDelresumerecordRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHrmEmployeeDelresumerecordResponse
	ResumeId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiHrmEmployeeDelresumerecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeDelresumerecordRequest) SetResumeId(resumeId2 string) {
	this.ResumeId = resumeId2
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetResumeId() string {
	return this.ResumeId
}
func (this *OapiHrmEmployeeDelresumerecordRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.delresumerecord"
}
func (this *OapiHrmEmployeeDelresumerecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeDelresumerecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeDelresumerecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["resume_id"] = this.ResumeId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiHrmEmployeeDelresumerecordRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ResumeId, "resumeId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeDelresumerecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeDelresumerecordResponse struct {
	taobao.TaobaoResponse
}
