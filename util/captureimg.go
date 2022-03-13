package util

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
)

func CaptureStream(url string) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)

	err := ffmpeg.Input(url).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		panic(err)
		return nil, err
	}
	return buf, nil
}
