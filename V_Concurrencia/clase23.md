# WaitGroup

## Sincronización de rutinas con Wait Groups en Go

En el ámbito de la programación concurrente, Go proporciona herramientas eficientes para gestionar múltiples tareas simultáneamente. Una de estas herramientas clave son los wait groups, que permiten sincronizar la ejecución de varias rutinas. En esta guía, te mostramos su funcionamiento y cómo utilizarlos en tus proyectos.

### ¿Qué es un Wait Group?

Un wait group en Go es un mecanismo que actúa como un contador del número de gorutinas que deben completar su trabajo antes de que el programa principal continúe. Es una alternativa efectiva a los canales cuando solo necesitas asegurarte de que todas tus tareas concurrentes han finalizado.

### Implementación de un Wait Group en Go

La implementación de un wait group en Go es sencilla y se realiza mediante los siguientes pasos:

1. **Crear un Wait Group**: Primero, importa el paquete `sync` que proporciona la estructura de wait group. Luego, crea una variable de tipo `sync.WaitGroup`.

    ```go
    var wg sync.WaitGroup
    ```

2. **Agregar al contador**: Cada vez que inicias una nueva gorutina, incrementa el contador del wait group con el método `Add`.

    ```go
    wg.Add(1)
    ```

3. **Disminuir el contador**: Dentro de cada gorutina, asegúrate de decrementar el contador del wait group cuando la gorutina finalice utilizando `Done`.

    ```go
    defer wg.Done()
    ```

4. **Esperar a que todas las gorutinas terminen**: Usa `wg.Wait()` en tu función principal para bloquearla hasta que todas las gorutinas hayan terminado su ejecución, es decir, hasta que el contador vuelva a cero.

    ```go
    wg.Wait()
    ```

### Ejemplo práctico de uso de Wait Groups

Para ilustrar cómo funciona un wait group, aquí tienes un ejemplo práctico donde se lanzan múltiples tareas que simulan un proceso prolongado:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func doSomething(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Goroutine %d inició\n", id)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d terminó\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go doSomething(i, &wg)
    }

    wg.Wait()
    fmt.Println("Todas las gorutinas han terminado.")
}
```

### Beneficios del uso de Wait Groups

- **Simplicidad**: Implementar wait groups es fácil y directo, facilitando la sincronización de tareas concurrentes.
- **Control**: Permite un control efectivo y seguro sobre la finalización de las gorutinas.
- **Legibilidad**: Mejora la legibilidad del código, haciendo evidente para otros programadores cuándo y cómo se sincronizan las tareas.
