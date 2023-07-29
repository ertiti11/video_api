package main

import (
	"fmt"
	"io"
	"log"
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
		utils.ToDash(e.Record.Get("source").(string), e.Record.Get("id").(string))
		// execute_transform(e.Record.Get("id").(string), e.Record.Get("source").(string), e.Record.Get("titulo").(string))

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// or you can also use the shorter e.Router.GET("/articles/:slug", handler, middlewares...)
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error al obtener el directorio de trabajo actual: %v", err)
		}
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/:id",
			Handler: func(c echo.Context) error {
				videopath := filepath.Join(cwd, "pb_data/storage/4ronlqa5jkr2oda", c.PathParam("id"), "outlast.mpd")
				mpdFile, err := os.Open(videopath)
				if err != nil {
					log.Fatal("error", err)
					return nil
				}
				defer mpdFile.Close()

				mpdData, err := io.ReadAll(mpdFile)
				if err != nil {
					log.Fatal("error:", err)
					return nil
				}
				c.Response().Header().Set("Content-Type", "application/dash+xml")
				c.Response().Write(mpdData)

				return c.String(200, "")
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// or you can also use the shorter e.Router.GET("/articles/:slug", handler, middlewares...)
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error al obtener el directorio de trabajo actual: %v", err)
		}
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/:id/:stream_id",
			Handler: func(c echo.Context) error {
				// fmt.Sprintf("%s.m4s", c.PathParam(":stream_id")
				initSegmentPath := filepath.Join(cwd, c.PathParam("id"), c.PathParam("stream_id"))

				initSegmentFile, err := os.Open(initSegmentPath)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error al abrir el segmento de inicialización: %s", err.Error()))
				}
				defer initSegmentFile.Close()
				initSegmentData, err := io.ReadAll(initSegmentFile)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error al leer el segmento de inicialización: %s", err.Error()))
				}
				return c.Blob(http.StatusOK, "video/mp4", initSegmentData)

			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// or you can also use the shorter e.Router.GET("/articles/:slug", handler, middlewares...)
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error al obtener el directorio de trabajo actual: %v", err)
		}
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/:id/:segment_number",
			Handler: func(c echo.Context) error {
				segmentFilename := fmt.Sprintf("segment-%s.m4s", c.PathParam("segmentNumber"))
				segmentPath := filepath.Join(cwd, c.PathParam("id"), segmentFilename)

				segmentFile, err := os.Open(segmentPath)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error al abrir el segmento: %s", err.Error()))
				}
				defer segmentFile.Close()
				segmentData, err := io.ReadAll(segmentFile)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error al leer el segmento: %s", err.Error()))
				}

				return c.Blob(http.StatusOK, "video/mp4", segmentData)
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
