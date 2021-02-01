package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkAttendsGetusergroupRequest() *SmartworkAttendsGetusergroupRequest {
	return &SmartworkAttendsGetusergroupRequest{}
}

type SmartworkAttendsGetusergroupRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkAttendsGetusergroupResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *SmartworkAttendsGetusergroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkAttendsGetusergroupRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *SmartworkAttendsGetusergroupRequest) GetUserid() string {
	return this.Userid
}
func (this *SmartworkAttendsGetusergroupRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.attends.getusergroup"
}
func (this *SmartworkAttendsGetusergroupRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkAttendsGetusergroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkAttendsGetusergroupRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkAttendsGetusergroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkAttendsGetusergroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkAttendsGetusergroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *SmartworkAttendsGetusergroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkAttendsGetusergroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkAttendsGetusergroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkAttendsGetusergroupResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
