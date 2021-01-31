package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessinstanceListRequest() *SmartworkBpmsProcessinstanceListRequest {
	return &SmartworkBpmsProcessinstanceListRequest{}
}

type SmartworkBpmsProcessinstanceListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkBpmsProcessinstanceListResponse
	Cursor          int64
	EndTime         int64
	ProcessCode     string
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *SmartworkBpmsProcessinstanceListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetSize() int64 {
	return this.Size
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetUseridList() string {
	return this.UseridList
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.processinstance.list"
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessinstanceListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["process_code"] = this.ProcessCode
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *SmartworkBpmsProcessinstanceListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCode, "processCode"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessinstanceListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessinstanceListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
