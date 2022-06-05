package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func push(m string) {
	params := url.Values{}
	params.Add("token", token)
	params.Add("user", user)
	params.Add("title", title)
	params.Add("message", m)

	resp, err := http.PostForm("https://api.pushover.net/1/messages.json", params)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}

func main() {

	if title == "" {
		title, _ = os.Hostname()
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		reader.Scan()
		line := reader.Text()
		if len(line) == 0 {
			// exits if \n is sent without a message
			time.Sleep(time.Second)
			break
		}
		go push(line)
	}
}
