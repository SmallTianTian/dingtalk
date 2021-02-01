package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduRolesGetRequest() *OapiEduRolesGetRequest {
	return &OapiEduRolesGetRequest{}
}

type OapiEduRolesGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduRolesGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduRolesGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduRolesGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduRolesGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduRolesGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduRolesGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.roles.get"
}
func (this *OapiEduRolesGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduRolesGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduRolesGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduRolesGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduRolesGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduRolesGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduRolesGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduRolesGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduRolesGetResponse struct {
	taobao.TaobaoResponse
	Result  QueryUserRolesResponse `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type QueryUserRolesResponse struct {
	Advisor  []int64 `json:"advisor,omitempty"`
	Guardian []int64 `json:"guardian,omitempty"`
	Student  []int64 `json:"student,omitempty"`
	Teacher  []int64 `json:"teacher,omitempty"`
	Userid   string  `json:"userid,omitempty"`
}
