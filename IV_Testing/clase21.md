# Implementacion de morks en Go

## ¿Cómo implementar funciones simuladas para pruebas unitarias?

El arte de las pruebas unitarias nos permite verificar que nuestro código funcione correctamente y de manera eficiente, incluso cuando depende de otras funciones. Implementar funciones simuladas es una técnica poderosa que nos ayuda en este contexto. Aprenderemos a usar esta técnica de manera efectiva.

### ¿Qué son las funciones simuladas y por qué son útiles?

Las funciones simuladas (Mock Functions) son versiones ficticias de funciones reales que nos permiten probar unidades de código en aislamiento. Al usarlas, podemos:

- **Simular comportamientos complejos**: Sin depender de funciones que podrían ser costosas o difíciles de ejecutar en un entorno de prueba.
- **Asegurar que las funciones principales están aisladas**: Garantizando que la lógica principal se ejecute correctamente sin interferencias.
- **Facilitar la detección de errores**: Permitiendo cambios controlados en las respuestas de las funciones para verificar que se manejen adecuadamente los casos erróneos.

### ¿Cómo crear y utilizar funciones simuladas en Go?

#### Paso 1: Definición de estructuras y valores de prueba

Primero, definimos nuestras estructuras y casos de prueba:

```go
func main() {
    expectedEmployee := FullTimeEmployee{
        Person: Person{Age: 35, DNI: "1", Name: "Jungleu"},
        ID:     "1",
        Position: "CEO",
    }
}
```

Este bloque define un empleado con parámetros que utilizaremos en nuestras pruebas.

#### Paso 2: Almacenamiento y sustitución de funciones originales

Para modificar funciones adecuadamente sin perder sus versiones originales, las almacenamos:

```go
originalGetEmployeeByID := getEmployeeByID
originalGetPersonByID := getPersonByID
```

Luego, reemplazamos las funciones originales con las funciones simuladas.

#### Paso 3: Uso de rangos y comparación de resultados

Iteramos sobre nuestra tabla de casos de prueba y evaluamos:

```go
for _, test := range tests {
    ft, err := getFullTimeEmployeeByID(test.ID, test.DNI)
    if err != nil {
        log.Println("Error when getting employee:", err)
    }
    if ft.Person.Age != expectedEmployee.Person.Age {
        log.Printf("Expected age %d but got %d", expectedEmployee.Person.Age, ft.Person.Age)
    }
}
```

Es vital comparar propiedad por propiedad, siempre retornando las funciones originales al finalizar.

### ¿Cómo ejecutar y validar los resultados de las pruebas?

Utilizamos comandos específicos para ejecutar pruebas y verificar resultados. Con `go test`, podemos comprobar si nuestras pruebas pasan correctamente:

```sh
go test
```

Si unes estos conceptos con un enfoque meticuloso al almacenar y restaurar funciones, tu experiencia probando funciones complejas se enriquecerá notablemente.

### Reflexiones finales y mejores prácticas

- **Fijar valores originales**: Siempre mantén los valores originales de las funciones antes de hacer simulaciones.
- **Prueba cambios específicos**: Experimenta cambiando valores de prueba para garantizar que las condiciones bajo prueba se comporten como se espera.
- **Frecuencia en las pruebas**: Realiza pruebas de forma regular para captar errores de manera temprana y frecuente en el desarrollo.
