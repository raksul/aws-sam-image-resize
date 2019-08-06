package main

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/disintegration/imaging"
)

func Convert(r io.Reader, c *Config) (string, error) {
	srcImg, err := imaging.Decode(r)
	if err != nil {
		return "", err
	}

	dstImg := imaging.Resize(srcImg, c.Width, c.Height, imaging.Lanczos)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, dstImg, imaging.PNG)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
