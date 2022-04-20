package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	genpath := "genfile"                    // path to the genfile directory
	pjPath := path.Join(genpath, "pj_test") // path to the pj_test directory
	os.RemoveAll(pjPath)
	if err := os.MkdirAll(pjPath, os.ModePerm); err != nil {
		panic(err)
	}

	group := flag.String("group", "", "agent group")
	setting := flag.String("set", "", "设置项，分别是 主叫开始项(忽略，仅为兼容pjtest而设),被叫开始项,数量，设置项之间用逗号隔开，多个设置用竖线隔开。例如：6101,6001,99|6300,6200,100")
	idIp := flag.String("idip", "", "sip id ip")
	registrar := flag.String("reg", "", "sip registrar")
	realm := flag.String("realm", "*", "sip realm")
	password := flag.String("password", "", "sip password，各个账号需统一密码")
	port := flag.Int("port", 10000, "起始的本地端口，默认10000,每新增一个配置会自动加1")

	flag.Parse()
	lst := strings.Split(*setting, "|")
	testPathIndex := 0

	for _, v := range lst {
		lst2 := strings.Split(v, ",")
		_, err := strconv.ParseInt(lst2[0], 10, 64)
		if err != nil {
			panic(err)
		}

		callee, err := strconv.ParseInt(lst2[1], 10, 64)
		if err != nil {
			panic(err)
		}

		count, err := strconv.Atoi(lst2[2])
		if err != nil {
			panic(err)
		}

		content, testPath, configFile := "", "", ""
		for i := 0; i < count; i++ {
			if content != "" {
				content += "--next-account\n"
			}
			content += fmt.Sprintf("--id=sip:%s_%v@%s\n", *group, callee, *idIp)
			content += fmt.Sprintf("--registrar=sip:%s\n", *registrar)
			content += fmt.Sprintf("--realm=%s\n", *realm)
			content += fmt.Sprintf("--username=%s_%v\n", *group, callee)
			content += fmt.Sprintf("--password=%s\n\n", *password)

			callee++

			if (i+1)%4 == 0 || i == count-1 {
				testPath = path.Join(pjPath, "test"+strconv.Itoa(testPathIndex)) // path to the test directory
				if err := os.MkdirAll(testPath, os.ModePerm); err != nil {       // 创建测试目录
					panic(err)
				}
				configFile = path.Join(testPath, "config.cfg") // path to the config.cfg file

				if content != "" {
					content += "--auto-answer=200\n"
					content += "--auto-loop\n"
					content += "--local-port=" + strconv.Itoa(*port) + "\n"
					if err := ioutil.WriteFile(configFile, []byte(content), 0644); err != nil { // 写入配置文件
						panic(err)
					}
					content = ""
					*port++
				}
				testPathIndex++
			}
		}
	}
}
