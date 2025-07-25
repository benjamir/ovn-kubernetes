// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package sbdb

import "github.com/ovn-kubernetes/libovsdb/model"

const LogicalDPGroupTable = "Logical_DP_Group"

// LogicalDPGroup defines an object in Logical_DP_Group table
type LogicalDPGroup struct {
	UUID      string   `ovsdb:"_uuid"`
	Datapaths []string `ovsdb:"datapaths"`
}

func (a *LogicalDPGroup) GetUUID() string {
	return a.UUID
}

func (a *LogicalDPGroup) GetDatapaths() []string {
	return a.Datapaths
}

func copyLogicalDPGroupDatapaths(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalDPGroupDatapaths(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *LogicalDPGroup) DeepCopyInto(b *LogicalDPGroup) {
	*b = *a
	b.Datapaths = copyLogicalDPGroupDatapaths(a.Datapaths)
}

func (a *LogicalDPGroup) DeepCopy() *LogicalDPGroup {
	b := new(LogicalDPGroup)
	a.DeepCopyInto(b)
	return b
}

func (a *LogicalDPGroup) CloneModelInto(b model.Model) {
	c := b.(*LogicalDPGroup)
	a.DeepCopyInto(c)
}

func (a *LogicalDPGroup) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *LogicalDPGroup) Equals(b *LogicalDPGroup) bool {
	return a.UUID == b.UUID &&
		equalLogicalDPGroupDatapaths(a.Datapaths, b.Datapaths)
}

func (a *LogicalDPGroup) EqualsModel(b model.Model) bool {
	c := b.(*LogicalDPGroup)
	return a.Equals(c)
}

var _ model.CloneableModel = &LogicalDPGroup{}
var _ model.ComparableModel = &LogicalDPGroup{}
