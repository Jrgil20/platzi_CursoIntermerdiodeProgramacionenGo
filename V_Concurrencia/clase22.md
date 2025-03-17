# Unbuffered  channel y Buffered channel

## ¿Qué son los canales en Go?

Los canales en Go son una herramienta esencial para la comunicación entre Go routines, permitiendo la transferencia de datos de manera segura y sincronizada. Son cruciales para la concurrencia en Go, y comprender su funcionamiento es vital para evitar bloqueos y errores en el código.

## Tipos de canales en Go

En Go, existen dos tipos de canales que se distinguen por la presencia o ausencia de un buffer. Cada uno tiene casos de uso específicos y presenta ventajas y desafíos únicos en la programación concurrente.

### Canal sin buffer

Un canal sin buffer no tiene capacidad para almacenar mensajes, lo que provoca un bloqueo si no hay una Go routine lista para recibir el mensaje en el momento en que se envía. Este tipo de canal es útil cuando la comunicación debe ser inmediata y sin retrasos.

```go
package main

func main() {
    ch := make(chan int)
    ch <- 1
    fmt.Println(<-ch)
}
```

En el ejemplo anterior, después de definir `ch` como un canal para enteros, el intento de enviar `1` produce un bloqueo porque no hay ninguna Go routine lista para recibir el mensaje.

### Canal con buffer

Un canal con buffer permite almacenar un número específico de mensajes antes de bloquear. Esto es útil cuando se necesita cierta flexibilidad en el envío y recepción de mensajes, permitiendo que una parte del programa continúe su ejecución sin esperar.

```go
package main

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

En este ejemplo, el canal `ch` tiene espacio para tres valores, por lo que no se bloquea hasta superar el límite del buffer.

## Limitaciones de los canales

Los canales, tanto con como sin buffer, pueden llevar a bloqueos si no se gestionan adecuadamente. Es esencial asegurarse de que:

- Un canal sin buffer siempre tenga una Go routine preparada para recibir antes de enviar un mensaje.
- No se supere la capacidad de un canal con buffer sin tener el control adecuado para procesar o recibir los mensajes.

Cuando se sobrepasa la capacidad de un canal con buffer, como al intentar enviar un cuarto valor en un buffer de capacidad tres, el programa se bloqueará nuevamente.

## Alternativas para manejar concurrencia

Además de los canales, Go ofrece otros mecanismos para manejar concurrencia que serán explorados en lecciones futuras. Estos métodos pueden ser más simples o intuitivos dependiendo del contexto de uso, y ofrecen flexibilidad para optimizar la comunicación y sincronización entre Go routines.

## Recomendaciones prácticas

Para dominar el uso de canales en Go, ten en cuenta:

- Entender el contexto y uso específico de cada tipo de canal.
- Practicar con distintas configuraciones para identificar los momentos en los que un canal se bloqueará.
- Explorar alternativas a los canales para mejorar la eficiencia y rendimiento de tu aplicación concurrente.
