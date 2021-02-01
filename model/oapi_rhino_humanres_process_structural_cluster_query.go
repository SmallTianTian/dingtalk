package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoHumanresProcessStructuralClusterQueryRequest() *OapiRhinoHumanresProcessStructuralClusterQueryRequest {
	return &OapiRhinoHumanresProcessStructuralClusterQueryRequest{}
}

type OapiRhinoHumanresProcessStructuralClusterQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiRhinoHumanresProcessStructuralClusterQueryResponse
	BizIdProcessList string
	TopHttpMethod    string
	TopResponseType  string
	Userid           string
}

func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) SetBizIdProcessList(bizIdProcessList2 string) {
	this.BizIdProcessList = bizIdProcessList2
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetBizIdProcessList() string {
	return this.BizIdProcessList
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.humanres.process.structural.cluster.query"
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id_process_list"] = this.BizIdProcessList
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.BizIdProcessList, 20, "bizIdProcessList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoHumanresProcessStructuralClusterQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoHumanresProcessStructuralClusterQueryResponse struct {
	taobao.TaobaoResponse

	ExternalMsgInfo string  `json:"external_msg_info,omitempty"`
	Hsfcode         int64   `json:"hsfcode,omitempty"`
	Model           []Model `json:"model,omitempty"`
	Success         bool    `json:"success,omitempty"`
}
type ProcessStructuralClusterDto struct {
	Accessory string `json:"accessory,omitempty"`
	Action    string `json:"action,omitempty"`
	Id        int64  `json:"id,omitempty"`
	Part      string `json:"part,omitempty"`
	Position  string `json:"position,omitempty"`
	Stitch    string `json:"stitch,omitempty"`
}
