package ifaces

import (
	"gorm.io/gorm/schema"
)

type Modeler interface {
	schema.Tabler
	EntityName() string
}
