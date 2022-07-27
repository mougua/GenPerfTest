package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	genpath := "genfile"                   // path to the genfile directory
	fileDir := path.Join(genpath, "users") // path to the sipp_test directory
	if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
		panic(err)
	}

	prefix := flag.String("prefix", "", "agent prefix")
	user_from := flag.Int("from", 1000, "开始")
	user_to := flag.Int("to", 1001, "结束")
	flag.Parse()
	content := ""
	template := `<include>
		<user id="%s">
		<params>
			<param name="password" value="888888"/>
			<param name="vm-password" value="%s"/>
		</params>
		<variables>
			<variable name="toll_allow" value="domestic,international,local"/>
			<variable name="accountcode" value="%s"/>
			<variable name="user_context" value="default"/>
			<variable name="effective_caller_id_name" value="Extension %s"/>
			<variable name="effective_caller_id_number" value="%s"/>
			<variable name="outbound_caller_id_name" value="$${outbound_caller_name}"/>
			<variable name="outbound_caller_id_number" value="$${outbound_caller_id}"/>
			<variable name="callgroup" value="techsupport"/>
		</variables>
		</user>
	</include>

	`

	for i := *user_from; i <= *user_to; i++ {
		s := fmt.Sprintf("%s%d", *prefix, i)
		content += fmt.Sprintf(template, s, s, s, s, s)
		if err := ioutil.WriteFile(path.Join(fileDir, "test-users.xml"), []byte(content), 0644); err != nil {
			panic(err)
		}
	}

}
