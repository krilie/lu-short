package broker

import "time"

// 用户注册
const MemberRegistered = "member_registered"

type MemberRegisteredMsg struct {
	MemberId     string
	RegisterTime time.Time
}
