package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpDeptgroupSyncuserRequest() *CorpDeptgroupSyncuserRequest {
	return &CorpDeptgroupSyncuserRequest{}
}

type CorpDeptgroupSyncuserRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDeptgroupSyncuserResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpDeptgroupSyncuserRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDeptgroupSyncuserRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *CorpDeptgroupSyncuserRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *CorpDeptgroupSyncuserRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpDeptgroupSyncuserRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpDeptgroupSyncuserRequest) GetApiMethodName() string {
	return "dingtalk.corp.deptgroup.syncuser"
}
func (this *CorpDeptgroupSyncuserRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDeptgroupSyncuserRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDeptgroupSyncuserRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDeptgroupSyncuserRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDeptgroupSyncuserRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDeptgroupSyncuserRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpDeptgroupSyncuserRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpDeptgroupSyncuserRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDeptgroupSyncuserRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDeptgroupSyncuserResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
