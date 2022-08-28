package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

// GobCodec Gob类型编解码
type GobCodec struct {
	conn io.ReadWriteCloser // 通常是通过TCP或者Unix建立Socket时得到的链接实例
	buf  *bufio.Writer      // 防止阻塞而创建的带缓冲的Writer
	dec  *gob.Decoder       // Decoder
	enc  *gob.Encoder       // Encoder
}

var _ Codec = (*GobCodec)(nil) // 检查GobCodec是否实现Codec接口

// NewGobCodec 根据连接生成Gob类型编解码器
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

// ReadHeader 解析Header
func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

// ReadBody 解析Body
func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

// Write 写回内容
func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()

	if err = c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

// Close 关闭连接
func (c *GobCodec) Close() error {
	return c.conn.Close()
}
