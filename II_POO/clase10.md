# interfaces

## ¿Cómo funcionan las interfaces en la programación orientada a objetos?

Las interfaces son fundamentales en la programación orientada a objetos, ya que facilitan el polimorfismo y varios patrones de diseño. Actúan como "contratos" que obligan a las clases que las implementan a definir todos los métodos especificados. Su importancia radica en la simplificación del código y la gestión eficiente del comportamiento de distintas clases.

## ¿Qué es una interfaz en programación?

Una interfaz es un conjunto de métodos que una clase debe implementar para cumplir con un contrato específico. Esto implica:

- **Definición de contrato**: Una interfaz garantiza que cualquier clase que la implemente tendrá los métodos especificados.
- **Flexibilidad en la implementación**: Aunque una interfaz define qué métodos deben existir, no especifica cómo deben ser implementados.

Por ejemplo, en TypeScript, una interfaz llamada `printInfo` puede requerir una función `getMessage` que devuelva un string. Cualquier clase que implemente esta interfaz debe definir el método `getMessage`.

## ¿Cómo se implementan las interfaces en TypeScript y Go?

### Uso de interfaces en TypeScript

- **Declaración**: Se utiliza la palabra clave `interface` para definir los métodos requeridos.
- **Implementación**: Las clases usan `implements` para adoptar la interfaz.

Un ejemplo en TypeScript muestra cómo una clase puede implementar varias interfaces, manteniendo su lógica interna única. Esto permite a los desarrolladores implementar la funcionalidad necesaria de manera específica a la estructura de su aplicación, apoyando el polimorfismo.

### Implementación implícita en Go

A diferencia de TypeScript, Go utiliza un enfoque implícito para las interfaces:

- **Declaración**: Similar a otros lenguajes, Go permite definir un tipo de interfaz. Las clases que coinciden con la interfaz la implementan automáticamente.
- **Polimorfismo**: A través de métodos comunes en diferentes estructuras (`struct`), Go permite operaciones polimórficas utilizando interfaces.

El siguiente código muestra la creación de una interfaz `printInfo` en Go:

```go
type printInfo interface {
    getMessage() string
}
```

Dos structs, `TemporaryEmployee` y `FullTimeEmployee`, pueden implementar esta interfaz implícitamente si definen el método `getMessage` correctamente.

## ¿Por qué son importantes las interfaces?

El uso de interfaces promueve un diseño más limpio y modular, y reduce la duplicación de código. Con interfaces, no es necesario definir múltiples funciones `printMessage` para cada clase, sino que se puede usar un solo método genérico que acepte cualquier objeto que implemente la interfaz. Esto:

- **Reduce la repetición**: Evita la creación de funciones duplicadas para cada clase.
- **Facilita la extensión**: Al añadir nuevas clases, solo se necesita implementar la interfaz sin cambiar el código existente.
- **Fomenta el polimorfismo**: Permite tratar diferentes clases a través de una misma interfaz, expandiendo las posibilidades de reutilización y adaptación del código.
