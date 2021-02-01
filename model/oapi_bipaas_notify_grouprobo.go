package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiBipaasNotifyGrouprobotRequest() *OapiBipaasNotifyGrouprobotRequest {
	return &OapiBipaasNotifyGrouprobotRequest{}
}

type OapiBipaasNotifyGrouprobotRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBipaasNotifyGrouprobotResponse
	RobotNotify     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBipaasNotifyGrouprobotRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBipaasNotifyGrouprobotRequest) SetRobotNotify(robotNotify2 string) {
	this.RobotNotify = robotNotify2
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetRobotNotify() string {
	return this.RobotNotify
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetApiMethodName() string {
	return "dingtalk.oapi.bipaas.notify.grouprobot"
}
func (this *OapiBipaasNotifyGrouprobotRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBipaasNotifyGrouprobotRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBipaasNotifyGrouprobotRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["robot_notify"] = this.RobotNotify
	return txtParams
}
func (this *OapiBipaasNotifyGrouprobotRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBipaasNotifyGrouprobotRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupRobotNotifyDto struct {
	AntcloudOperatorIds []string `json:"antcloud_operator_ids,omitempty"`
	AntcloudTenantId    string   `json:"antcloud_tenant_id,omitempty"`
	MessageBody         string   `json:"message_body,omitempty"`
	MessageType         string   `json:"message_type,omitempty"`
	RobotUrl            string   `json:"robot_url,omitempty"`
}
type OapiBipaasNotifyGrouprobotResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
