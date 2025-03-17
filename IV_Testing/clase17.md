# Testing

## ¿Qué es el testing y por qué es importante?

El testing es crucial en el desarrollo de software. Aunque es un tema amplio, aquí nos enfocaremos en las pruebas unitarias para nuestro código. Estas pruebas verifican el comportamiento de funciones individuales, asegurando que el código funcione correctamente y ayudando a detectar errores tempranamente. También mediremos el code coverage, que indica el porcentaje de código cubierto por las pruebas.

## ¿Cómo creamos pruebas unitarias en Go?

Para crear pruebas unitarias en Go, comenzamos con un archivo principal, `main.go`, que contiene una función simple de suma. Luego, creamos un archivo de pruebas, `main_test.go`, donde escribimos las pruebas para la función de suma. Este archivo debe seguir un formato específico para que Go lo reconozca, usando el sufijo `_test.go`.

### Configuración inicial

Primero, definimos el paquete principal y la función que vamos a probar:

```go
package main

func sum(x, y int) int {
    return x + y
}
```

### Creación de la prueba unitaria

En el archivo `main_test.go`, definimos el paquete y la función de prueba. Es esencial importar el paquete `testing`, ya que proporciona las herramientas necesarias para verificar los resultados de las pruebas.

```go
package main

import "testing"

func TestSum(t *testing.T) {
    total := sum(5, 5)
    if total != 10 {
        t.Errorf("Suma incorrecta, obtuvimos %d, pero esperábamos %d", total, 10)
    }
}
```

### Ejecutando las pruebas

Para ejecutar las pruebas, utilizamos el comando `go test` desde la terminal. Si configuramos correctamente nuestro módulo con `go mod init`, las pruebas deberían ejecutarse sin problemas. En caso de errores, Go reportará los fallos y sus ubicaciones específicas en el código.

## ¿Cómo utilizamos tablas de prueba para múltiples casos?

Un patrón común para aplicar pruebas en Go es el uso de tablas de prueba. Esta técnica nos permite definir múltiples escenarios de prueba en un slice de structs, facilitando la iteración y ejecución de cada caso.

### Creación de una tabla de pruebas

Definimos una estructura para nuestros casos de prueba, que incluya los valores de entrada y el resultado esperado.

```go
func TestSum(t *testing.T) {
    testCases := []struct{a, b, expected int}{
        {1, 2, 3},
        {2, 2, 4},
        {25, 26, 51},
    }

    for _, tc := range testCases {
        result := sum(tc.a, tc.b)
        if result != tc.expected {
            t.Errorf("Suma incorrecta, obtuvimos %d, pero esperábamos %d", result, tc.expected)
        }
    }
}
```

### Ventajas de las tablas de pruebas

Este enfoque organiza el proceso de pruebas y permite agregar y modificar casos fácilmente sin cambiar la lógica de prueba. Además, proporciona un marco claro para identificar fallos específicos cuando ocurren.

## ¿Qué impacto tiene el code coverage?

El code coverage es una métrica que indica qué parte de nuestro código ha sido ejecutada durante las pruebas. Aunque Go no obliga a cubrir el 100% del código, un bajo porcentaje de cobertura podría ocultar errores en partes no probadas. Go ofrece comandos para generar informes de cobertura, permitiéndonos identificar y mejorar áreas deficientes en nuestras pruebas.

### Integración del code coverage

Para calcular el code coverage, podemos usar el comando `go test -cover`, que mostrará un resumen del porcentaje de código cubierto por las pruebas. Esta práctica nos ayuda a mantener un estándar de calidad alto y a asegurar que todas las partes críticas del código estén debidamente verificadas.

El uso consciente de pruebas unitarias y el seguimiento del code coverage no solo mejoran la calidad del software, sino que también aumentan la confianza en nuestras aplicaciones. Invito a continuar explorando y aplicando estos conceptos en sus proyectos para asegurar un desarrollo más sólido y eficiente.