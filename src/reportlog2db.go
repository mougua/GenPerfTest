package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func run() {
	f := flag.String("f", "", "file path")
	isLocal := flag.Bool("l", false, "is local")
	flag.Parse()
	if *f == "" {
		panic("file is null")
	}
	file, err := os.OpenFile(*f, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	bufOut := ""

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		bufOut += handleLine(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
	if *isLocal {
		bufOut += `
		update ivr_history_details
		set tenant_id = '10000',
			ivr_id='8'
		where tenant_id = '10003';

		update ivr_histories
		set tenant_id = '10000',
			ivr_id='8'
		where tenant_id = '10003';
		`
	}

	fileOut := *f + ".sql"
	if err := ioutil.WriteFile(fileOut, []byte(bufOut), 0644); err != nil {
		panic(err)
	}
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

func main() {
	// s := handleLine(`05-31 22:16:27,572 INFO  [tevent.HandleIVRRoutedEvent] classHandleIVRRoutedEvent,onEvent(),ivrHistory:IvrHistory{id=null, callId='b8c4100a-bd94-41fa-9a35-93e2968f6f52', startTime=Tue May 31 22:16:27 CST 2022, endTime=Tue May 31 22:16:27 CST 2022, ani='1913682215012', dnis='95598', callType=2, tenantId=10003, exitStatus=2, ivrNode=17717, previousIvrNode=-1, thisQueue='null', campaignId=null, campaignContactId=null, curIvrNode=17717, ivrId=107, previousIvrId=-1, routeType='SUCCESS', ivrDtmfDigit='null', ivrDigitType='httpRequest', target='{"request_headers":{"Content-Type":"application/json;charset=UTF-8"},"trans_url":"http://10.147.22.26:8010/GDYKZX/gbc-mssc-restful-service/saveIVRConversation","request_body":"{\"lhhm\":\"${_qqzsjh}\",\"gddwbm\":\"${_gsdbm}\",\"dqbm\":\"${_dqbm}\",\"dhnr\":\"${_zndhnr}\"\n,\"callid\":\"$${cti_call_id}\"}","trans_request_body":"{\"lhhm\":\"13682215012\",\"gddwbm\":\"0319\",\"dqbm\":\"031900\",\"dhnr\":\"#B:2022-05-31 22:16:27尊敬的客户您好！我是智能客服小赫兹，听您说话就能办理业务，请说出您的业务需求，如查电费、查停电等，您请说\"\n,\"callid\":\"b8c4100a-bd94-41fa-9a35-93e2968f6f52\"}","req_timeout":"5","_custom_result_key":"","url":"${_BCDHNR}/saveIVRConversation","result":{"request":{"body":"{\"lhhm\":\"13682215012\",\"gddwbm\":\"0319\",\"dqbm\":\"031900\",\"dhnr\":\"#B:2022-05-31 22:16:27尊敬的客户您好！我是智能客服小赫兹，听您说话就能办理业务，请说出您的业务需求，如查电费、查停电等，您请说\"\n,\"callid\":\"b8c4100a-bd94-41fa-9a35-93e2968f6f52\"}","url":"http://10.147.22.26:8010/GDYKZX/gbc-mssc-restful-service/saveIVRConversation","method":"POST","headers":{"Content-Length":291,"Content-Type":"application/json;charset=UTF-8"}},"response":{"body":{"data":"保存信息成功","msg":"请求成功！","time":1654006587559,"code":200},"res":1,"headers":{"access-control-allow-methods":"GET, POST, PUT, DELETE, PATCH, OPTIONS","date":"Tue, 31 May 2022 14:16:27 GMT","transfer-encoding":"chunked","access-control-allow-origin":"*","x-mg-code":"200","content-type":"application/json;charset=UTF-8","x-mg-timestamp":"1654006587560","access-control-allow-headers":"DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token","x-mg-node":"yx-wfw-v23-845f954547-g5pxn","x-mg-traceid":"waY85nhhTKOhpZY0176AgA","x-mg-span":"111","x-request-id":"waY85nhhTKOhpZY0176AgA"},"code":200}},"http_method":"POST"}', ivrStatus='end_ivr_component', targetType='', executeContext='inbound_ivr', channelStatus='null', eventSequence=422119, fwm='null', result='', extension='null', typeText='null', typeCode='null', ivrSessionId='f0c7e30c-998c-42de-b286-8250bad63d08'} | Report10003-thread-1`)
	// fmt.Println(s)
	run()
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
func handleLine(line string) string {
	columnList := []string{"call_id", "start_time", "end_time", "ani", "dnis", "call_type", "tenant_id", "exit_status", "ivr_node", "previous_ivr_node", "this_queue", "campaign_id", "campaign_contact_id", "cur_ivr_node", "ivr_id", "previous_ivr_id", "route_type", "ivr_dtmf_digit", "ivr_digit_type", "target", "ivr_status", "target_type", "execute_context", "channel_status", "event_sequence", "fwm_code", "result", "type_text", "type_code", "ivr_session_id"}
	if !strings.Contains(line, "classHandleIVRRoutedEvent,onEvent(),ivrHistory:") {
		return ""
	}
	// fmt.Println(line)

	insertSql1 := "insert into ivr_histories ("
	insertSql2 := "insert into ivr_history_details ("

	isStart := strings.Contains(line, "endTime=null")

	insertSqlValues := "values ("
	s := strings.TrimSpace(line[strings.Index(line, "id="):strings.Index(line, "} |")])
	arr := strings.Split(s, ",")
	var startTimeBuf string
	for i, v := range arr {
		tmp := strings.Split(v, "=")

		column := strings.TrimSpace(snakeString(tmp[0]))
		if !Find(columnList, column) {
			continue
		}

		if column == "fwm" {
			column = "fwm_code"
		}
		var value string
		var idx int
		// 处理复杂value的字段

		if column == "target" {
			value = strings.TrimSpace(line[strings.Index(line, "target=")+7:])
			idx = strings.Index(value, "}', ")
			if idx < 0 {
				value = "''"
			} else {
				value = strings.TrimSpace(value[:idx+2])
			}
		} else if column == "ivr_dtmf_digit" {
			value = strings.TrimSpace(tmp[1])
			if value == "'null'" || value == "'" {
				value = "null"
			} else if value[0] == '\'' && value[len(value)-1] != '\'' {
				value += "'"
			}
		} else {
			// println(column)
			value = strings.TrimSpace(tmp[1])
			if value == "'null'" {
				value = "null"
			} else if !strings.Contains(column, "time") || value == "null" {
				value = tmp[1]
			} else {
				value = "'" + value + "'"
			}
		}
		if column == "start_time" {
			startTimeBuf = value
		}
		if column == "end_time" && isStart {
			value = startTimeBuf
		}

		if isStart {
			insertSql1 += column
		}
		insertSql2 += column
		insertSqlValues += value

		if i != len(arr)-1 {
			if isStart {
				insertSql1 += ","
			}
			insertSql2 += ","
			insertSqlValues += ","
		} else {
			if isStart {
				insertSql1 += ")"
			}

			insertSql2 += ")"
			insertSqlValues += ");\n"
		}
	}
	if isStart {
		return insertSql1 + insertSqlValues + insertSql2 + insertSqlValues
	} else {
		return insertSql2 + insertSqlValues
	}
}
