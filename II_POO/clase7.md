# Metodos y funciones

## ¿Cómo implementar métodos en Go usando receiver functions?

Al explorar Go, encontramos una característica clave que simula el comportamiento de clases y métodos en otros lenguajes: las receiver functions. Aunque al principio puede parecer inusual, su comprensión revela su gran potencial. A continuación, veremos cómo usar métodos set y get con receiver functions para structs en Go.

### ¿Qué son y cómo funcionan las receiver functions en Go?

Las receiver functions son funciones que permiten asociar métodos a un struct, emulando la funcionalidad de clases en otros lenguajes. Esto se logra declarando una función con un parámetro receptor que apunta a la instancia del struct donde se aplicará el método. Así se crean métodos en Go:

```go
func (e *Employee) SetID(id int) {
    e.ID = id
}
```

### ¿Cómo definir métodos set y get usando receiver functions?

#### Método setID

Para establecer un método que modifique una propiedad, como el ID de un empleado, se especifica una receiver function. Aquí se muestra cómo implementar el método setID:

```go
func (e *Employee) SetID(id int) {
    e.ID = id
}
```

- **Parámetro receptor:** `(e *Employee)` - Permite que el método acceda a la instancia actual del struct Employee.
- **Modificar propiedad:** Utiliza `e.ID` para cambiar el valor del ID.

#### Método setName

Para modificar el nombre de un empleado, se sigue un enfoque similar:

```go
func (e *Employee) SetName(name string) {
    e.Name = name
}
```

- **Nuevo valor:** El método recibe un nuevo valor para `Name` y lo asigna a la propiedad del struct.

### Métodos getID y getName

Además de modificar, también es necesario recuperar información. Aquí se demuestra cómo crear métodos get:

```go
func (e *Employee) GetID() int {
    return e.ID
}
func (e *Employee) GetName() string {
    return e.Name
}
```

### Ejemplo de uso de estos métodos

Aquí hay un pequeño ejemplo para ilustrar cómo estos métodos pueden ser utilizados juntos:

```go
employee := Employee{}
employee.SetID(5)
employee.SetName("John Doe")

fmt.Println("ID:", employee.GetID()) // Imprime "ID: 5"
fmt.Println("Name:", employee.GetName()) // Imprime "Name: John Doe"
```

### La potencia del enfoque de Go

Este enfoque puede parecer diferente a la programación orientada a objetos tradicional, pero ofrece una poderosa flexibilidad y mantiene la simplicidad de Go. Sin embargo, surgen preguntas comunes:

- **Constructores:** Aunque Go no tiene constructores tradicionales, permite inicializar structs de forma eficaz.
- **Herencia e interfaces:** Go no sigue una jerarquía de clase pero ofrece interfaces para la reutilización de código y comportamientos genéricos.
