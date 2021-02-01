package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCcoserviceServicegroupAddmemberRequest() *CcoserviceServicegroupAddmemberRequest {
	return &CcoserviceServicegroupAddmemberRequest{}
}

type CcoserviceServicegroupAddmemberRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CcoserviceServicegroupAddmemberResponse
	OpenGroupId     string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CcoserviceServicegroupAddmemberRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CcoserviceServicegroupAddmemberRequest) SetOpenGroupId(openGroupId2 string) {
	this.OpenGroupId = openGroupId2
}
func (this *CcoserviceServicegroupAddmemberRequest) GetOpenGroupId() string {
	return this.OpenGroupId
}
func (this *CcoserviceServicegroupAddmemberRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CcoserviceServicegroupAddmemberRequest) GetUserid() string {
	return this.Userid
}
func (this *CcoserviceServicegroupAddmemberRequest) GetApiMethodName() string {
	return "dingtalk.ccoservice.servicegroup.addmember"
}
func (this *CcoserviceServicegroupAddmemberRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CcoserviceServicegroupAddmemberRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CcoserviceServicegroupAddmemberRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CcoserviceServicegroupAddmemberRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CcoserviceServicegroupAddmemberRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CcoserviceServicegroupAddmemberRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_group_id"] = this.OpenGroupId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CcoserviceServicegroupAddmemberRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenGroupId, "openGroupId"); err != nil {
		return err
	}
	return nil
}
func (this *CcoserviceServicegroupAddmemberRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CcoserviceServicegroupAddmemberRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CcoserviceServicegroupAddmemberResponse struct {
	taobao.TaobaoResponse
	Result AddMembersResponseModel `json:"result,omitempty"`
}
type AddMembersResponseModel struct {
	DingtalkId string `json:"dingtalk_id,omitempty"`
}
