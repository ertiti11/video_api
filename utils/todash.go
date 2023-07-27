package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ToDash(videoName string)


func Leer(videoname string) {
	BASE_VIDEO_PATH := filepath.Join("backend\\pb_data\\storage\\4ronlqa5jkr2oda")

	mpdPath := filepath.Join(BASE_VIDEO_PATH, videoname, "dash_out.mpd")
	mpdFile, err := os.Open(mpdPath)
	if err != nil {
		print("error")
	}
	defer mpdFile.Close()
	mpdData, err2 := io.ReadAll(mpdFile)
	if err2 != nil {
		print("error")
	}
	fmt.Println(mpdData)
	// c.Response().Header().Set("Content-Type", "application/dash+xml")
	// _, err = c.Response().Write(mpdData)
}
