/*
 * Copyright (c) 2021 Jim Hoffman <jim@securebytes.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
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
