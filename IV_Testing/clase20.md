# MORKS

## ¿Qué es el testing con mocks y su importancia?

El testing con mocks es una técnica crucial en el desarrollo de software que permite probar unidades de código de manera aislada, sin depender de sus dependencias externas. Por ejemplo, si una función consulta una base de datos, en lugar de hacer llamadas reales durante el test, se utilizan mocks para simular ese comportamiento. Esto permite probar la lógica de la función sin influencias externas.

## ¿Cómo se implementa un mock en Go?

Para implementar un mock en Go, se comienza creando funciones que realizan acciones específicas, como obtener datos de una base de datos. Luego, estas funciones se redefinen como variables para poder asignarles diferentes comportamientos durante el test.

### Ejemplo básico

#### Definición de funciones originales

```go
func getPersonByDNI(dni string) (Person, error) {
    // Simula lectura de base de datos
    return Person{}, nil
}

func getEmployeeByID(id int) (Employee, error) {
    // Simula lectura de base de datos
    return Employee{}, nil
}
```

#### Reasignación de funciones como variables

```go
var getPersonByDNI = func(dni string) (Person, error) {
    return Person{}, nil
}

var getEmployeeByID = func(id int) (Employee, error) {
    return Employee{}, nil
}
```

## ¿Cómo crear y utilizar un mock en tests?

Una vez que las funciones están definidas como variables, se pueden redefinir dentro del test para simular el comportamiento deseado.

### Ejemplo de test con mocks

```go
func TestGetFullTimeEmployeeByID(t *testing.T) {
    // Definimos la función mock
    mockGetEmployeeByID := func(id int) (Employee, error) {
        return Employee{ID: 1, Position: "CEO"}, nil
    }

    mockGetPersonByDNI := func(dni string) (Person, error) {
        return Person{Name: "John Doe", Age: 35, DNI: "1"}, nil
    }

    // Asignamos los mocks a nuestras variables de función
    getEmployeeByID = mockGetEmployeeByID
    getPersonByDNI = mockGetPersonByDNI

    // Aquí se procede con las pruebas utilizando los mocks
    // ...
}
```

Este enfoque facilita aislar la lógica que se desea probar, evitando que las pruebas dependan de otros componentes del sistema como bases de datos o servicios externos.

## ¿Por qué es importante la modularidad y la reutilización con mocks?

Utilizando mocks, los desarrolladores pueden:

- **Reducir dependencias externas:** Las pruebas no se ven afectadas por servicios externos o estados de base de datos.
- **Acelerar tiempos de prueba:** Las ejecuciones son
