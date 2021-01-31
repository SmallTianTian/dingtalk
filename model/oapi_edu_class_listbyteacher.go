package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduClassListbyteacherRequest() *OapiEduClassListbyteacherRequest {
	return &OapiEduClassListbyteacherRequest{}
}

type OapiEduClassListbyteacherRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassListbyteacherResponse
	FilterParam     string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduClassListbyteacherRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassListbyteacherRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassListbyteacherRequest) SetFilterParam(filterParam2 string) {
	this.FilterParam = filterParam2
}
func (this *OapiEduClassListbyteacherRequest) GetFilterParam() string {
	return this.FilterParam
}
func (this *OapiEduClassListbyteacherRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduClassListbyteacherRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduClassListbyteacherRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.listbyteacher"
}
func (this *OapiEduClassListbyteacherRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassListbyteacherRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassListbyteacherRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassListbyteacherRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassListbyteacherRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["filter_param"] = this.FilterParam
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduClassListbyteacherRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduClassListbyteacherRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassListbyteacherRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenClassParam struct {
	GroupType string `json:"group_type,omitempty"`
	Role      string `json:"role,omitempty"`
}
type OapiEduClassListbyteacherResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
