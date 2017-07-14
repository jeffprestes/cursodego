package manipulador

import "html/template"

//Modelos armazena os modelos de pagina que serão executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
