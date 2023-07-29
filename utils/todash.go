package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ToDash(source string, id string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error al obtener el directorio de trabajo actual: %v", err)
	}
	inputFile := filepath.Join(cwd, "pb_data/storage/4ronlqa5jkr2oda", id, source)
	outputDir := filepath.Join(cwd, "pb_data/storage/4ronlqa5jkr2oda", id)
	outputFileName := "outlast.mpd"

	// Crear el directorio de salida si no existe
	// if _, err := os.Stat(outputDir); os.IsNotExist(err) {
	// 	err := os.Mkdir(outputDir, os.ModePerm)
	// 	if err != nil {
	// 		log.Fatalf("Error al crear el directorio de salida: %v", err)
	// 	}
	// }

	// Comando de ffmpeg para la conversión DASH

	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "libx264", "-b:v", "1M", "-c:a", "aac", "-b:a", "128k", "-f", "dash",
		"-init_seg_name", filepath.Join("pb_data/storage/4ronlqa5jkr2oda", id, "$RepresentationID$"),
		"-media_seg_name", filepath.Join("pb_data/storage/4ronlqa5jkr2oda", id, "segment-$Number%09d$.m4s"),
		filepath.Join(outputDir, outputFileName))

	// Capturar la salida del comando en caso de que haya errores
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error al ejecutar ffmpeg: %v\nSalida de ffmpeg:\n%s", err, output)
	}

	fmt.Println("Conversión completada con éxito.")

}

func Leer(videoname string) {
	BASE_VIDEO_PATH := filepath.Join("pb_data/storage/4ronlqa5jkr2oda")

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
