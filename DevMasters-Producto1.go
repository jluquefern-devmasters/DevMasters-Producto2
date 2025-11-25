package main

import (
    "fmt"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, `
            <html>
            <body>
                <h1>Soy alumno de la UOC y miembro de DevMasters</h1>
                <img src="/static/UOCLogo.png" width="300">
            </body>
            </html>
        `)
    })

    fmt.Println("Servidor ejecutándose en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
