# Constructores

## ¿Qué son los constructores y por qué son importantes?

Los constructores son esenciales en la programación orientada a objetos, ya que permiten crear instancias de clases y convertirlas en objetos. Su función principal es preconfigurar o precargar propiedades necesarias para el funcionamiento y personalización de estos objetos. Independientemente del lenguaje, ya sea TypeScript, Python o Go, los constructores tienen el mismo objetivo: asignar valores iniciales a las propiedades de un objeto al momento de su creación.

## ¿Cómo se implementan los constructores en diferentes lenguajes de programación?

### En TypeScript

En TypeScript, los constructores son una parte integral de las clases. Permiten definir y asignar valores a propiedades específicas al crear un objeto. Por ejemplo, un constructor que recibe parámetros como `name` e `id` los asignará y utilizará al instanciar un nuevo objeto de la clase:

```typescript
class Employee {
    constructor(public name: string, public id: number) {}
}

const employee = new Employee('John', 1);
```

### En Python

Python utiliza el método especial `__init__` para inicializar objetos con las propiedades deseadas:

```python
class Employee:
        def __init__(self, name, id):
                self.name = name
                self.id = id

employee = Employee('John', 1)
```

### En Go

En Go, los constructores no existen como tal, pero se pueden emular mediante funciones personalizadas. Utilizando `structs`, se puede definir una estructura y una función para crear instancias con valores iniciales:

```go
type Employee struct {
        ID       int
        Name     string
        Vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
        return &Employee{
                ID:       id,
                Name:     name,
                Vacation: vacation,
        }
}

func main() {
        e1 := NewEmployee(1, "Alice", true)
        fmt.Println(e1)
}
```

## ¿Cómo agregar propiedades con valores predeterminados en Go?

En Go, si no se especifican valores al inicializar un `struct`, el lenguaje asigna valores por defecto: `0` para enteros, `""` para strings y `false` para booleanos. Esto es crucial para evitar errores al suponer que las propiedades son `nil` sin inicialización explícita.

```go
// Ejemplo: struct con valores predeterminados
employee := Employee{}
fmt.Println(employee)  // Output: Employee{ID:0 Name:"" Vacation:false}
```

## ¿Cuáles son las diferentes maneras para crear objetos en Go?

### Fórmula número uno: Inicialización directa

Podemos inicializar un objeto directamente sin especificar valores para las propiedades, lo cual asignará valores predeterminados.

### Fórmula número dos: Inicialización explícita

Definiendo explícitamente las propiedades desde su creación, se asignarán los valores especificados:

```go
employee := Employee{
        ID:       1,
        Name:     "Bob",
        Vacation: true,
}
```

### Fórmula número tres: Uso del operador `new`

El operador `new` es una alternativa válida, pero devuelve un apuntador en vez de una copia del objeto:

```go
employee := new(Employee)
```

### Fórmula número cuatro: Uso de una función constructora

Crear una función constructora personalizada ofrece flexibilidad y control detallado sobre el proceso de inicialización, permitiendo ajustar el comportamiento del objeto según se necesite antes de su creación definitiva.

```go
func NewEmployee(id int, name string, vacation bool) *Employee {
        return &Employee{
                ID:       id,
                Name:     name,
                Vacation: vacation,
        }
}
```

## ¿Por qué optar por funciones constructoras en Go?

El uso de funciones constructoras es especialmente valioso en Go, ya que permite definir valores iniciales y ejecutar lógica adicional antes de retornar el objeto. Esto incluye la inicialización de servicios, carga de dependencias o cualquier configuración necesaria. Se recomienda el uso de funciones constructoras para maximizar la personalización y eficiencia del proceso de creación de objetos en Go, emulando el comportamiento de los constructores en otros lenguajes de programación orientada a objetos.
