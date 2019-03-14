package ini

import (
	"os"
	"bufio"
	"strings"
)

type M map[string]string

type Section M
type Conf map[string]Section

var (
	conf Conf = make(Conf)	// 初始化 conf
	currentSection string = "default"
)

func Read(filename string) (Conf, error) {
	file, err := os.Open(filename)
	if err != nil {
		return conf, err
	}
	defer file.Close()

	conf[currentSection] = make(Section)	// 初始化默认 section，conf['default']

	scanner := bufio.NewScanner(file)	// bufio.Scanner 实现逐行读取
	for {
		if ok := scanner.Scan(); !ok {	// 逐行扫描，文件尾或遇到错误会返回 false
			if err := scanner.Err(); err != nil {	// 返回第一个非 io.eof 的错误，如果是 io.eof，则返回的是 nil，如果有错误打印错误，并 exit
				return conf, err
			}
			break	// 执行到这说明文件读取完毕
		}	
		parse(scanner.Text())	// 返回最近一次 Scan 调用生成的文本
	}
	return conf, nil
}

func parse(line string) {
	if isNodeAnnotation(line) { 
		return 
	}
	if ok, section := isNodeSection(line); ok {
		currentSection = section
		if _, isExists := conf[currentSection]; !isExists {
			conf[currentSection] = make(Section)	// 如果该 section 不存在，则初始化
		}
		return
	}
	if ok, item := isNodeItem(line); ok {
		for k,v := range item {
			conf[currentSection][k] = v
		}
	}
}

func isNodeAnnotation(line string) bool {
	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
		return true
	}
	return false
}

func isNodeSection(line string) (bool, string) {
	if strings.HasPrefix(line, "[") {
		return true, strings.Trim(line, "[] ")	// 可以同时去除两端的"["、"]"、" "
	}
	return false, ""
}

func isNodeItem(line string) (bool, M) {
	item := strings.SplitN(line, "=", 2)
	if len(item) == 2 {
		return true, M{ strings.TrimSpace(item[0]) : strings.TrimSpace(item[1]) }
	}
	return false, nil	
}