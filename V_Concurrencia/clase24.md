# Bufered Channels ( como semaforos)

## ¿Qué son los channels con buffer en Go?

En el ámbito de la programación concurrente en Go, los channels con buffer son fundamentales. Funcionan como "semáforos" que limitan la cantidad de procesos concurrentes, regulando el flujo de ejecución de las goroutines. Al definir un buffer, se controla cuántas goroutines pueden estar activas simultáneamente, evitando bloqueos por sobrecarga.

## ¿Cómo se implementa un channel con buffer?

Para ilustrar el uso de channels con buffer, veamos un ejemplo práctico. Primero, se crea un canal para transmitir datos de tipo entero y se define un buffer que permite manejar múltiples datos antes de ser leídos.

### Código de ejemplo

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func doSomething(id int, wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    fmt.Printf("Inicio del proceso: %d\n", id)
    time.Sleep(4 * time.Second)
    fmt.Printf("Fin del proceso: %d\n", id)
    <-ch
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan int, 2)

    for i := 0; i < 10; i++ {
        wg.Add(1)
        ch <- 1
        go doSomething(i, &wg, ch)
    }

    wg.Wait()
}
```

### Explicación del código

- **Creación del canal**: Se crea un canal `ch` con un buffer de tamaño 2, permitiendo contener hasta dos elementos antes de requerir lectura.
- **Función `doSomething`**: Simula un proceso que tarda cuatro segundos en completarse, utilizando el canal para indicar la finalización de una goroutine y liberar espacio en el buffer.
- **Control del flujo concurrente**: Con un ciclo `for`, se lanzan diez goroutines, pero solo dos se ejecutan simultáneamente debido al buffer del canal. Las demás esperan su turno.

## ¿Cómo actúa el buffer como un semáforo?

El buffer controla el ritmo de ejecución de las goroutines. En el código anterior:

- Cada vez que `doSomething` finaliza, se libera un elemento del canal (`<-ch`), permitiendo que una nueva goroutine inicie.
- Con un buffer de tamaño 2, solo dos goroutines se ejecutan simultáneamente, gestionando eficientemente los recursos.

## ¿Qué sucede si modificamos la capacidad del buffer?

Aumentar la capacidad del buffer permite más procesos concurrentes. Por ejemplo, con un buffer de 5, hasta cinco goroutines pueden ejecutarse simultáneamente antes de requerir liberar espacio.

### Prueba con un buffer mayor

```go
ch := make(chan int, 5)
```

Al ejecutar el programa con esta configuración, más goroutines inician y finalizan en paralelo, ideal para balancear concurrencia y control.

## ¿Cuándo utilizar channels con buffer?

Los channels con buffer son esenciales en escenarios donde la carga de procesos concurrentes puede saturar el sistema:

- **Control de concurrencia**: Limitar goroutines concurrentes previene sobrecargas y asegura un uso eficiente de recursos.
- **Trabajo en colas (job queues)**: Gestionan la ejecución de trabajos acumulados rápidamente.

Aprovechando los channels con buffer, puedes optimizar la concurrencia en tus programas en Go, asegurando rendimiento y estabilidad.
