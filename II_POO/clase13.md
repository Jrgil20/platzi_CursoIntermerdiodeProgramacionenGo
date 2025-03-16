# funciones Anonimas

## ¿Qué es una función anónima en Go?

Las funciones anónimas, también llamadas funciones sin nombre, son un concepto interesante en programación y Go las implementa de manera efectiva. Estas funciones se crean para ser usadas una sola vez y no se pueden referenciar nuevamente, ya que no tienen nombre. Este enfoque es útil cuando estamos seguros de que la funcionalidad solo se necesita ejecutar una vez.

## ¿Cómo implementamos funciones anónimas en Go?

Veamos cómo se utilizan las funciones anónimas en Go. Primero, creamos un archivo llamado `functions.go`. Dentro de este archivo, definimos el paquete principal y una función `main`. Creamos una variable `x` e implementamos una función anónima que duplica su valor.

```go
package main

func main() {
    x := 5
    y := func() int {
        return x * 2
    }()
    println(y)
}
```

En este bloque, la función que obtiene el doble del valor de `x` se declara e invoca inmediatamente. La función anónima se asigna a una variable `y`, que se imprime a continuación. Sin embargo, si necesitamos reutilizar esta función con otro valor, corremos el riesgo de duplicar código.

## ¿Cuáles son las desventajas de repetir código con funciones anónimas?

El ejemplo anterior muestra que si intentamos reutilizar el código de la función con otra variable, estaríamos repitiendo código, lo cual va en contra del principio DRY (Don't Repeat Yourself). Este principio promueve evitar la duplicación en el diseño de software.

```go
package main

func main() {
    x := 5
    y := func() int {
        return x * 2
    }()
    println(y)

    z := 10
    doubleZ := func() int {
        return z * 2
    }()
    println(doubleZ)
}
```

Aunque este enfoque puede funcionar, no es una buena práctica. Repetiríamos el mismo bloque de código en lugar de abstraerlo en una función.

## ¿Cómo se utilizan funciones anónimas con Goroutines?

Otra utilidad clave de las funciones anónimas es en la ejecución concurrente con Goroutines en Go. Las funciones anónimas nos permiten realizar procesos largos sin bloquear el flujo general del programa.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan bool)
    go func() {
        fmt.Println("Iniciando la función...")
        time.Sleep(5 * time.Second)
        fmt.Println("Función terminada.")
        done <- true
    }()
    <-done
    fmt.Println("Programa completado.")
}
```

En este ejemplo, hemos usado una función anónima para simular un proceso largo que espera 5 segundos antes de finalizar. Esto demuestra la potencia de las funciones anónimas al permitirnos ejecutar código de manera concurrente usando Goroutines.

## ¿Cuándo es recomendable usar funciones anónimas?

El uso de funciones anónimas es práctico en situaciones donde:

- Necesitamos realizar una operación sencilla una sola vez.
- Deseamos aplicar funcionalidad en tiempo de ejecución sin saturar el espacio de nombres con funciones no reutilizables.
- En conjunción con Goroutines para manejar procesos concurrentes.

Sin embargo, es crucial recordar que contar con bloques de código duplicado complica el mantenimiento del mismo y no es recomendado. Por lo tanto, el diseño en Go debe tener en cuenta el equilibrio entre funcionalidad, simplicidad y reutilización.
