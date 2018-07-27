package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	//	"io"
	"net/url"
	"strings"
)

func zlib_decode(input []byte) (r []byte, err error) {
	reader, err := zlib.NewReader(bytes.NewBuffer(input))
	if err != nil {
		return
	}

	var n int64
	var out bytes.Buffer
	if n, err = out.ReadFrom(reader); err != nil {
		return
	}

	r = out.Bytes()[:n]
	err = reader.Close()
	return
}

func main() {
	str := "Q0qCeYRq+fMVbmoxUjyRDcMSEGLsed4jcdf4pqJjGiJ0E3uCcMxayqmZQrpiTAzAK1tapbsWwp6Gz+AeInAGcF8xSlCzsZliRJX1u6JiHGr9vQHTv7JvaweguB5KLbE2RJLnnLd35zRNcpuBgtKoSWjClTiZ8mBtzz8J3UioOhh4nGTNMQoCMRCF4bu8OsVMstmZTO8NvIDoFIHIaowgiHfXdMK2//fgvfHwO4wCzt1Pw4/16jDOtKYUi2iAv261/3dZefaxHXaS85RnvcCiRiIqixAXEQnoW5u735G3BgOnrKxpiar4fAMAAP//LP0kbA=="
	DECODE_LIMIT := 20
	for i := 0; i < DECODE_LIMIT; i++ {
		if strings.Contains(str, "%") {
			str, _ = url.QueryUnescape(str)
		} else {
			break
		}
	}
	if strings.Contains(str, " ") {
		str = strings.Replace(str, " ", "+", -1)
	}
	str += "==="[:(4-len(str)%4)%4]

	if decode, err := base64.StdEncoding.DecodeString(str); err != nil {
		fmt.Println("shit0", err)
		return
	} else {

		input := []byte(decode)[128:]
		r, err := zlib_decode(input)
		fmt.Println(string(r), err)

	}
}
