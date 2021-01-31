package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewIsvCallRemoveuserlistRequest() *IsvCallRemoveuserlistRequest {
	return &IsvCallRemoveuserlistRequest{}
}

type IsvCallRemoveuserlistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            IsvCallRemoveuserlistResponse
	StaffIdList     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *IsvCallRemoveuserlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *IsvCallRemoveuserlistRequest) SetStaffIdList(staffIdList2 string) {
	this.StaffIdList = staffIdList2
}
func (this *IsvCallRemoveuserlistRequest) GetStaffIdList() string {
	return this.StaffIdList
}
func (this *IsvCallRemoveuserlistRequest) GetApiMethodName() string {
	return "dingtalk.isv.call.removeuserlist"
}
func (this *IsvCallRemoveuserlistRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *IsvCallRemoveuserlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *IsvCallRemoveuserlistRequest) GetTopApiCallType() string {
	return "top"
}
func (this *IsvCallRemoveuserlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *IsvCallRemoveuserlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *IsvCallRemoveuserlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["staff_id_list"] = this.StaffIdList
	return txtParams
}
func (this *IsvCallRemoveuserlistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StaffIdList, "staffIdList"); err != nil {
		return err
	}
	return nil
}
func (this *IsvCallRemoveuserlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *IsvCallRemoveuserlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IsvCallRemoveuserlistResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
