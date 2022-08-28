package codec

import "io"

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}

// Codec 编解码接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// NewCodecFunc Codec构造函数类型
type NewCodecFunc func(closer io.ReadWriteCloser) Codec

// Type 编解码类型
type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

// NewCodecFuncMap 存放编解码类型与其工厂方法
var NewCodecFuncMap map[Type]NewCodecFunc

// Header 存放除参数之外的请求信息
type Header struct {
	ServiceMethod string // 服务名和方法名，通常与 Go 语言中的结构体和方法相映射
	Seq           uint64 // 请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求
	Error         string // 错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中
}
