package qwhisper

import (
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/whisper"
)

type Message struct {
	ref     *whisper.Message
	Flags   int32  `json:"flags"`
	Payload string `json:"payload"`
	From    string `json:"from"`
}

func ToQMessage(msg *whisper.Message) *Message {
	return &Message{
		ref:     msg,
		Flags:   int32(msg.Flags),
		Payload: "0x" + ethutil.Bytes2Hex(msg.Payload),
		From:    "0x" + ethutil.Bytes2Hex(crypto.FromECDSAPub(msg.Recover())),
	}
}
