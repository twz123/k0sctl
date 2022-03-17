package cluster

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestPermStringUnmarshalWithOctal(t *testing.T) {
	u := UploadFile{}
	tmp := t.TempDir()
	ioutil.WriteFile(path.Join(tmp, "foo"), []byte("bar"), 0666)

	makeYaml := func(perm string) []byte {
		return []byte(fmt.Sprintf(`
src: %s
dstDir: %s
perm: %s
`, tmp, tmp, perm))
	}

	for _, test := range []struct{ name, perm string }{
		{"withOctal", "0755"},
		{"withString", `"0755"`},
	} {
		t.Run(test.name, func(t *testing.T) {
			yml := makeYaml(test.perm)
			require.NoError(t, yaml.Unmarshal(yml, &u))
			require.Equal(t, "0755", u.PermString)
		})
	}

	for _, test := range []struct{ name, perm string }{
		{"withInvalidString", "u+rwx"},
		{"withInvalidNumber", "0800"},
		{"withZero", "0"},
	} {
		t.Run(test.name, func(t *testing.T) {
			yml := makeYaml(test.perm)
			require.Error(t, yaml.Unmarshal(yml, &u))
		})
	}
}
