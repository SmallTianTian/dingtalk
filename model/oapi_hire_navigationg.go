package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHireNavigationGetRequest() *OapiHireNavigationGetRequest {
	return &OapiHireNavigationGetRequest{}
}

type OapiHireNavigationGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireNavigationGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiHireNavigationGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireNavigationGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireNavigationGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHireNavigationGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHireNavigationGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.navigation.get"
}
func (this *OapiHireNavigationGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireNavigationGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireNavigationGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireNavigationGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireNavigationGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiHireNavigationGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiHireNavigationGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireNavigationGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireNavigationGetResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
