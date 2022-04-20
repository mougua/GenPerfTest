package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	genpath := "genfile"                        // path to the genfile directory
	sippPath := path.Join(genpath, "sipp_test") // path to the sipp_test directory
	if err := os.MkdirAll(sippPath, os.ModePerm); err != nil {
		panic(err)
	}

	group := flag.String("group", "", "agent group") //"10003"
	setting := flag.String("set", "", "设置项，分别是 主叫开始项,被叫开始项,数量，设置项之间用逗号隔开，多个设置用竖线隔开。例如：\"6101,6001,99|6300,6200,100\"")
	head := flag.String("head", "SEQUENTIAL", "文件头部设置，可以为SEQUENTIAL或RANDOM")
	flag.Parse()
	lst := strings.Split(*setting, "|")
	content := *head + "\n"
	for _, v := range lst {
		lst2 := strings.Split(v, ",")
		caller, err := strconv.ParseInt(lst2[0], 10, 64)
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

		for i := 0; i < count; i++ {
			content += *group + "_" + strconv.FormatInt(caller, 10) + ";" +
				*group + "_" + strconv.FormatInt(callee, 10) + "\n"
			caller++
			callee++
		}
	}

	if err := ioutil.WriteFile(path.Join(sippPath, "t.csv"), []byte(content), 0644); err != nil {
		panic(err)
	}
}
