package b64

import (
	"encoding/base64"
	"io"

	mc "github.com/multiformats/go-multicodec"
	base "github.com/multiformats/go-multicodec/base"
)

var (
	HeaderPath = "/base64/"
	Header     = mc.Header([]byte(HeaderPath))
	multic     = mc.NewMulticodecFromCodec(Codec(), Header)
)

type codec struct{}

type decoder struct {
	r io.Reader
}

func (d decoder) Decode(v interface{}) error {
	out, ok := v.(io.Writer)
	if !ok {
		return base.ErrExpectedWriter
	}

	_, err := io.Copy(out, d.r)
	return err
}

func (codec) Decoder(r io.Reader) mc.Decoder {
	return decoder{base64.NewDecoder(base64.StdEncoding, r)}
}

type encoder struct {
	w io.WriteCloser
}

func (e encoder) Encode(v interface{}) error {
	in, ok := v.(io.Reader)
	if !ok {
		return base.ErrExpectedReader
	}

	_, err := io.Copy(e.w, in)
	if err != nil {
		return err
	}
	return e.w.Close()
}

func (codec) Encoder(w io.Writer) mc.Encoder {
	return encoder{base64.NewEncoder(base64.StdEncoding, w)}
}

func Codec() mc.Codec {
	return codec{}
}

func Multicodec() mc.Multicodec {
	return multic
}
