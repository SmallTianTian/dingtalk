package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingtaxUserPushRequest() *OapiDingtaxUserPushRequest {
	return &OapiDingtaxUserPushRequest{}
}

type OapiDingtaxUserPushRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingtaxUserPushResponse
	SourceRegion    string
	TopHttpMethod   string
	TopResponseType string
	UserInfoList    string
}

func (this *OapiDingtaxUserPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingtaxUserPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingtaxUserPushRequest) SetSourceRegion(sourceRegion2 string) {
	this.SourceRegion = sourceRegion2
}
func (this *OapiDingtaxUserPushRequest) GetSourceRegion() string {
	return this.SourceRegion
}
func (this *OapiDingtaxUserPushRequest) SetUserInfoList(userInfoList2 string) {
	this.UserInfoList = userInfoList2
}
func (this *OapiDingtaxUserPushRequest) GetUserInfoList() string {
	return this.UserInfoList
}
func (this *OapiDingtaxUserPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingtax.user.push"
}
func (this *OapiDingtaxUserPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingtaxUserPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingtaxUserPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingtaxUserPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingtaxUserPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["source_region"] = this.SourceRegion
	txtParams["user_info_list"] = this.UserInfoList
	return txtParams
}
func (this *OapiDingtaxUserPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.SourceRegion, "sourceRegion"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingtaxUserPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingtaxUserPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DingTaxUserInfoDTO struct {
	NewUser      string `json:"new_user,omitempty"`
	TaxationType string `json:"taxation_type,omitempty"`
	UserMobile   string `json:"user_mobile,omitempty"`
	UserRole     string `json:"user_role,omitempty"`
}
type OapiDingtaxUserPushResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
