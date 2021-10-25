package main

import (
	"io/ioutil"
	"os"
	"bufio"
	"io"
	"strings"
	"golang.org/x/text/encoding/simplifiedchinese"
	"fmt"
)

func main() {
	bat, err := os.OpenFile("ts_merge.bat", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		//如果是.m3u8结尾的文件
		if !f.IsDir() && f.Name()[len(f.Name()) - 5:] == ".m3u8" {
			cmd, err := ts_merge(f.Name())
			if err != nil {
				fmt.Println(err)
				continue
			}
			bat.Write([]byte(cmd))
		}
	}
	bat.Close()
}

func ts_merge(m3u8 string) (cmd string, err error) {
	f, err := os.Open(m3u8)
	if err != nil {
		return
	}

	//commands
	cms := make([]string, 0)
	r := bufio.NewReader(f)
	var splSize int
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(bytes)
		if line[:7] == "file://" {
			spl := strings.Split(line, "/")
			if splSize == 0 {
				splSize = len(spl)
			}
			cms = append(cms, ".\\" + spl[splSize - 2] + "\\" + spl[splSize - 1])
		}
	}
	f.Close()
	// ANSI编码格式，避免cmd中文乱码
	fn, err := simplifiedchinese.GBK.NewEncoder().String(m3u8[:len(m3u8) - 5])
	if err != nil {
		return
	}
	cmd = "copy/b " + strings.Join(cms, " + ") + " \"" + fn + ".ts\" & "
	return
}
