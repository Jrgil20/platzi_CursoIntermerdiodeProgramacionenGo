# Profiling en Go

## ¿Qué es el profiling y cómo puede mejorar tu código?

El profiling es una técnica esencial que permite identificar las áreas críticas de tu código, ayudándote a optimizar su rendimiento. Los desarrolladores a menudo se enfrentan a la incertidumbre sobre por qué sus programas son lentos o consumen demasiados recursos. Con el profiling, puedes detectar exactamente qué partes del código necesitan atención para mejorar la eficiencia.

## ¿Cómo implementar una función de Fibonacci recursiva en Go?

Para calcular la serie de Fibonacci de manera recursiva en Go, puedes agregar la siguiente función. Esta función toma un número entero como parámetro y devuelve otro entero.

```go
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}
```

Aquí, se establece el caso base de la recursión: si `n` es menor o igual a uno, simplemente retorna `n`. De lo contrario, la función se llama a sí misma con `n-1` y `n-2`, sumando sus resultados.

## ¿Cómo crear pruebas para la función de Fibonacci en Go?

Para asegurarte de que la función de Fibonacci funcione correctamente, puedes crear pruebas unitarias en Go. Estas pruebas verifican diferentes escenarios, comprobando que la función devuelve los resultados esperados.

```go
func TestFibonacci(t *testing.T) {
    var tests = []struct{
        a, n int
    }{
        {1, 1},  // Fibonacci(1) = 1
        {8, 21}, // Fibonacci(8) = 21
        {50, 12586269025}, // Fibonacci(50)
    }

    for _, tt := range tests {
        result := Fibonacci(tt.a)
        if result != tt.n {
            t.Errorf("Fibonacci(%d) = %d; want %d", tt.a, result, tt.n)
        }
    }
}
```

El paso final es crear una tabla de casos que suma los diferentes inputs y outputs esperados, e iterar sobre ellos para verificar que la función arroja el resultado correcto.

## ¿Cómo utilizar el profiling en Go?

Dado que ejecutar las pruebas para Fibonacci de 50 toma mucho tiempo, aquí es donde entra en juego el profiling. Utiliza Go para generar un archivo de perfil que te ayudará a entender mejor dónde se va el tiempo de ejecución.

Ejecuta tu test con el comando de profiling y genera un archivo de salida:

```sh
go test -cpuprofile cpu.out
```

Analiza el archivo generado usando `go tool` para observar el time profiling:

```sh
go tool pprof cpu.out
```

Una vez dentro del modo interactivo, puedes usar `top` para ver las funciones que más tiempo están consumiendo:

```sh
top
```

La función de Fibonacci probablemente estará en la parte superior de la lista, debido a su implementación recursiva ineficiente.

## ¿Cómo interpretar y compartir los resultados del profiling?

Para entender en profundidad los resultados, puedes listar en detalle los tiempos de cada línea dentro de la función de Fibonacci:

```sh
list Fibonacci
```

Además, es posible exportar los resultados a un formato web o PDF para compartir con tu equipo, lo cual puede ser muy útil para colaborar en mejoras.

Para exportar como SVG:

```sh
web
```

Para exportar como PDF:

```sh
pdf
```

Estas medidas facilitan visualizar exactamente dónde se pasa la mayoría del tiempo de ejecución, identificando las líneas problemáticas en tu código que necesitan ser optimizadas.

El profiling es una herramienta clave para analizar y mejorar el rendimiento de tus aplicaciones. Ahora tienes el conocimiento para aplicarlo efectivamente. ¡Continúa aprendiendo y mejorando tu código con cada nueva técnica que explores!
