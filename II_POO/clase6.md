# Struc vs class

## Definición de clases en Go usando structs

En el ámbito de la programación, es esencial comprender cómo definir clases y sus propiedades para el desarrollo de software. Aunque lenguajes como Python y JavaScript tienen sus propias formas de definir clases, en Go utilizamos structs. Los structs nos permiten crear nuevos tipos de datos con propiedades específicas, actuando como plantillas para desarrollar objetos con características definidas.

## Comparación entre un struct en Go y una clase en Python

Un struct en Go es similar a una clase en Python en términos de estructura. En Python, definimos una clase y especificamos propiedades como `name` e `ID`. En Go, los structs cumplen el rol de las clases, permitiéndonos definir propiedades dentro del struct.

Por ejemplo, para crear un `Employee`:

```go
package main

type Employee struct {
    ID   int
    Name string
}
```

Aquí definimos un struct llamado `Employee` con dos propiedades: `ID` de tipo entero y `Name` de tipo string. Esto proporciona una estructura base para instanciar objetos `Employee` con esas características.

## Creación e instanciación de structs en Go

Después de definir un struct, el siguiente paso es instanciarlo, es decir, crear un objeto con las propiedades definidas.

En Python, crear una nueva instancia es sencillo gracias a su sintaxis directa utilizando la palabra clave `class`. En Go, instanciamos un struct usando llaves:

```go
func main() {
    employee := Employee{ID: 1, Name: "John"}
    fmt.Println(employee)
}
```

## Comportamiento por defecto de los structs en Go

Una característica interesante de Go es que cuando defines un struct sin asignar valores específicos, las propiedades toman valores por defecto.

- Los enteros (`int`) se inicializan en cero.
- Las cadenas de texto (`string`) se inicializan como cadenas vacías.

```go
employee := Employee{}
fmt.Println(employee) // Output: {0 ""}
```

Este comportamiento muestra que, incluso si no asignamos valores explícitamente durante la instancia, el lenguaje gestiona los valores predeterminados.

## Modificación de propiedades de un struct

Modificar las propiedades de un struct es tan directo como en otros lenguajes. Utilizamos la sintaxis de acceso a propiedades:

```go
employee.ID = 10
employee.Name = "Jane"
fmt.Println(employee) // Output: {10 "Jane"}
```

Esta metodología es similar a cómo se modificarían las propiedades en objetos de otros lenguajes, facilitando la adaptación para los programadores acostumbrados a lenguajes orientados a objetos.

Go ofrece diversas herramientas para trabajar con structs, asemejándose a la funcionalidad de clases en Python o TypeScript. Sin embargo, su enfoque es simple y eficiente, lo que lo hace atractivo para aquellos que buscan escribir código limpio y ordenado. En la próxima lección, exploraremos cómo agregar métodos a los structs para ampliar su funcionalidad, llevando a Go a niveles de abstracción similares a los lenguajes orientados a objetos tradicionales.
