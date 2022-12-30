package main

import (
	"fmt"
	"os"
	"time"
)

// PathError /
type PathError struct {
	fileName   string
	message    string
	createTime time.Time
}

func (p *PathError) Error() string {
	return fmt.Sprintf("fileName=%s \ncreateTime=%s \nmessage=%s", p.fileName, p.createTime, p.message)
}

// 打开文件
func optionFile(fileName string) error {
	open, err := os.Open(fileName)
	if err != nil {
		return &PathError{fileName: fileName,
			message:    "打开失败",
			createTime: time.Now(),
		}
	}
	defer open.Close()
	return nil
}

func main() {
	err := optionFile("C:\\Users\\jinjianghai\\Downloads\\abca.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("打开失败,", v)
	}
}
