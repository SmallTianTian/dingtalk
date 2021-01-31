package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGetcolumnvalRequest() *OapiAttendanceGetcolumnvalRequest {
	return &OapiAttendanceGetcolumnvalRequest{}
}

type OapiAttendanceGetcolumnvalRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetcolumnvalResponse
	ColumnIdList    string
	FromDate        time.Time
	ToDate          time.Time
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAttendanceGetcolumnvalRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetcolumnvalRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetcolumnvalRequest) SetColumnIdList(columnIdList2 string) {
	this.ColumnIdList = columnIdList2
}
func (this *OapiAttendanceGetcolumnvalRequest) GetColumnIdList() string {
	return this.ColumnIdList
}
func (this *OapiAttendanceGetcolumnvalRequest) SetFromDate(fromDate2 time.Time) {
	this.FromDate = fromDate2
}
func (this *OapiAttendanceGetcolumnvalRequest) GetFromDate() time.Time {
	return this.FromDate
}
func (this *OapiAttendanceGetcolumnvalRequest) SetToDate(toDate2 time.Time) {
	this.ToDate = toDate2
}
func (this *OapiAttendanceGetcolumnvalRequest) GetToDate() time.Time {
	return this.ToDate
}
func (this *OapiAttendanceGetcolumnvalRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceGetcolumnvalRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceGetcolumnvalRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getcolumnval"
}
func (this *OapiAttendanceGetcolumnvalRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetcolumnvalRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetcolumnvalRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetcolumnvalRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetcolumnvalRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["column_id_list"] = this.ColumnIdList
	txtParams["from_date"] = this.FromDate
	txtParams["to_date"] = this.ToDate
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceGetcolumnvalRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ColumnIdList, 20, "columnIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetcolumnvalRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetcolumnvalRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetcolumnvalResponse struct {
	taobao.TaobaoResponse
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
	Result  ColumnValListForTopVo `json:"result,omitempty"`
}
type ColumnDayAndVal struct {
	Date  time.Time `json:"date,omitempty"`
	Value string    `json:"value,omitempty"`
}
type ColumnValForTopVo struct {
	ColumnVals []ColumnDayAndVal `json:"column_vals,omitempty"`
	ColumnVo   ColumnForTopVo    `json:"column_vo,omitempty"`
	FixedValue string            `json:"fixed_value,omitempty"`
}
type ColumnValListForTopVo struct {
	ColumnVals []ColumnValForTopVo `json:"column_vals,omitempty"`
}
