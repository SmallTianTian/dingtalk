package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFugongHealthDataListRequest() *OapiFugongHealthDataListRequest {
	return &OapiFugongHealthDataListRequest{}
}

type OapiFugongHealthDataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiFugongHealthDataListResponse
	ActionDate        string
	Offset            int64
	ProcessInstanceId string
	Size              int64
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiFugongHealthDataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFugongHealthDataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFugongHealthDataListRequest) SetActionDate(actionDate2 string) {
	this.ActionDate = actionDate2
}
func (this *OapiFugongHealthDataListRequest) GetActionDate() string {
	return this.ActionDate
}
func (this *OapiFugongHealthDataListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiFugongHealthDataListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiFugongHealthDataListRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *OapiFugongHealthDataListRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *OapiFugongHealthDataListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiFugongHealthDataListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiFugongHealthDataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.fugong.health_data.list"
}
func (this *OapiFugongHealthDataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFugongHealthDataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFugongHealthDataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFugongHealthDataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFugongHealthDataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action_date"] = this.ActionDate
	txtParams["offset"] = this.Offset
	txtParams["process_instance_id"] = this.ProcessInstanceId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiFugongHealthDataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActionDate, "actionDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFugongHealthDataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFugongHealthDataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFugongHealthDataListResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
}
