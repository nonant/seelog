// Copyright (c) 2012 - Cloud Instruments Co., Ltd.
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package seelog

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// fileWriter is used to write to a file.
type httpWriter struct {
	innerWriter io.WriteCloser
	url         string
}

// Creates a new file and a corresponding writer. Returns error, if the file couldn't be created.
func NewHttpWriter(url string) (writer *httpWriter) {
	newWriter := new(httpWriter)
	newWriter.url = url

	return newWriter
}

func (fw *httpWriter) Close() error {
	// if fw.innerWriter != nil {
	// 	err := fw.innerWriter.Close()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fw.innerWriter = nil
	// }
	return nil
}

// Create folder and file on WriteLog/Write first call
func (fw *httpWriter) Write(bytes []byte) (n int, err error) {
	if fw.url != "" {
		messagebody := ioutil.NopCloser(strings.NewReader(string(bytes)))

		client := &http.Client{}
		r, err := http.NewRequest("POST", fw.url, messagebody)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}

		r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

		resp, err := client.Do(r)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
		defer resp.Body.Close()
	}
	return 0, nil
}

func (fw *httpWriter) String() string {
	return fmt.Sprintf("Http writer: %s", fw.url)
}
