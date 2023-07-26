package main

import (
	"fmt"
	"os"
	//"os/exec"
)

func convertToDash(mp4FileName string) {
	// Verificar si se proporcionó un archivo MP4 como argumento
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run main.go archivo.mp4")
		os.Exit(1)
	}

	
	// outputFileName := "./pb_data/output.mpd" // Nombre del archivo DASH de salida

	// // Comando de conversión usando ffmpeg
	// cmd := exec.Command("ffmpeg",
	// 	"-i", mp4FileName,
	// 	"-c:v", "libx264",
	// 	"-c:a", "aac",
	// 	"-b:v", "4000k",
	// 	"-b:a", "128k",
	// 	"-map", "0",
	// 	"-f", "dash",
	// 	"-window_size", "10",
	// 	"-extra_window_size", "5",
	// 	"-min_seg_duration", "5000000",
	// 	"-use_template", "1",
	// 	"-use_timeline", "1",
	// 	"-init_seg_name", "init-stream$RepresentationID$.m4s",
	// 	"-media_seg_name", "chunk-stream$RepresentationID$-$Number%05d$.m4s",
	// 	outputFileName,
	// )

	// // Redirigir la salida estándar y de error
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// // Ejecutar el comando
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("Error durante la conversión:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Conversión completada exitosamente.")
	
}
