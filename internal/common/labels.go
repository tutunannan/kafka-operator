package common

import (
	"strings"
)

const (
	LabelCrName    = "app.kubernetes.io/Name"
	LabelComponent = "app.kubernetes.io/component"
	LabelManagedBy = "app.kubernetes.io/managed-by"
)

type RoleLabels struct {
	InstanceName string
	Name         string
}

func (r *RoleLabels) GetLabels() map[string]string {
	res := map[string]string{
		LabelCrName:    strings.ToLower(r.InstanceName),
		LabelManagedBy: "kafka-operator",
	}
	if r.Name != "" {
		res[LabelComponent] = r.Name
	}
	return res
}
