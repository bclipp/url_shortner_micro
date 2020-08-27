package msgpack

import (
	"github.com/vmihailenco/msgpack"
	"github.com/bclipp/url_shortner_micro/go/src/urlshortner"
	"github.com/pkg/errors"

)

type Redirect struct {}

func (r *Redirect) Decode(input []byte) (*urlshortner.Redirect, error) {
	redirect := &urlshortner.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect,nil
}

func (r * Redirect) Encode(input *urlshortner.Redirect)([]byte, error){
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
