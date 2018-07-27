package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  renamedc: 2
  d: [3, 4]
key: !!str |
    -----BEGIN RSA PRIVATE KEY-----
    MIICXgIBAAKBgQC3torTchV3Fg7CnUfr5rNSGNT97xIUpoJq6Rs1ttj6asWkvlII
    GcYwH/FJQtA+0/9Ul+7vrjq8tvyLo0ESwns3+xp0exmLFgOHTgc781g2Rpw4IRXU
    W5uwfe6saD35Bolq7cYe2MdWqswKoGDIqsycDsuuefjtkSHq7wGI3KJ8yQIDAQAB
    AoGBAKLVASd5Lt1mh56nmV/WE4rox6BrjjMPgvkPtDZf4uHSUWw75KmZbripF0xA
    FCQb82wfXoDChP1Pk3iRMtRV3K+5Rn2m++fPA5DMWQZMLdTJXFVZgWty2tqFlQ03
    WqArGqGUTbZrTruUFRGHnBOGP5Ep8+CwxclbVvaNUD9RphgBAkEA11Br/PdJCnS4
    BuvtVbXI+oT3epqb9kljw54gophyXlGAjg/d/DPKJBZ2Tk0Rf368i0G8D/9DDLv7
    Uu5Km8InYQJBANptdBKO7Mf9qHXZZn85kZ2HexSYCd+HeoZHA6lD+O8/psGEfnNq
    8B1J2E96grtipFib6pEC5NzAHVqt0vVeFmkCQQCbaxen76+vdnTnuiEEv0UlYBlr
    THLLokn7CAcylqvnXRlYciegRpng1r3q40KALzFVq8teGLmDVaWiRRNhvxHBAkEA
    htJwhtP5iygixzs9bOIX42gwrHF8BqEVG7TRqbTu/p4HTAOAGbW0KB/giI5SC1/D
    mX4DtyZWJXvfoo0QXauJoQJAHQhVXLLeMLGOF/QVfMohEWEtixNHruag6/QKZQ+f
    ShIKJ0PM63cKQ22TeDnnDzG9h8bRaDmUDz0r+qlY19YhnA==
    -----END RSA PRIVATE KEY-----
fuck: what
`

type T struct {
	A string
	B struct {
		//RenamedC int `yaml:"RennamedC"`
		RenamedC int
		D        []int `yaml:",flow"`
	}
	Key   string
	Ddata string
}

func getTicketPrivateKey(privatekey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privatekey))
	if block == nil {
		return nil, errors.New("tticket private key decode fail")
	}

	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return rsa, nil
	default:
		return nil, errors.New("wrong block type ")
	}
}

func loadKey(path string) ([]byte, error) {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error while read file :" + path)
		return nil, err
	}
	return key, err
}

var data1 = `
woater_model_v1:
    area_ids:
       -  1152921590595125259  #无锡九里河公园商圈
       -  1152921506063122454  #泰州长泰路商圈
       -  1152921524690026511  #南京万寿商圈
    need_filter: true
woater_model_v2:
    area_ids:
       - 1152921590595125259  #无锡九里河公园商圈
       - 1152921506063122454  #泰州长泰路商圈
       - 1152921524690026511  #南京万寿商圈
    need_filter: true
    city_params:
       - [11, 35, 80, 500] #成都
       - [17, 35, 80, 500] #成都
       - [95, 35, 80, 500] #泰州
       - [47, 35, 80, 500] #无锡
       - [9, 35, 80, 500] #郑州
       - [0, 35, 80, 500] #兜底
    black_time: 6
`

type DynamicWoaterConfigV2 struct {
	AreaIds    []string   `yaml:"area_ids"`
	NeedFilter bool       `yaml:"need_filter"`
	CityParams [][]string `yaml:"city_params"`
	BlackTime  int        `yaml:"black_time"`
}

type DynamicWoaterConfigV1 struct {
	AreaIds    []string `yaml:"area_ids"`
	NeedFilter bool     `yaml:"need_filter"`
}

type DynamicConfig struct {
	WoaterModelV1 DynamicWoaterConfigV1 `yaml:"woater_model_v1"`
	WoaterModelV2 DynamicWoaterConfigV2 `yaml:"woater_model_v2"`
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
	fmt.Printf("data +++++++%v+++++++++++\n", t.Key)
	b, _ := getTicketPrivateKey([]byte(t.Key))
	fmt.Printf("decode +++++++%v+++++++++++\n", *b)

	fmt.Println("------------------------------------------------------------------\n")
	f, err := loadKey("./pem")
	fmt.Printf("file +++++++%v++++++++%v+++\n", string(f), time.Now())
	b1, _ := getTicketPrivateKey([]byte(f))
	fmt.Printf("decode +++++++%v+++++++%v++++\n", *b1, time.Now())

	fmt.Printf("================================\n")
	fmt.Printf("=============%v===================\n", t.B.RenamedC)
	//if *b != *b1 {
	//	fmt.Printf("fuck you\n")
	//}

	//d, err := yaml.Marshal(&t)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	//m := make(map[interface{}]interface{})

	//err = yaml.Unmarshal([]byte(data), &m)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- m:\n%v\n\n", m)

	//d, err = yaml.Marshal(&m) //if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- m dump:\n%s\n\n", string(d))

	t1 := DynamicConfig{}
	err = yaml.Unmarshal([]byte(data1), &t1)
	if err != nil {
		log.Fatalf("data1 error: %v", err)
	}
	fmt.Println("data1", t1)

}
