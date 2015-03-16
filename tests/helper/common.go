package helper

import "gitlab.com/q-dev/q-client/common"

func FromHex(h string) []byte {
	if common.IsHex(h) {
		h = h[2:]
	}

	return common.Hex2Bytes(h)
}
