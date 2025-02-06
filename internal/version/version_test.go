package version

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_ShortObjectName(t *testing.T) {
	cases := []struct {
		FullName string

		Expected string
	}{
		{
			FullName: "",
			Expected: "",
		},
		{
			FullName: "dependency",
			Expected: "dependency",
		},
		{
			FullName: "vendor/dependency",
			Expected: "dependency",
		},
	}

	for _, tCase := range cases {
		t.Run(fmt.Sprintf("%s -> %s", tCase.FullName, tCase.Expected), func(t *testing.T) {
			vers := &Version{
				Object: tCase.FullName,
			}

			assert.Equal(t, tCase.Expected, vers.ShortObjectName())
		})
	}
}
