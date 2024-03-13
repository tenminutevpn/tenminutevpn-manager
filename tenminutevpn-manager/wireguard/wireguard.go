package wireguard

import (
	"fmt"
	"os/exec"
	"strings"
)

func GenKey() string {
	cmd := exec.Command("wg", "genkey")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	key := strings.TrimSpace(string(out))
	return key
}
