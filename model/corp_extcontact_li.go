package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpExtcontactListRequest() *CorpExtcontactListRequest {
	return &CorpExtcontactListRequest{}
}

type CorpExtcontactListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtcontactListResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtcontactListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtcontactListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpExtcontactListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpExtcontactListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpExtcontactListRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpExtcontactListRequest) GetApiMethodName() string {
	return "dingtalk.corp.extcontact.list"
}
func (this *CorpExtcontactListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtcontactListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtcontactListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtcontactListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtcontactListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtcontactListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpExtcontactListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxValue(int(this.Size), 100, "size"); err != nil {
		return err
	}
	return nil
}
func (this *CorpExtcontactListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtcontactListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtcontactListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
