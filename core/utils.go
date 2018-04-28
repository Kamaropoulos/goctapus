package goctapus

import (
	"net"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func isValidPort(str string) (bool, string) {
	if i, err := strconv.Atoi(str); err == nil {
		if (1 <= i) && (i <= 65535) {
			return true, ""
		} else {
			return false, "A port number should be between 1 and 65535"
		}
	}
	return false, "Not a valid port number"
}

func isUsablePort(str string) (bool, string) {
	if res, err := isValidPort(str); !res {
		return res, err
	} else {
		ln, err1 := net.Listen("tcp", ":"+str)

		if err1 != nil {
			return false, err1.Error()
		}

		_ = ln.Close()

		return true, ""
	}
}
