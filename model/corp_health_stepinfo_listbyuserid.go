package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHealthStepinfoListbyuseridRequest() *CorpHealthStepinfoListbyuseridRequest {
	return &CorpHealthStepinfoListbyuseridRequest{}
}

type CorpHealthStepinfoListbyuseridRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHealthStepinfoListbyuseridResponse
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *CorpHealthStepinfoListbyuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHealthStepinfoListbyuseridRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetStatDate() string {
	return this.StatDate
}
func (this *CorpHealthStepinfoListbyuseridRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetUserids() string {
	return this.Userids
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetApiMethodName() string {
	return "dingtalk.corp.health.stepinfo.listbyuserid"
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHealthStepinfoListbyuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHealthStepinfoListbyuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHealthStepinfoListbyuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["stat_date"] = this.StatDate
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *CorpHealthStepinfoListbyuseridRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StatDate, "statDate"); err != nil {
		return err
	}
	return nil
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHealthStepinfoListbyuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpHealthStepinfoListbyuseridResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
