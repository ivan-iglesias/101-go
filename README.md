# GO

*Siguiendo el curso de Gentleman Programming, [Curso INTENSIVO de GO ! - te enseño TODO lo que tienes que saber](https://www.youtube.com/watch?v=z-NtGea-378)*

**Lenguaje de programación** desarrollado por Google cuya finalidad es resolver problemas de **concurrencia y escalabilidad** en sistemas modernos. Es altamente eficiente gracias a su **modelo de concurrencia incorporado**, su **compilación rápida** y su **sintaxis simple y clara**. Go es un lenguaje **compilado**, de **tipado estático**, diseñado para facilitar el desarrollo de software robusto y mantenible.

Usos comunes:

- Servidores web
- Herramientas de línea de comandos
- Aplicaciones web
- Servicios en la nube

[Ejemplos GO](https://gobyexample.com/)

## 🛠️ Instalación

Eliminar instalación previa

```
sudo rm -rf /usr/local/go
```

Extraer los archivos descargados de la [web oficial](https://go.dev/dl/)

```
sudo tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz
```

Añadir al final del fichero `~/.bashrc` la siguiente línea

```
export PATH=$PATH:/usr/local/go/bin
```

Con ello añadimos la ruta `/usr/local/go/bin` en la variable de entorno PATH. Para que surja efecto, cerramos y abrimos el terminal.

Confirmar que está instalado

```
go version
```

## 🚀 Primeros pasos

Crea el fichero `go.mod`, en el cual aparece el módulo `learning-go` y la versión que estamos usando.

```
go mod init learning-go
```

Creamos el archivo `main.go`, donde crearemos nuestro primer programa

```go
// Define el punto de entrada de nuestra aplicaicón
package main

// Importar un paquete de GO
import "fmt"

func main() {
  fmt.Println("Hola Mundo, soy Iván")
}

```

Lo ejecutamos con el siguiente comando, que buscara el fichero `main.go`

```
go run .
```

## 📦 Tipos de datos

|Tipo|Descripción|Valor por defecto|
|-|-|-|
|bool|true / false - flag o condicionales|false|
|string|cadena de caracteres|""|
|int, int8, int16, int32, int64|entero|0|
|uint, uint8, unit16, uint32, uint64|entero sin signo|0|
|float32, float64|representar valores númerico reales (decimales)|0|
|byte| === uint8, para trabajar con datos binarios|0|
|rune| === int32, representar un solo caracter que ocupa más de un byte|0|
|complex64, complex128|cuando tiene una parte real y una imaginaria (N + i)|0 + 0i|
