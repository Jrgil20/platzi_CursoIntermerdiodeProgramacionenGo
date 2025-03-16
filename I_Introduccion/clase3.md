# Repaso General: variables, condicionales, slices y maps

## ¿Qué conceptos repasaremos en este curso?

Comenzaremos con una revisión general de conceptos clave de cursos anteriores que son esenciales para avanzar en este curso. Estos principios nos ayudarán a entender cómo se integran las nuevas características de Go.

### Temas específicos a revisar

- **Structs**: Cómo definirlos y utilizarlos.
- **Apuntadores**: Uso y significado del ampersand (&) y del asterisco (*).
- **Google Teams**: Creación y uso.
- **Canales (Channels)**: Comunicación entre Go Routines y control del programa según sea necesario.

## ¿Cómo preparar nuestro entorno de desarrollo con Go?

Para este curso, utilizaremos Visual Studio Code (VS Code) como editor de texto, y es importante tener instalada la extensión de Go. Configura tu entorno siguiendo estos pasos:

1. Abre VS Code y busca en la pestaña de extensiones “Go”.
2. Instala la extensión adecuada para asegurar una experiencia fluida en el desarrollo de tus aplicaciones.

## ¿Cómo estructura un programa básico en Go?

Para iniciar con un programa básico en Go, sigue estos pasos:

1. Crea un archivo nuevo llamado `main.go`.
2. Define un paquete principal:

```go
package main

func main() {
    // Punto de entrada del programa
}
```

Este archivo será el punto de entrada de tu aplicación, crucial para su compilación y ejecución.

## ¿Cómo se declaran variables en Go?

Go ofrece múltiples maneras de declarar variables, tanto de forma explícita como implícita:

- **Explícita**: Utiliza `var`, el nombre y el tipo de la variable.

```go
var x int = 8
```

- **Implícita**: Mediante el uso de `:=`.

```go
y := 8
```

Ambas formas tienen su uso, dependiendo de si deseas ser explícito respecto al tipo de tu variable para evitar errores de compatibilidad.

## ¿Cómo se manejan los errores en Go?

A diferencia de otros lenguajes, Go maneja los errores de manera explícita, sin las tradicionales estructuras `try` y `catch`.

Al intentar parsear un entero desde un string:

```go
myValue, err := strconv.ParseInt("8", 0, 64)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Parsed value:", myValue)
}
```

El manejo explícito permite identificar claramente el error que se produce y gestionarlo adecuadamente.

## ¿Cómo crear y manejar mapas en Go?

Los mapas en Go son estructuras clave-valor que se definen como sigue:

```go
m := make(map[string]int)
m["clave"] = 6
fmt.Println(m["clave"])
```

Asegúrate de que el tipo de las llaves y los valores coincida con lo especificado al crear el mapa para evitar errores de compilación.

## ¿Cómo utilizar slices en Go?

Los slices son similares a arreglos pero con más flexibilidad. Se crean y se pueden iterar de la siguiente manera:

```go
s := []int{1, 2, 3}
for index, value := range s {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

Puedes agregar más elementos usando la función `append`:

```go
s = append(s, 16)
```

Esto te permite modificar la longitud del slice dinámicamente.
