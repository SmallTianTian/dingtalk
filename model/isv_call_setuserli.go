package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewIsvCallSetuserlistRequest() *IsvCallSetuserlistRequest {
	return &IsvCallSetuserlistRequest{}
}

type IsvCallSetuserlistRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            IsvCallSetuserlistResponse
	StaffIdList     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *IsvCallSetuserlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *IsvCallSetuserlistRequest) SetStaffIdList(staffIdList2 string) {
	this.StaffIdList = staffIdList2
}
func (this *IsvCallSetuserlistRequest) GetStaffIdList() string {
	return this.StaffIdList
}
func (this *IsvCallSetuserlistRequest) GetApiMethodName() string {
	return "dingtalk.isv.call.setuserlist"
}
func (this *IsvCallSetuserlistRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *IsvCallSetuserlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *IsvCallSetuserlistRequest) GetTopApiCallType() string {
	return "top"
}
func (this *IsvCallSetuserlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *IsvCallSetuserlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *IsvCallSetuserlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["staff_id_list"] = this.StaffIdList
	return txtParams
}
func (this *IsvCallSetuserlistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StaffIdList, "staffIdList"); err != nil {
		return err
	}
	return nil
}
func (this *IsvCallSetuserlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *IsvCallSetuserlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IsvCallSetuserlistResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
