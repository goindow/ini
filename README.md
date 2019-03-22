# ini
Package toolbox 封装了一个 .ini 格式文件读取器

## 特性
- 注释符 ";"、"#"
- 键值分割符 "="
- 支持 section，在 [section] 后的内容都属于该 section，默认的 section 为 default
- 重复定义 section，不同键会追加到该 section 中，相同则后者覆盖前者

## 索引
- [Read](#Read)

### 使用
```go
import (
	"github.com/goindow/ini"
)
```

### Read
- 如果 err != nil，则返回值一个空 map
```go
file := ""

if conf, err := ini.read(file); err != nil {
	fmt.Println(conf)
}

// Output:
// map[default:map[a:1 b:2 c:1 + 1 = 2] user:map[gender:male name:hyb age:99] profile:map[email:76788424@qq.com github:github.com/goindow]]
```