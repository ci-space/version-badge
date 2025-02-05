package version

import "strings"

type Version struct {
	Object  string
	Version string
}

func (v *Version) ShortObjectName() string {
	parts := strings.Split(v.Object, "/")

	return parts[len(parts)-1]
}
