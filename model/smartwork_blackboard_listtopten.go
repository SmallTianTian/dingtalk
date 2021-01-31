package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBlackboardListtoptenRequest() *SmartworkBlackboardListtoptenRequest {
	return &SmartworkBlackboardListtoptenRequest{}
}

type SmartworkBlackboardListtoptenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkBlackboardListtoptenResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *SmartworkBlackboardListtoptenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBlackboardListtoptenRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *SmartworkBlackboardListtoptenRequest) GetUserid() string {
	return this.Userid
}
func (this *SmartworkBlackboardListtoptenRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.blackboard.listtopten"
}
func (this *SmartworkBlackboardListtoptenRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBlackboardListtoptenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBlackboardListtoptenRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBlackboardListtoptenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBlackboardListtoptenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBlackboardListtoptenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *SmartworkBlackboardListtoptenRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBlackboardListtoptenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBlackboardListtoptenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBlackboardListtoptenResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
