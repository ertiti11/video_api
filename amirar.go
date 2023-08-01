package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Definir la URL base que deseas utilizar en la etiqueta <BaseURL>
	baseURL := "http://ejemplo.com/carpeta/segmentos/"

	// Definir una plantilla del archivo MPD con el marcador de posición para la URL base
	mpdTemplate := `
<MPD xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns="urn:mpeg:dash:schema:mpd:2011"
	xmlns:xlink="http://www.w3.org/1999/xlink"
	xsi:schemaLocation="urn:mpeg:DASH:schema:MPD:2011 http://standards.iso.org/ittf/PubliclyAvailableStandards/MPEG-DASH_schema_files/DASH-MPD.xsd"
	profiles="urn:mpeg:dash:profile:isoff-live:2011"
	type="static"
	mediaPresentationDuration="PT40.5S"
	maxSegmentDuration="PT5.0S"
	minBufferTime="PT20.8S">
	<Period id="0" start="PT0.0S">
		<AdaptationSet id="0" contentType="video" startWithSAP="1" segmentAlignment="true" bitstreamSwitching="true" frameRate="24000/1001" maxWidth="640" maxHeight="360" par="16:9" lang="und">
			<Representation id="0" mimeType="video/mp4" codecs="avc1.64001e" bandwidth="1000000" width="640" height="360" sar="1:1">
				<SegmentTemplate timescale="24000" initialization="init-stream$RepresentationID$.m4s" media="chunk-stream$RepresentationID$-$Number%05d$.m4s" startNumber="1">
					<SegmentTimeline>
						<S t="0" d="250250" r="2" />
						<S d="222222" />
					</SegmentTimeline>
				</SegmentTemplate>
			</Representation>
		</AdaptationSet>
	</Period>
</MPD>
`

	// Reemplazar el marcador de posición en la plantilla con la URL base real
	mpdContent := strings.ReplaceAll(mpdTemplate, "$BaseURL$", baseURL)

	// Guardar el contenido del MPD en un archivo temporal
	tmpFile := "temp.mpd"
	err := os.WriteFile(tmpFile, []byte(mpdContent), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo temporal:", err)
		return
	}
	defer os.Remove(tmpFile)

	// Generar el archivo MPD utilizando ffmpeg
	cmd := exec.Command("ffmpeg", "-i", "video.mp4", "-f", "dash", "output.mpd")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar el archivo MPD:", err)
		return
	}

	// Mover el archivo MPD generado al directorio deseado
	err = os.Rename("output.mpd", "videos/output/output.mpd")
	if err != nil {
		fmt.Println("Error al mover el archivo MPD:", err)
		return
	}

	fmt.Println("Archivo MPD generado y movido con éxito.")
}
