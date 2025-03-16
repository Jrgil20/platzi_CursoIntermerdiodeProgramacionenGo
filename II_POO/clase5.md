# Es Go orientada a objetos?

## ¿Go es un lenguaje orientado a objetos?

El debate sobre si Go es un lenguaje orientado a objetos ha existido desde su creación, con opiniones divididas. Este paradigma es predominante en el desarrollo de software moderno, basado en la idea de que casi todo puede ser visto como un objeto, como en lenguajes como Java. En este curso, asumiremos que Go incorpora elementos de la programación orientada a objetos. Sin embargo, la experiencia de cada desarrollador puede ofrecer diferentes perspectivas. Este artículo explorará cómo Go puede ser utilizado con un enfoque orientado a objetos, aunque con algunas diferencias.

### ¿Qué es la programación orientada a objetos?

La programación orientada a objetos (POO) es un paradigma que organiza el código y representa datos mediante estructuras llamadas objetos. Los objetos se definen a través de clases, que actúan como plantillas. Estas clases contienen propiedades y métodos que proporcionan una estructura clara y reutilizable para crear múltiples instancias de un objeto.

### ¿Cómo funcionan las clases en TypeScript?

Las clases en TypeScript son un ejemplo claro de la implementación de la POO en lenguajes modernos. Actúan como plantillas que definen propiedades y métodos. Por ejemplo, considera una clase `Empleado` que puede tener propiedades como `nombre` e `identificador`, así como métodos para interactuar con esas propiedades.

```typescript
class Empleado {
    constructor(public id: number, public nombre: string) {}

    asignarID(nuevoID: number) {
        this.id = nuevoID;
    }

    obtenerNombre() {
        return this.nombre;
    }
}

let empleado1 = new Empleado(1, "Juan");
empleado1.asignarID(2);
console.log(empleado1.obtenerNombre());
```

Aquí se muestra cómo se instancian objetos, es decir, cómo se crean entidades concretas a partir de la plantilla proporcionada por la clase.

### ¿Cómo Go implementa conceptos de programación orientada a objetos?

Go es conocido por su enfoque pragmático y eficiente en el desarrollo de software. Aunque carece de algunas características tradicionales de la POO, como clases y herencia, ofrece tipos y métodos que permiten emular este paradigma. En lugar de clases, Go utiliza estructuras (`structs`) para definir tipos personalizados con campos similares a las propiedades.

### ¿Qué son las estructuras en Go?

Las estructuras en Go son colecciones de campos que pueden asociar comportamiento mediante la definición de métodos. Permiten encapsular datos y comportamiento de manera similar a las clases en otros lenguajes orientados a objetos.

```go
type Empleado struct {
    ID   int
    Nombre string
}

func (e *Empleado) AsignarID(nuevoID int) {
    e.ID = nuevoID;
}

func (e *Empleado) ObtenerNombre() string {
    return e.Nombre;
}

func main() {
    empleado1 := Empleado{ID: 1, Nombre: "Juan"}
    empleado1.AsignarID(2)
    fmt.Println(empleado1.ObtenerNombre())
}
```

Esta estructura y los métodos asociados demuestran cómo Go puede simular la configuración y comportamiento de un objeto.
