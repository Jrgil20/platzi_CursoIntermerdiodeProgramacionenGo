# Herencia

## ¿Qué es la herencia en la programación orientada a objetos?

La herencia es un pilar fundamental en la programación orientada a objetos (POO) que permite a las clases obtener propiedades y comportamientos de otras clases. Esto facilita la reutilización del código y soporta el polimorfismo. Sin embargo, puede ser un tema controvertido debido a las diferentes implementaciones de jerarquías de clases en distintos lenguajes de programación, lo que puede afectar el código dependiente.

Mediante la herencia, se crea una clase base con propiedades comunes, de la cual otras clases pueden derivar. Estas clases derivadas heredan propiedades y comportamientos, reduciendo la duplicación de código.

## ¿Cómo se implementa la herencia con TypeScript?

En TypeScript, la herencia se implementa con la palabra clave `extends`. Esto permite que una clase derivada herede propiedades y métodos de una clase base:

```typescript
class BaseEmployee {
    // propiedades y métodos comunes
}

class TemporaryEmployee extends BaseEmployee {
    // propiedades y métodos específicos
}

class FullTimeEmployee extends BaseEmployee {
    // propiedades y métodos específicos
}
```

En este ejemplo, `TemporaryEmployee` y `FullTimeEmployee` heredan de `BaseEmployee`, lo que les permite acceder a sus métodos y propiedades públicas, evitando la duplicación de código.

## ¿Qué desafíos presenta la herencia?

Aunque útil, la herencia puede llevar a que las clases derivadas sobrescriban métodos de la clase base, lo que puede indicar un mal uso de la herencia y sugerir que la composición podría ser más adecuada.

Además, no todos los lenguajes manejan la herencia de la misma manera. Por ejemplo, en Go, no existe la herencia tradicional y se prefiere el principio de "composición sobre herencia". La composición permite que una clase tenga propiedades de otra sin ser necesariamente un tipo de esa clase.

## ¿Cómo se aplica la composición en Go?

En Go, la composición se logra mediante "campos anónimos". En lugar de heredar de una clase base, una estructura (struct) contiene campos que representan las propiedades de otras estructuras:

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    ID int
}

type FullTimeEmployee struct {
    Person
    Employee
}
```

En este ejemplo, `FullTimeEmployee` tiene las características de `Person` y `Employee` sin declarar explícitamente que es un `Person` o `Employee`, obteniendo los beneficios de ambas sin heredar.

## ¿Cómo afecta el polimorfismo?

El polimorfismo es otro concepto clave en POO. En TypeScript, es fácil implementar polimorfismo permitiendo que una función acepte una clase base como argumento, tratando cualquier clase derivada como esa clase base:

```typescript
function printMessage(employee: BaseEmployee) {
    console.log(employee.getMessage());
}
```

En Go, esto no es tan directo debido a la falta de herencia. Aunque la composición resuelve muchos problemas de diseño, limita el uso del polimorfismo tradicional. Sin embargo, Go utiliza interfaces para lograr un comportamiento similar, permitiendo que las estructuras compartan métodos sin necesidad de herencia directa.
