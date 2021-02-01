package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiBlackboardListtoptenRequest() *OapiBlackboardListtoptenRequest {
	return &OapiBlackboardListtoptenRequest{}
}

type OapiBlackboardListtoptenRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardListtoptenResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiBlackboardListtoptenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardListtoptenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardListtoptenRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiBlackboardListtoptenRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiBlackboardListtoptenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.listtopten"
}
func (this *OapiBlackboardListtoptenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardListtoptenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardListtoptenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardListtoptenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardListtoptenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiBlackboardListtoptenRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiBlackboardListtoptenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardListtoptenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBlackboardListtoptenResponse struct {
	taobao.TaobaoResponse
	BlackboardList []OapiBlackboardVo `json:"blackboard_list,omitempty"`
}
