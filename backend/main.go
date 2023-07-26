package main

import (
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/labstack/echo/v5"
    "net/http"
    "github.com/pocketbase/pocketbase/core"
    "fmt"
    "os"
    "os/exec"
)

func execute_transform(id string, source string, title string) {
    // Crea un objeto de tipo Cmd
    cmd := exec.Command("python3", "../mp4ToDash.py", source, id, title)
    fmt.Println("python3", "mp4ToDash.py", source, id, title)
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
    app := pocketbase.New()
    // &{{true 0afqhmo08psimyq 2023-07-20 23:31:53.191Z 2023-07-20 23:31:53.191Z} 0xc000196340 false false true map[source:video_uDMNls7SIC.mp4 titulo:lfaksjdflñkajdfñlkj32490u09234] <nil> 0xc0000b4180}
    
    app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
        //get title of the previous comment
        execute_transform(e.Record.Get("id").(string), e.Record.Get("source").(string), e.Record.Get("titulo").(string))
       

       return nil
    })


    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        // or you can also use the shorter e.Router.GET("/articles/:slug", handler, middlewares...)
        e.Router.AddRoute(echo.Route{
            Method: http.MethodGet,
            Path:   "/api/articles/:slug",
            Handler: func(c echo.Context) error {
                return c.String(http.StatusOK, c.PathParam("slug"))               
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

