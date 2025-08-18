package im

type ResponseData struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	GroupId      string `json:"GroupId,omitempty"`
	Type         string `json:"Type,omitempty"`
	TotalCount   int    `json:"TotalCount,omitempty"`
	Next         uint64 `json:"Next,omitempty"`
}

type ChatGroupBody struct {
	OwnerAccount string `json:"Owner_Account,omitempty"`
	Type         string `json:"Type,omitempty"`
	GroupId      string `json:"GroupId,omitempty"`
	Name         string `json:"Name,omitempty"`
}
type DestroyGroupReq struct {
	GroupId string `json:"GroupId"`
}
type GetGroupsRequest struct {
	Limit     int    `json:"Limit,omitempty"`
	Next      uint64 `json:"Next,omitempty"`
	GroupType string `json:"GroupType,omitempty"`
}

type GroupSystemNotificationReq struct {
	GroupId          string   `json:"GroupId"`
	Content          string   `json:"Content"`
	ToMembersAccount []string `json:"ToMembers_Account,omitempty"`
}

type MsgElem struct {
	MsgType    string         `json:"MsgType"`
	MsgContent map[string]any `json:"MsgContent"`
}

type SendGroupMsgReq struct {
	GroupId                 string    `json:"GroupId"`
	Random                  uint32    `json:"Random"`
	MsgBody                 []MsgElem `json:"MsgBody"`
	FromAccount             string    `json:"From_Account,omitempty"`
	MsgPriority             string    `json:"MsgPriority,omitempty"`
	OnlineOnlyFlag          int       `json:"OnlineOnlyFlag,omitempty"`
	SendMsgControl          []string  `json:"SendMsgControl,omitempty"`
	SupportMessageExtension int       `json:"SupportMessageExtension,omitempty"`
}

type SendGroupMsgResp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	MsgTime      int64  `json:"MsgTime"`
	MsgSeq       int64  `json:"MsgSeq"`
}

type SendC2CMsgReq struct {
	SyncOtherMachine int       `json:"SyncOtherMachine,omitempty"`
	FromAccount      string    `json:"From_Account,omitempty"`
	ToAccount        string    `json:"To_Account"`
	OnlineOnlyFlag   int       `json:"OnlineOnlyFlag,omitempty"`
	MsgSeq           uint32    `json:"MsgSeq,omitempty"`
	MsgRandom        uint32    `json:"MsgRandom"`
	MsgBody          []MsgElem `json:"MsgBody"`
}

type SendC2CMsgResp struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorCode    int    `json:"ErrorCode"`
	ErrorInfo    string `json:"ErrorInfo"`
	MsgTime      int64  `json:"MsgTime"`
	MsgKey       string `json:"MsgKey"`
}

type GroupMsgGetSimpleReq struct {
	GroupId         string `json:"GroupId"`
	ReqMsgNumber    int    `json:"ReqMsgNumber"`
	ReqMsgSeq       uint64 `json:"ReqMsgSeq,omitempty"`
	WithRecalledMsg int    `json:"WithRecalledMsg,omitempty"`
}

type GroupMsgRecallReq struct {
	GroupId    string `json:"GroupId"`
	MsgSeqList []struct {
		MsgSeq uint64 `json:"MsgSeq"`
	} `json:"MsgSeqList"`
	Reason string `json:"Reason,omitempty"`
}

type GetGroupInfoReq struct {
	GroupIdList    []string `json:"GroupIdList"`
	ResponseFilter *struct {
		GroupBaseInfoFilter []string `json:"GroupBaseInfoFilter,omitempty"`
		MemberInfoFilter    []string `json:"MemberInfoFilter,omitempty"`
	} `json:"ResponseFilter,omitempty"`
}

type GetGroupMemberInfoReq struct {
	GroupId          string   `json:"GroupId"`
	Limit            int      `json:"Limit,omitempty"`
	Offset           int      `json:"Offset,omitempty"`
	MemberInfoFilter []string `json:"MemberInfoFilter,omitempty"`
}

type addMember struct {
	Member_Account string `json:"Member_Account"`
}

type AddGroupMemberReq struct {
	GroupId    string      `json:"GroupId"`
	Silence    int         `json:"Silence,omitempty"`
	MemberList []addMember `json:"MemberList"`
}

type DeleteGroupMemberReq struct {
	GroupId             string   `json:"GroupId"`
	Silence             int      `json:"Silence,omitempty"`
	Reason              string   `json:"Reason,omitempty"`
	MemberToDel_Account []string `json:"MemberToDel_Account"`
}

type ForbidSendMsgReq struct {
	GroupId         string   `json:"GroupId"`
	Members_Account []string `json:"Members_Account"`
	MuteTime        uint32   `json:"MuteTime"`
}

type AccountImportReq struct {
	UserID  string `json:"UserID"`
	Nick    string `json:"Nick,omitempty"`
	FaceUrl string `json:"FaceUrl,omitempty"`
}

type AccountCheckReq struct {
	CheckItem []struct {
		UserID string `json:"UserID"`
	} `json:"CheckItem"`
}

type KickReq struct {
	UserID string `json:"UserID"`
}
type QueryOnlineStatusReq struct {
	ToAccount            []string `json:"To_Account"`
	IsNeedDetail         int      `json:"IsNeedDetail,omitempty"`
	IsReturnInstid       int      `json:"IsReturnInstid,omitempty"`
	IsReturnCustomStatus int      `json:"IsReturnCustomStatus,omitempty"`
}

// ========== Accounts ==========
type MultiAccountImportReq struct {
	Accounts []AccountImportReq `json:"Accounts"`
}

type AccountDeleteReq struct {
	DeleteItem []struct {
		UserID string `json:"UserID"`
	} `json:"DeleteItem"`
}

// ========== C2C ==========
type BatchSendC2CMsgReq struct {
	SyncOtherMachine int       `json:"SyncOtherMachine,omitempty"`
	FromAccount      string    `json:"From_Account,omitempty"`
	ToAccount        []string  `json:"To_Account"`
	MsgRandom        uint32    `json:"MsgRandom"`
	MsgBody          []MsgElem `json:"MsgBody"`
}

type AdminSetMsgReadReq struct {
	ReportAccount string `json:"Report_Account"`
	PeerAccount   string `json:"Peer_Account"`
	MsgReadTime   uint64 `json:"MsgReadTime,omitempty"`
}

type AdminMsgWithdrawReq struct {
	FromAccount string `json:"From_Account"`
	ToAccount   string `json:"To_Account"`
	MsgKey      string `json:"MsgKey"`
}

type GetC2CUnreadMsgNumReq struct {
	ToAccount   string   `json:"To_Account"`
	PeerAccount []string `json:"Peer_Account,omitempty"`
}

// ========== Profile ==========
type PortraitSetReq struct {
	FromAccount string `json:"From_Account"`
	ProfileItem []struct {
		Tag   string      `json:"Tag"`
		Value interface{} `json:"Value"`
	} `json:"ProfileItem"`
}

type PortraitGetReq struct {
	ToAccount            []string `json:"To_Account"`
	TagList              []string `json:"TagList,omitempty"`
	LastStandardSequence int      `json:"LastStandardSequence,omitempty"`
	LastCustomSequence   int      `json:"LastCustomSequence,omitempty"`
}

// ========== Global NoSpeaking ==========
type SetNoSpeakingReq struct {
	Set_Account            string `json:"Set_Account"`
	C2CmsgNospeakingTime   uint32 `json:"C2CmsgNospeakingTime,omitempty"`
	GroupmsgNospeakingTime uint32 `json:"GroupmsgNospeakingTime,omitempty"`
}

type GetNoSpeakingReq struct {
	Get_Account string `json:"Get_Account"`
}

// ========== Ops ==========
type GetAppInfoReq struct {
	RequestField []string `json:"RequestField,omitempty"`
}

type GetIPListReq struct{}
