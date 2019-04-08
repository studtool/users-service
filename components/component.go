package components

import (
	"github.com/studtool/common/logs"
	"github.com/studtool/common/utils"

	"github.com/studtool/users-service/config"
)

type Component struct {
	service   string
	component string
}

func NewComponent(name string) *Component {
	return &Component{
		service:   config.ServiceName,
		component: name,
	}
}

func (c *Component) LogFieldsFor(f interface{}) *logs.LogFields {
	return &logs.LogFields{
		Service:   c.service,
		Component: c.component,
		Function:  utils.NameOf(f),
	}
}
