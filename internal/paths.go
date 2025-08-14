package im

const (
	CreateGroupAPIPath        = "/v4/group_open_http_svc/create_group"
	DestroyGroupAPIPath       = "/v4/group_open_http_svc/destroy_group"
	GetAllGroupsAPIPath       = "/v4/group_open_http_svc/get_appid_group_list"
	SendGroupSystemNotifyPath = "/v4/group_open_http_svc/send_group_system_notification"
	GetGroupInfoAPIPath       = "/v4/group_open_http_svc/get_group_info"
	GetGroupMemberInfoAPIPath = "/v4/group_open_http_svc/get_group_member_info"
	AddGroupMemberAPIPath     = "/v4/group_open_http_svc/add_group_member"
	DeleteGroupMemberAPIPath  = "/v4/group_open_http_svc/delete_group_member"
	ForbidSendMsgAPIPath      = "/v4/group_open_http_svc/forbid_send_msg"
)

const (
	SendGroupMsgAPIPath      = "/v4/group_open_http_svc/send_group_msg"
	GroupMsgGetSimpleAPIPath = "/v4/group_open_http_svc/group_msg_get_simple"
	GroupMsgRecallAPIPath    = "/v4/group_open_http_svc/group_msg_recall"
	SendC2CMsgAPIPath        = "/v4/openim/sendmsg"
)

const (
	AccountImportAPIPath     = "/v4/im_open_login_svc/account_import"
	AccountCheckAPIPath      = "/v4/im_open_login_svc/account_check"
	KickAPIPath              = "/v4/im_open_login_svc/kick"
	QueryOnlineStatusAPIPath = "/v4/openim/query_online_status"
)
