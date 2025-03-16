# Repaso general: Gorutines y apuntadores

## ¿Qué son y cómo usar correctamente las Go Routines y canales en Go?

En el ámbito de la programación concurrente, las Go Routines y los canales son componentes cruciales que pueden optimizar el rendimiento de tus proyectos en Go. Comprender su funcionamiento te permitirá desarrollar aplicaciones más eficientes y robustas. A continuación, se detallan sus fundamentos, cómo implementarlos en tu código, y se concluye con los conceptos de apuntadores para ofrecer una visión más completa de su uso práctico.

### ¿Cómo funcionan las Go Routines y su implementación?

Las Go Routines son la herramienta principal en Go para ejecutar tareas de manera concurrente. Puedes pensar en ellas como funciones que se ejecutan de forma asíncrona en el mismo hilo o en hilos diferentes, permitiéndote realizar múltiples tareas simultáneamente sin bloquear el flujo del programa principal.

Para iniciar una Go Routine, simplemente utiliza la palabra clave `go` antes de llamar a la función deseada:

```go
func DoSomething() {
    time.Sleep(3 * time.Second)
    fmt.Println("done")
}

func main() {
    go DoSomething()
    fmt.Println("This will print before 'done' due to asynchrony")
}
```

Es importante notar que una Go Routine por sí sola no es monitoreada, lo que significa que el mensaje "done" podría no imprimirse si `main` termina antes de que `DoSomething` se complete. Aquí es donde los canales son útiles.

### ¿Qué rol juegan los canales en la comunicación entre Go Routines?

Los canales en Go permiten que múltiples Go Routines se comuniquen entre sí. Un canal es esencialmente una tubería a través de la cual puedes enviar y recibir valores entre diferentes partes de tu programa.

Para interactuar mediante un canal, primero debes crearlo y luego modificar tu función para que acepte el canal como parámetro:

```go
func DoSomething(c chan int) {
    time.Sleep(3 * time.Second)
    fmt.Println("done")
    c <- 1 // Envía un valor al canal
}

func main() {
    c := make(chan int)
    go DoSomething(c)
    <-c // Espera a recibir un valor del canal antes de continuar
}
```

En este ejemplo, `main` se bloqueará hasta que `DoSomething` envíe un valor al canal, asegurando que el mensaje "done" se imprima antes de que el programa finalice por completo.

### ¿Cuál es la importancia de los apuntadores en Go?

Los apuntadores en Go son esenciales cuando necesitas manejar un valor por su referencia, en lugar de por su copia. Esto es útil para optimizar el uso de memoria y permitir que las funciones modifiquen el estado de una variable fuera de su ámbito local.

Imagina que tienes la variable `g` y quieres manejar su dirección de memoria:

```go
var g = 25
fmt.Println(g) // Imprime 25

var h *int = &g // h es un apuntador a la dirección de g
fmt.Println(h)  // Imprime la dirección de memoria de g
```

Para acceder al valor real donde apunta `h`, utilizamos el operador asterisco (`*`):

```go
fmt.Println(*h) // Imprime 25, el valor almacenado en la dirección apuntada por h
```

Es crucial distinguir entre el uso del operador dirección (`&`) para referirse a la dirección de memoria de una variable, y el operador asterisco para acceder al valor almacenado en esa dirección. Este conocimiento te ayudará a evitar posibles errores e inconvenientes en el manejo de memoria en Go.
