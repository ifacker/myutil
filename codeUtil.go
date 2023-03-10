package code

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"unsafe"
)

// unicode 转 string
func Unicode2String(form string) (result string, err error) {
	run := func(form string) (to []rune) {
		bs, err := hex.DecodeString(form[2:])
		if err != nil {
			//fmt.Println(err)
			return
		}
		for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
			binary.Read(br, binary.BigEndian, &r)
			to = append(to, rune(r))
		}
		return
	}

	formTmp := strings.Replace(form, `\U`, `\u`, -1)
	for {
		if formTmp == "" {
			break
		}

		if len(formTmp) >= 2 && formTmp[:2] == `\u` {
			index := strings.IndexAny(formTmp, `\u`)
			tmp := formTmp[index : index+6]
			strTmp := string(run(tmp))
			result += strTmp
			formTmp = formTmp[index+6:]
			//old := formTmp[]
			//fmt.Println(tmp)

		} else {
			strTmp := formTmp[:1]
			result += strTmp
			if len(formTmp) > 1 {
				formTmp = formTmp[1:]
			} else {
				formTmp = ""
			}
		}

	}
	return
}

// string 转 byte
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// byte 转 string
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}
