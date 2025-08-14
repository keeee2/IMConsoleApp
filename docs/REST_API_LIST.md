# REST API List (Tencent RTC Chat)

> Source: https://trtc.io/document/34621?product=chat&menulabel=core%20sdk&platform=unity%EF%BC%88game%20solution%EF%BC%89

이 문서는 콘솔 앱이 사용하는 REST API 엔드포인트 목록을 카테고리별로 정리한 것입니다. 모든 경로 상수는 `internal/paths.go`에 정의되어 있습니다.

## Account Management
- `v4/im_open_login_svc/account_import` — 단일 계정 가져오기(등록)
- `v4/im_open_login_svc/multiaccount_import` — 다중 계정 가져오기(등록)
- `v4/im_open_login_svc/account_delete` — 계정 삭제
- `v4/im_open_login_svc/account_check` — 계정 상태 조회
- `v4/im_open_login_svc/kick` — 로그인 상태 무효화
- `v4/openim/query_online_status` — 온라인 상태 조회

## One-to-One Message
- `v4/openim/sendmsg` — 1:1 메시지 전송(단건)
- `v4/openim/batchsendmsg` — 1:1 메시지 전송(다중 대상)
- `v4/openim/importmsg` — 1:1 메시지 가져오기(히스토리)
- `v4/openim/admin_getroammsg` — 1:1 로밍 메시지 조회
- `v4/openim/admin_msgwithdraw` — 1:1 메시지 회수
- `v4/openim/admin_set_msg_read` — 1:1 읽음 처리
- `v4/openim/get_c2c_unread_msg_num` — 1:1 미읽음 개수 조회
- `v4/openim/modify_c2c_msg` — 1:1 과거 메시지 수정

## Pushing to All Users
- `v4/all_member_push/im_push` — 전체 푸시
- `v4/all_member_push/im_set_attr_name` — 앱 속성명 설정
- `v4/all_member_push/im_get_attr_name` — 앱 속성명 조회
- `v4/all_member_push/im_get_attr` — 유저 속성 조회
- `v4/all_member_push/im_set_attr` — 유저 속성 설정
- `v4/all_member_push/im_remove_attr` — 유저 속성 삭제
- `v4/all_member_push/im_get_tag` — 유저 태그 조회
- `v4/all_member_push/im_add_tag` — 유저 태그 추가
- `v4/all_member_push/im_remove_tag` — 유저 태그 삭제
- `v4/all_member_push/im_remove_all_tags` — 유저 전체 태그 삭제

## Profile Management
- `v4/profile/portrait_set` — 프로필 설정
- `v4/profile/portrait_get` — 프로필 조회

## Relationship Chain Management
- `v4/sns/friend_add` — 친구 추가
- `v4/sns/friend_import` — 친구 일괄 가져오기
- `v4/sns/friend_update` — 친구 정보 업데이트
- `v4/sns/friend_delete` — 친구 삭제
- `v4/sns/friend_delete_all` — 전체 친구 삭제
- `v4/sns/friend_check` — 친구 관계 확인
- `v4/sns/friend_get` — 친구 목록 조회
- `v4/sns/friend_get_list` — 지정 친구 조회
- `v4/sns/black_list_add` — 차단 추가
- `v4/sns/black_list_delete` — 차단 해제
- `v4/sns/black_list_get` — 차단 목록 조회
- `v4/sns/black_list_check` — 상호 차단 여부 확인
- `v4/sns/group_add` — 사용자 리스트 추가
- `v4/sns/group_delete` — 사용자 리스트 삭제
- `v4/sns/group_get` — 사용자 리스트 조회

## Following & Follower
- `v4/follow/follow_add` — 팔로우
- `v4/follow/follow_delete` — 언팔로우
- `v4/follow/follow_check` — 팔로우 관계 확인
- `v4/follow/follow_get` — 팔로잉/팔로워/맞팔 목록
- `v4/follow/follow_get_info` — 팔로잉/팔로워/맞팔 수

## Recent Contacts
- `v4/recentcontact/get_list` — 대화 목록 조회
- `v4/recentcontact/delete` — 대화 삭제
- `v4/recentcontact/create_contact_group` — 대화 그룹 데이터 생성
- `v4/recentcontact/del_contact_group` — 대화 그룹 데이터 삭제
- `v4/recentcontact/update_contact_group` — 대화 그룹 데이터 업데이트
- `v4/recentcontact/search_contact_group` — 대화 그룹 마크 검색
- `v4/recentcontact/mark_contact` — 대화 마크 생성/업데이트
- `v4/recentcontact/get_contact_group` — 대화 그룹 마크 조회

## Group Management
- `v4/group_open_http_svc/get_appid_group_list` — 앱 내 전체 그룹 조회
- `v4/group_open_http_svc/create_group` — 그룹 생성
- `v4/group_open_http_svc/get_group_info` — 그룹 프로필 조회
- `v4/group_open_http_svc/get_group_member_info` — 멤버 프로필 조회
- `v4/group_open_http_svc/modify_group_base_info` — 그룹 프로필 수정
- `v4/group_open_http_svc/add_group_member` — 멤버 추가
- `v4/group_open_http_svc/delete_group_member` — 멤버 삭제
- `v4/group_open_http_svc/modify_group_member_info` — 멤버 프로필 수정
- `v4/group_open_http_svc/destroy_group` — 그룹 해산
- `v4/group_open_http_svc/get_joined_group_list` — 유저 가입 그룹 조회
- `v4/group_open_http_svc/get_role_in_group` — 그룹 내 역할 조회
- `v4/group_open_http_svc/forbid_send_msg` — 멤버 음소거/해제
- `v4/group_open_http_svc/get_group_shutted_uin` — 음소거 멤버 목록
- `v4/group_open_http_svc/send_group_msg` — 그룹 일반 메시지 전송
- `v4/group_open_http_svc/send_group_system_notification` — 그룹 시스템 메시지 전송
- `v4/group_open_http_svc/group_msg_recall` — 그룹 메시지 회수
- `v4/group_open_http_svc/change_group_owner` — 그룹장 변경
- `v4/group_open_http_svc/import_group` — 그룹 프로필 가져오기
- `v4/group_open_http_svc/import_group_msg` — 그룹 메시지 가져오기
- `v4/group_open_http_svc/import_group_member` — 그룹 멤버 가져오기
- `v4/group_open_http_svc/set_unread_msg_num` — 멤버 미읽음 수 설정
- `v4/group_open_http_svc/delete_group_msg_by_sender` — 특정 발신자 메시지 삭제
- `v4/group_open_http_svc/group_msg_get_simple` — 그룹 메시지 히스토리 조회
- `v4/group_open_http_svc/get_online_member_num` — AV 그룹 온라인 수 조회
- `v4/group_open_attr_http_svc/get_group_attr` — 그룹 커스텀 속성 조회
- `v4/group_open_http_svc/get_group_ban_member` — 그룹 밴 멤버 목록
- `v4/group_open_http_svc/ban_group_member` — 그룹 밴 설정
- `v4/group_open_http_svc/unban_group_member` — 그룹 밴 해제
- `v4/group_open_http_svc/modify_group_attr` — 그룹 커스텀 속성 수정
- `v4/group_open_http_svc/clear_group_attr` — 그룹 커스텀 속성 초기화
- `v4/group_open_http_svc/set_group_attr` — 그룹 커스텀 속성 설정
- `v4/group_open_http_svc/delete_group_attr` — 그룹 커스텀 속성 삭제
- `v4/openim/modify_group_msg` — 그룹 과거 메시지 수정
- `v4/group_open_http_svc/send_broadcast_msg` — 모든 AV 그룹에 브로드캐스트
- `v4/group_open_http_svc/get_group_counter` — 그룹 카운터 조회
- `v4/group_open_http_svc/update_group_counter` — 그룹 카운터 업데이트
- `v4/group_open_http_svc/delete_group_counter` — 그룹 카운터 삭제

## Global Mute Management
- `v4/openconfigsvr/setnospeaking` — 전역 음소거 설정
- `v4/openconfigsvr/getnospeaking` — 전역 음소거 조회

## Operations Management
- `v4/openconfigsvr/getappinfo` — 운영 데이터 조회
- `v4/open_msg_svc/get_history` — 최근 메시지 다운로드
- `v4/ConfigSvc/GetIPList` — 서버 IP 조회

---

> 변경 내역: `internal/paths.go` 전체 경로 상수 보강, 본 문서 신설.
