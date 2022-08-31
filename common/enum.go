package common

type MQTopic string

const (
	LoginTopic                MQTopic = "login"
	AssignRolePermissionTopic MQTopic = "assign-role-permission"
	RemoveRolePermissionTopic MQTopic = "remove-role-permission"
	DeleteRoleTopic           MQTopic = "delete-role"
	DeletePermissionTopic     MQTopic = "delete-permission"
)
