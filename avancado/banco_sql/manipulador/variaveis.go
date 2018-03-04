package manipulador

import "html/template"

//ModeloOla armazena o modelo de pagina ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena o modelos de pagina Local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
