package broker

import (
	"github.com/cesc1802/auth-service/common"
)

type Message struct {
	Value MessageValue
	Topic common.MQTopic
}

type MessageValue struct {
	RoleIDs       []uint
	PermissionIDs []uint
	UserID        uint
}
