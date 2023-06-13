package feishutalk

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	JobName := flag.String("job-name", "", "JOB_NAME")
	BuildDisplayName := flag.String("build-display_name", "", "BUILD_DISPLAY_NAME")
	Message := flag.String("message", "", "MESSAGE")
	Author := flag.String("author", "", "AUTHOR")

	flag.Parse()

	if *JobName == "" || *BuildDisplayName == "" || *Message == "" || *Author == "" {
		fmt.Println("Usage: go run main.go -job-name <JOB_NAME> -build-display_name <BUILD_DISPLAY_NAME> -message <MESSAGE> -author <AUTHOR>")
		return
	}

	url := "https://open.feishu.cn/open-apis/bot/v2/hook/602800d2-411a-4758-bffb"
	jsonStr := []byte(fmt.Sprintf(`{
    	"msg_type": "interactive",
    	"card": {
        	"elements": [{
            	"tag": "div",
            	"text": {
					"content": "📋 **任务名称**: %s \n🔢 **任务编号**: %s \n🌟 **构建状态**: <font color='red'>%s</font> \n👤 **执   行 者**: %s \n <at id=all></at>",
                	"tag": "lark_md"
            	}
        	}],
        	"header": {
            	"title": {
                	"content": "DevOps Platform",
        	        "tag": "plain_text"
            	}
        	}
    	}
	}`, *JobName, *BuildDisplayName, *Message, *Author))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Status:", resp.Status)
}
