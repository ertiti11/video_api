package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"video_api/utils"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func execute_transform(id string, source string, title string) {
	// Crea un objeto de tipo Cmd
	cmd := exec.Command("python", "../mp4ToDash.py", source, id, title)
	// Redirecciona la salida del comando a la salida estándar de la consola
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Ejecuta el comando
	if err := cmd.Run(); err != nil {
		// Si ocurre un error, imprímelo
		fmt.Println("Error:", err)
		return
	}
}

func main() {
	utils.Leer("video")
	BASE_VIDEO_PATH := filepath.Join("backend\\pb_data\\storage\\4ronlqa5jkr2oda")
	fmt.Println(BASE_VIDEO_PATH)
	app := pocketbase.New()

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		execute_transform(e.Record.Get("id").(string), e.Record.Get("source").(string), e.Record.Get("titulo").(string))

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// or you can also use the shorter e.Router.GET("/articles/:slug", handler, middlewares...)

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/:video/dash_out.mpd",
			Handler: func(c echo.Context) error {

				return c.String(1, "hola que ase")
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})
	if err := app.Start(); err != nil {
		print("error")
	}
}
