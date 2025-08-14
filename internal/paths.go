package im

// Auto-generated on 2025-08-14 from Tencent RTC Chat RESTful API List.
// Source: https://trtc.io/document/34621?product=chat&menulabel=core%20sdk&platform=unity%EF%BC%88game%20solution%EF%BC%89

// ========== Account Management ==========
const (
	AccountImportAPIPath        = "/v4/im_open_login_svc/account_import"
	MultiAccountImportAPIPath   = "/v4/im_open_login_svc/multiaccount_import"
	AccountDeleteAPIPath        = "/v4/im_open_login_svc/account_delete"
	AccountCheckAPIPath         = "/v4/im_open_login_svc/account_check"
	KickAPIPath                 = "/v4/im_open_login_svc/kick"
	QueryOnlineStatusAPIPath    = "/v4/openim/query_online_status"
)

// ========== One-to-One Message ==========
const (
	SendC2CMsgAPIPath           = "/v4/openim/sendmsg"
	BatchSendC2CMsgAPIPath      = "/v4/openim/batchsendmsg"
	ImportC2CMsgAPIPath         = "/v4/openim/importmsg"
	AdminGetRoamMsgAPIPath      = "/v4/openim/admin_getroammsg"
	AdminMsgWithdrawAPIPath     = "/v4/openim/admin_msgwithdraw"
	AdminSetMsgReadAPIPath      = "/v4/openim/admin_set_msg_read"
	GetC2CUnreadMsgNumAPIPath   = "/v4/openim/get_c2c_unread_msg_num"
	ModifyC2CMsgAPIPath         = "/v4/openim/modify_c2c_msg"
)

// ========== Pushing to All Users ==========
const (
	AllMemberPushAPIPath        = "/v4/all_member_push/im_push"
	SetAppAttrNameAPIPath       = "/v4/all_member_push/im_set_attr_name"
	GetAppAttrNameAPIPath       = "/v4/all_member_push/im_get_attr_name"
	GetUserAttrAPIPath          = "/v4/all_member_push/im_get_attr"
	SetUserAttrAPIPath          = "/v4/all_member_push/im_set_attr"
	RemoveUserAttrAPIPath       = "/v4/all_member_push/im_remove_attr"
	GetUserTagAPIPath           = "/v4/all_member_push/im_get_tag"
	AddUserTagAPIPath           = "/v4/all_member_push/im_add_tag"
	RemoveUserTagAPIPath        = "/v4/all_member_push/im_remove_tag"
	RemoveAllUserTagsAPIPath    = "/v4/all_member_push/im_remove_all_tags"
)

// ========== Profile Management ==========
const (
	PortraitSetAPIPath          = "/v4/profile/portrait_set"
	PortraitGetAPIPath          = "/v4/profile/portrait_get"
)

// ========== Relationship Chain Management ==========
const (
	FriendAddAPIPath            = "/v4/sns/friend_add"
	FriendImportAPIPath         = "/v4/sns/friend_import"
	FriendUpdateAPIPath         = "/v4/sns/friend_update"
	FriendDeleteAPIPath         = "/v4/sns/friend_delete"
	FriendDeleteAllAPIPath      = "/v4/sns/friend_delete_all"
	FriendCheckAPIPath          = "/v4/sns/friend_check"
	FriendGetAPIPath            = "/v4/sns/friend_get"
	FriendGetListAPIPath        = "/v4/sns/friend_get_list"
	BlackListAddAPIPath         = "/v4/sns/black_list_add"
	BlackListDeleteAPIPath      = "/v4/sns/black_list_delete"
	BlackListGetAPIPath         = "/v4/sns/black_list_get"
	BlackListCheckAPIPath       = "/v4/sns/black_list_check"
	SnsGroupAddAPIPath          = "/v4/sns/group_add"
	SnsGroupDeleteAPIPath       = "/v4/sns/group_delete"
	SnsGroupGetAPIPath          = "/v4/sns/group_get"
)

// ========== Following And Follower ==========
const (
	FollowAddAPIPath            = "/v4/follow/follow_add"
	FollowDeleteAPIPath         = "/v4/follow/follow_delete"
	FollowCheckAPIPath          = "/v4/follow/follow_check"
	FollowGetAPIPath            = "/v4/follow/follow_get"
	FollowGetInfoAPIPath        = "/v4/follow/follow_get_info"
)

// ========== Recent Contacts ==========
const (
	RecentContactGetListAPIPath   = "/v4/recentcontact/get_list"
	RecentContactDeleteAPIPath    = "/v4/recentcontact/delete"
	CreateContactGroupAPIPath     = "/v4/recentcontact/create_contact_group"
	DelContactGroupAPIPath        = "/v4/recentcontact/del_contact_group"
	UpdateContactGroupAPIPath     = "/v4/recentcontact/update_contact_group"
	SearchContactGroupAPIPath     = "/v4/recentcontact/search_contact_group"
	MarkContactAPIPath            = "/v4/recentcontact/mark_contact"
	GetContactGroupAPIPath        = "/v4/recentcontact/get_contact_group"
)

// ========== Group Management ==========
const (
	GetAllGroupsAPIPath           = "/v4/group_open_http_svc/get_appid_group_list"
	CreateGroupAPIPath            = "/v4/group_open_http_svc/create_group"
	GetGroupInfoAPIPath           = "/v4/group_open_http_svc/get_group_info"
	GetGroupMemberInfoAPIPath     = "/v4/group_open_http_svc/get_group_member_info"
	ModifyGroupBaseInfoAPIPath    = "/v4/group_open_http_svc/modify_group_base_info"
	AddGroupMemberAPIPath         = "/v4/group_open_http_svc/add_group_member"
	DeleteGroupMemberAPIPath      = "/v4/group_open_http_svc/delete_group_member"
	ModifyGroupMemberInfoAPIPath  = "/v4/group_open_http_svc/modify_group_member_info"
	DestroyGroupAPIPath           = "/v4/group_open_http_svc/destroy_group"
	GetJoinedGroupListAPIPath     = "/v4/group_open_http_svc/get_joined_group_list"
	GetRoleInGroupAPIPath         = "/v4/group_open_http_svc/get_role_in_group"
	ForbidSendMsgAPIPath          = "/v4/group_open_http_svc/forbid_send_msg"
	GetGroupShuttedUinAPIPath     = "/v4/group_open_http_svc/get_group_shutted_uin"
	SendGroupMsgAPIPath           = "/v4/group_open_http_svc/send_group_msg"
	SendGroupSystemNotifyPath     = "/v4/group_open_http_svc/send_group_system_notification"
	GroupMsgRecallAPIPath         = "/v4/group_open_http_svc/group_msg_recall"
	ChangeGroupOwnerAPIPath       = "/v4/group_open_http_svc/change_group_owner"
	ImportGroupAPIPath            = "/v4/group_open_http_svc/import_group"
	ImportGroupMsgAPIPath         = "/v4/group_open_http_svc/import_group_msg"
	ImportGroupMemberAPIPath      = "/v4/group_open_http_svc/import_group_member"
	SetUnreadMsgNumAPIPath        = "/v4/group_open_http_svc/set_unread_msg_num"
	DeleteGroupMsgBySenderAPIPath = "/v4/group_open_http_svc/delete_group_msg_by_sender"
	GroupMsgGetSimpleAPIPath      = "/v4/group_open_http_svc/group_msg_get_simple"
	GetOnlineMemberNumAPIPath     = "/v4/group_open_http_svc/get_online_member_num"
	GetGroupAttrAPIPath           = "/v4/group_open_attr_http_svc/get_group_attr"
	GetGroupBanMemberAPIPath      = "/v4/group_open_http_svc/get_group_ban_member"
	BanGroupMemberAPIPath         = "/v4/group_open_http_svc/ban_group_member"
	UnbanGroupMemberAPIPath       = "/v4/group_open_http_svc/unban_group_member"
	ModifyGroupAttrAPIPath        = "/v4/group_open_http_svc/modify_group_attr"
	ClearGroupAttrAPIPath         = "/v4/group_open_http_svc/clear_group_attr"
	SetGroupAttrAPIPath           = "/v4/group_open_http_svc/set_group_attr"
	DeleteGroupAttrAPIPath        = "/v4/group_open_http_svc/delete_group_attr"
	ModifyGroupMsgAPIPath         = "/v4/openim/modify_group_msg"
	SendBroadcastMsgAPIPath       = "/v4/group_open_http_svc/send_broadcast_msg"
	GetGroupCounterAPIPath        = "/v4/group_open_http_svc/get_group_counter"
	UpdateGroupCounterAPIPath     = "/v4/group_open_http_svc/update_group_counter"
	DeleteGroupCounterAPIPath     = "/v4/group_open_http_svc/delete_group_counter"
)

// ========== Global Mute Management ==========
const (
	SetNoSpeakingAPIPath          = "/v4/openconfigsvr/setnospeaking"
	GetNoSpeakingAPIPath          = "/v4/openconfigsvr/getnospeaking"
)

// ========== Operations Management ==========
const (
	GetAppInfoAPIPath             = "/v4/openconfigsvr/getappinfo"
	GetHistoryAPIPath             = "/v4/open_msg_svc/get_history"
	GetIPListAPIPath              = "/v4/ConfigSvc/GetIPList"
)
