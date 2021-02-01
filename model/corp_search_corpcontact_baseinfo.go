package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpSearchCorpcontactBaseinfoRequest() *CorpSearchCorpcontactBaseinfoRequest {
	return &CorpSearchCorpcontactBaseinfoRequest{}
}

type CorpSearchCorpcontactBaseinfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpSearchCorpcontactBaseinfoResponse
	Offset          int64
	Query           string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpSearchCorpcontactBaseinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpSearchCorpcontactBaseinfoRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpSearchCorpcontactBaseinfoRequest) SetQuery(query2 string) {
	this.Query = query2
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetQuery() string {
	return this.Query
}
func (this *CorpSearchCorpcontactBaseinfoRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetApiMethodName() string {
	return "dingtalk.corp.search.corpcontact.baseinfo"
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpSearchCorpcontactBaseinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpSearchCorpcontactBaseinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpSearchCorpcontactBaseinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["query"] = this.Query
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpSearchCorpcontactBaseinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpSearchCorpcontactBaseinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpSearchCorpcontactBaseinfoResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type GroupContactResult struct {
	AliTmpDept  string `json:"ali_tmp_dept,omitempty"`
	FlowerName  string `json:"flower_name,omitempty"`
	JobNumber   string `json:"job_number,omitempty"`
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Userid      string `json:"userid,omitempty"`
	WorkStation string `json:"work_station,omitempty"`
}
type PageResult struct {
	HasMore   bool                 `json:"has_more,omitempty"`
	Offset    int64                `json:"offset,omitempty"`
	Total     int64                `json:"total,omitempty"`
	ValueList []GroupContactResult `json:"value_list,omitempty"`
}
