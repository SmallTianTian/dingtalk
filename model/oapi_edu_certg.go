package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCertGetRequest() *OapiEduCertGetRequest {
	return &OapiEduCertGetRequest{}
}

type OapiEduCertGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCertGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduCertGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCertGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCertGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduCertGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduCertGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.cert.get"
}
func (this *OapiEduCertGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCertGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCertGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCertGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCertGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduCertGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCertGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCertGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCertGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenQueryCertResponse `json:"result,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
type Certdata struct {
	CanCert    bool  `json:"can_cert,omitempty"`
	CertLevel  int64 `json:"cert_level,omitempty"`
	CertStatus int64 `json:"cert_status,omitempty"`
}
type OpenPracticalTaskData struct {
	Finish   bool   `json:"finish,omitempty"`
	TaskCode string `json:"task_code,omitempty"`
}
type OpenQueryCertResponse struct {
	CertDatas         []Certdata              `json:"cert_datas,omitempty"`
	CurrentCertLevel  int64                   `json:"current_cert_level,omitempty"`
	PracticalTaskData []OpenPracticalTaskData `json:"practical_task_data,omitempty"`
}
