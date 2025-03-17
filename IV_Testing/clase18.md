# code coverage

- code coverage is a metric that measures the amount of code that is executed during the test suite

## ¿Qué es code coverage y por qué es importante?

El code coverage es una métrica crucial en el desarrollo de software que mide qué cantidad del código está cubierta por pruebas automáticas. Sirve para evaluar la calidad y efectividad de las pruebas implementadas. Cuanto más código esté probado, mayor confianza podemos tener en su desempeño y menor es el riesgo de errores en producción. A continuación, profundizaremos en cómo usar herramientas de coverage en Go para evaluar nuestras funciones y mejorar su calidad.

## ¿Cómo comprobar la cobertura de código en Go?

Para empezar, es esencial tener nuestra función completa, pero sin pruebas aún. En este caso, tenemos la función `getMax`, que necesita ser evaluada.

### Ejecutar pruebas con cobertura

Utilizamos el comando:

```sh
go test -cover
```

Esto nos indica qué porcentaje del código está siendo cubierto por las pruebas actuales. Por ejemplo, un resultado del 25% indica una cobertura baja y la necesidad de mejorarla.

### Generar un archivo de perfil de cobertura

Para identificar exactamente qué partes del código no han sido probadas, creamos un archivo de cobertura usando:

```sh
go test -coverprofile=coverage.out
```

Este archivo servirá para obtener detalles específicos de la cobertura.

## ¿Cómo visualizar un análisis legible de la cobertura?

Una vez que tenemos nuestro archivo de cobertura, debemos utilizar herramientas para interpretar la información adecuadamente.

### Análisis con Go tool cover

Usamos la herramienta `cover` para generar un resumen más legible:

```sh
go tool cover -func=coverage.out
```

Esto nos da un porcentaje claro y comprensible de la cobertura de cada archivo o función.

### Visualización en HTML

Para una comprensión más gráfica, podemos convertir nuestro archivo en un reporte HTML:

```sh
go tool cover -html=coverage.out
```

Esto abre una nueva pestaña en el navegador, mostrando el código probado en verde y el no probado en rojo. De este modo, podemos detectar cuáles funciones, como `getMax`, no están totalmente cubiertas por los tests.

## ¿Cómo mejorar la cobertura modificando pruebas?

Al identificar que nuestras pruebas no cubren todos los escenarios, se debe actuar para probar eficazmente.

### Modificar los casos de prueba

En la función `getMax`, inicialmente solo probamos casos donde el primer número era mayor que el segundo. Para mejorarlo:

```go
func TestMax(t *testing.T) {
    tables := []struct {
        a, b, n int
    }{
        {3, 2, 3},
        {2, 3, 3},
    }

    for _, item := range tables {
        max := getMax(item.a, item.b)
        if max != item.n {
            t.Errorf("getMax was incorrect, got: %d, want: %d.", max, item.n)
        }
    }
}
```

Esto incluye un caso adicional donde el segundo número es mayor.

### Regenerar el reporte de cobertura

Después de modificar los tests, generamos un nuevo reporte:

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Visualizaremos cómo ahora el 100% de la función `getMax` está cubierto si se han considerado todas las condiciones.
