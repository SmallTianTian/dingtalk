package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHealthStepinfoListbyuseridRequest() *OapiHealthStepinfoListbyuseridRequest {
	return &OapiHealthStepinfoListbyuseridRequest{}
}

type OapiHealthStepinfoListbyuseridRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHealthStepinfoListbyuseridResponse
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiHealthStepinfoListbyuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHealthStepinfoListbyuseridRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiHealthStepinfoListbyuseridRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.health.stepinfo.listbyuserid"
}
func (this *OapiHealthStepinfoListbyuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHealthStepinfoListbyuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHealthStepinfoListbyuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["stat_date"] = this.StatDate
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiHealthStepinfoListbyuseridRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StatDate, "statDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHealthStepinfoListbyuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHealthStepinfoListbyuseridResponse struct {
	taobao.TaobaoResponse
	Errcode      int64             `json:"errcode,omitempty"`
	Errmsg       string            `json:"errmsg,omitempty"`
	StepinfoList []BasicStepInfoVo `json:"stepinfo_list,omitempty"`
}
