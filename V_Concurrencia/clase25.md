# Channel con lectura y escritura

## ¿Cómo utilizar canales unidireccionales en Go?

La programación concurrente en Go es muy poderosa, especialmente cuando se usan canales para enviar y recibir datos. Sin embargo, no siempre es ideal usar canales bidireccionales. En algunos casos, definir claramente los tipos de canales, como aquellos usados solo para escribir o solo para leer, puede prevenir errores como bloqueos en las gorutinas.

## ¿Qué es un canal y cómo crear uno?

En Go, los canales permiten la comunicación entre gorutinas (funciones concurrentes), facilitando el paso de mensajes entre ellas. Los canales pueden cerrarse para indicar que no se enviarán más datos. Para cerrar un canal se usa la función `close`.

### Código básico de un canal en Go para enviar y recibir enteros

```go
c := make(chan int)
```

## ¿Cómo definir canales para solo lectura o solo escritura?

La especificación de si un canal es solo para lectura o solo para escritura se realiza mediante la dirección de una flecha (`<-`). Al establecer la dirección para la función que recibe o envía valores, se mitigan los posibles riesgos de un uso inapropiado.

### Ejemplo de definición de canal solo para escritura

```go
func generator(c chan<- int) {
    // código que envía valores
}
```

### Ejemplo de definición de canal solo para lectura

```go
func print(c <-chan int) {
    // código que lee valores
}
```

## ¿Cómo prevenir bloqueos con canales?

Go alertará si se intenta usar un canal de manera incorrecta según la definición establecida. Intentar usar un canal definido como solo de lectura para enviar datos genera un error en la compilación, ayudando a prevenir bloqueos.

## ¿Cuál es la estructura del pipeline con canales?

La estructura típica en un pipeline es conectar diferentes etapas de procesamiento a través de canales. Por ejemplo, un canal podría tomar números generados, duplicarlos y luego imprimirlos.

### Código para una estructura de pipeline básica

```go
func main() {
    generator := make(chan int)
    dobles := make(chan int)

    go generator(generator) // Genera valores y envía
    go double(generator, dobles) // Duplica valores y reenvía
    print(dobles) // Imprime valores
}
```

Cada parte del pipeline tiene su canal de comunicación, estableciendo un flujo ordenado entre funciones.

## ¿Por qué es importante entender los tipos de canal?

Entender cómo definir y utilizar tipos de canal específicos en tus programas es esencial para asegurar una correcta sincronización y comunicación entre gorutinas. Los errores pueden ser complejos de diagnosticar, ya que un canal mal utilizado puede obstruir la ejecución completa de tu programa.
