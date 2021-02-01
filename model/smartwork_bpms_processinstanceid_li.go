package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessinstanceidListRequest() *SmartworkBpmsProcessinstanceidListRequest {
	return &SmartworkBpmsProcessinstanceidListRequest{}
}

type SmartworkBpmsProcessinstanceidListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkBpmsProcessinstanceidListResponse
	Cursor          int64
	EndTime         int64
	ProcessCode     string
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *SmartworkBpmsProcessinstanceidListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetSize() int64 {
	return this.Size
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetUseridList() string {
	return this.UseridList
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.processinstanceid.list"
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessinstanceidListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["process_code"] = this.ProcessCode
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *SmartworkBpmsProcessinstanceidListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCode, "processCode"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessinstanceidListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessinstanceidListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
