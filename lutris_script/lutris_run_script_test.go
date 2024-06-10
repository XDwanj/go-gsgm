package lutris_script

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

func Test(t *testing.T) {
	mp := make(map[string]string)
	mp["PREFIX"] = "hello"
	mp["WINE"] = "world"
	system := &SystemDetail{
		Locale:   "utf8",
		PostExit: "/bin/bash",
		// Env:      mp,
	}

	yl, err := yaml.Marshal(system)
	if err != nil {
		return
	}

	fmt.Println(string(yl))
}
