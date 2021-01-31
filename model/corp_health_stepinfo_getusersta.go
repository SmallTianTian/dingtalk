package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHealthStepinfoGetuserstatusRequest() *CorpHealthStepinfoGetuserstatusRequest {
	return &CorpHealthStepinfoGetuserstatusRequest{}
}

type CorpHealthStepinfoGetuserstatusRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHealthStepinfoGetuserstatusResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpHealthStepinfoGetuserstatusRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHealthStepinfoGetuserstatusRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetApiMethodName() string {
	return "dingtalk.corp.health.stepinfo.getuserstatus"
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHealthStepinfoGetuserstatusRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHealthStepinfoGetuserstatusRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHealthStepinfoGetuserstatusRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpHealthStepinfoGetuserstatusRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHealthStepinfoGetuserstatusRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpHealthStepinfoGetuserstatusResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
