package group

import (
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/utils"
)

type OnGroupListener interface {
	OnMemberEnter(groupId string, memberList string)
	OnMemberLeave(groupId string, member string)
	OnMemberInvited(groupId string, opUser string, memberList string)
	OnMemberKicked(groupId string, opUser string, memberList string)
	OnGroupCreated(groupId string)
	OnGroupInfoChanged(groupId string, groupInfo string)
	OnReceiveJoinApplication(groupId string, member string, opReason string)
	OnApplicationProcessed(groupId string, opUser string, AgreeOrReject int32, opReason string)
}

func (u *Group) SetGroupListener(callback OnGroupListener) {
	if callback == nil {
		return
	}
	u.listener = callback
}

func (u *Group) CreateGroup(callback common.Base, groupBaseInfo string, memberList string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "CreateGroup args: ", groupBaseInfo, memberList)
		var unmarshalCreateGroupBaseInfoParam sdk_params_callback.CreateGroupBaseInfoParam
		common.JsonUnmarshalAndArgsValidate(groupBaseInfo, &unmarshalCreateGroupBaseInfoParam, callback, operationID)
		var unmarshalCreateGroupMemberRoleParam sdk_params_callback.CreateGroupMemberRoleParam
		common.JsonUnmarshalAndArgsValidate(memberList, &unmarshalCreateGroupMemberRoleParam, callback, operationID)
		result := u.createGroup(callback, unmarshalCreateGroupBaseInfoParam, unmarshalCreateGroupMemberRoleParam, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, "CreateGroup callback: ", utils.StructToJsonString(result))
	}()
}

func (u *Group) JoinGroup(callback common.Base, groupID, reqMsg string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, reqMsg)
		u.joinGroup(groupID, reqMsg, callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.JoinGroupCallback))
		log.NewInfo(operationID, "JoinGroup callback: ", utils.StructToJsonString(sdk_params_callback.JoinGroupCallback))
	}()
}


func (u *Group) QuitGroup(callback common.Base, groupID string,  operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID)
		u.quitGroup(groupID,  callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.QuitGroupCallback))
		log.NewInfo(operationID, "QuitGroup callback: ", utils.StructToJsonString(sdk_params_callback.QuitGroupCallback))
	}()
}


func (u *Group) GetJoinedGroupList(callback common.Base, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ")
		groupList := u.getJoinedGroupList(callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupList)))
		log.NewInfo(operationID, "QuitGroup callback: ", utils.StructToJsonString(utils.StructToJsonString(groupList)))
	}()
}


func (u *Group) GetGroupsInfo(callback common.Base, groupIDList string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupIDList)
		var unmarshalGetGroupsInfoParam sdk_params_callback.GetGroupsInfoParam
		common.JsonUnmarshalAndArgsValidate(groupIDList, &unmarshalGetGroupsInfoParam, callback, operationID)
		groupsInfoList := u.getGroupsInfo(unmarshalGetGroupsInfoParam, callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupsInfoList)))
		log.NewInfo(operationID, "GetGroupsInfo callback: ", utils.StructToJsonString(utils.StructToJsonString(groupsInfoList)))

	}()
}


func (u *Group) SetGroupInfo( callback common.Base, groupInfo string, groupID string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupInfo, groupID)
		var unmarshalSetGroupInfoParam sdk_params_callback.SetGroupInfoParam
		common.JsonUnmarshalAndArgsValidate(groupInfo, &unmarshalSetGroupInfoParam, callback, operationID)
		u.setGroupInfo( callback, unmarshalSetGroupInfoParam, groupID, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(sdk_params_callback.SetGroupInfoCallback)))
		log.NewInfo(operationID, "SetGroupInfo callback: ", utils.StructToJsonString(sdk_params_callback.SetGroupInfoCallback))
	}()
}


func (u *Group) GetGroupMemberList(callback common.Base, groupID string, filter int32, next int32,  operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, filter, next)
		groupMemberList := u.getGroupMemberList( callback, groupID, filter, next, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupMemberList)))
		log.NewInfo(operationID, "GetGroupMemberList callback: ", utils.StructToJsonString(groupMemberList))
	}()
}

func (u *Group) GetGroupMembersInfo(callback common.Base, groupID string, userList string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, userList)
		var unmarshalParam sdk_params_callback.GetGroupMembersInfoParam
		groupMemberList := u.getGroupMembersInfo( callback, groupID, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupMemberList)))
		log.NewInfo(operationID, "GetGroupMembersInfo callback: ", utils.StructToJsonString(groupMemberList))
	}()
}

func (u *Group) KickGroupMember(callback common.Base, groupID string, reason string, userList string,  operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, reason, userList)
		var unmarshalParam sdk_params_callback.KickGroupMemberParam
		result := u.kickGroupMember(callback,  groupID,  unmarshalParam, reason, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(result)))
		log.NewInfo(operationID, "GetGroupMembersInfo callback: ", utils.StructToJsonString(result))
	}()
}


func (u *Group) TransferGroupOwner(callback common.Base, groupID, newOwnerUserID string,  operationID string) {
	if callback == nil {
		return
	}
	go func() {
		u.transferGroupOwner(callback, groupID, newOwnerUserID, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(sdk_params_callback.TransferGroupOwnerCallback)))
	}()
}

func (u *Group) InviteUserToGroup(callback common.Base, groupID, reason string, userList string,  operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, reason, userList)
		var unmarshalParam sdk_params_callback.InviteUserToGroupParam
		result := u.inviteUserToGroup(callback,  groupID, reason, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(result)))
		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(result))
	}()
}

func (u *Group) GetGroupApplicationList(callback common.Base, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ",)
		result := u.getGroupApplicationList(callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(result)))
		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(result))
	}()
}

func (u *Group) AcceptGroupApplication(callback common.Base, groupID, fromUserID,  handleMsg string, operationID string){
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, fromUserID,  handleMsg)
		u.processGroupApplication(callback, groupID, fromUserID,  handleMsg, 1, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.AcceptGroupApplicationCallback))
		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(sdk_params_callback.AcceptGroupApplicationCallback))
	}()
}

func (u *Group) RefuseGroupApplication(callback common.Base, groupID, fromUserID,  handleMsg string, operationID string){
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, fromUserID,  handleMsg)
		u.processGroupApplication(callback, groupID, fromUserID,  handleMsg, -1, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.RefuseGroupApplicationCallback))
		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(sdk_params_callback.RefuseGroupApplicationCallback))
	}()
}


/*
func (u *UserRelated) TsetGetGroupApplicationList(callback Base) string {
	if callback == nil {
		return ""
	}

	r, err := u.getGroupApplicationList()
	if err != nil {
		sdkLog("getGroupApplicationList faild, ", err.Error())
		callback.OnError(ErrCodeGroup, err.Error())
		return ""
	}
	jsonResult, err := json.Marshal(r)
	if err != nil {
		sdkLog("getGroupApplicationList faild, ", err.Error())
		callback.OnError(ErrCodeGroup, err.Error())
		return ""
	}
	callback.OnSuccess(string(jsonResult))
	return string(jsonResult)

}*/
