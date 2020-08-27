package json

import (
	"encoding/json"
	"github.com/bclipp/url_shortner_micro/go/src/urlshortener"
	"github.com/pkg/errors"

)

type Redirect struct {}

func (r *Redirect) Decode(input []byte) (*urlshortener.Redirect, error) {
	redirect := &urlshortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect,nil
}

func (r * Redirect) Encode(input *urlshortener.Redirect)([]byte, error){
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
