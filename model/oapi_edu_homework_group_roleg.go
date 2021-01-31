package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkGroupRoleGetRequest() *OapiEduHomeworkGroupRoleGetRequest {
	return &OapiEduHomeworkGroupRoleGetRequest{}
}

type OapiEduHomeworkGroupRoleGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkGroupRoleGetResponse
	BizCode         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduHomeworkGroupRoleGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkGroupRoleGetRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkGroupRoleGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.group.role.get"
}
func (this *OapiEduHomeworkGroupRoleGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkGroupRoleGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkGroupRoleGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduHomeworkGroupRoleGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkGroupRoleGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkGroupRoleGetResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
