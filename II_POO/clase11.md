# Abstact factory

## ¿Qué es un patrón de diseño y cómo se relaciona con las interfaces?

Imagina que tienes un proyecto de software complejo; un plano arquitectónico es esencial para diseñar su estructura. De manera similar, en la programación, los patrones de diseño son soluciones probadas que te ayudan a resolver problemas comunes. Son especialmente importantes en situaciones donde el polimorfismo y las interfaces juegan un papel crucial.

## ¿Cómo se aplican los patrones de diseño en el caso de las notificaciones?

El patrón de diseño Abstract Factory es útil para crear familias de objetos similares. Por ejemplo, si desarrollas un software que maneja múltiples tipos de notificaciones como SMS y email, el reto es diseñar un programa que gestione estos diferentes tipos de notificaciones de forma polimórfica. Veamos cómo se logra paso a paso.

## Implementación del patrón Abstract Factory en Go

### ¿Cómo se definen las interfaces en Go para este patrón?

En Go, una interfaz define un conjunto de funciones que cualquier tipo debe implementar para cumplir con esa interfaz. Aquí creamos interfaces para manejar notificaciones:

```go
package main

type iNotificationFactory interface {
    sendNotification()
    getSender() iSender
}

type iSender interface {
    getSenderMethod() string
    getSenderChannel() string
}
```

Estas interfaces son el esqueleto que cada tipo de notificación necesitará implementar.

### ¿Cómo se implementa una notificación SMS?

Una notificación SMS es una estructura específica que no solo implementa las interfaces, sino que proporciona una definición concreta.

```go
type SMSNotification struct{}

func (SMSNotification) sendNotification() {
    fmt.Println("Enviando notificación vía SMS")
}

type SMSNotificationSender struct{}

func (SMSNotificationSender) getSenderMethod() string {
    return "SMS"
}

func (SMSNotificationSender) getSenderChannel() string {
    return "Twilio"
}

func (SMSNotification) getSender() iSender {
    return &SMSNotificationSender{}
}
```

Aquí, hemos creado un método que envía la notificación y un sender que define los detalles técnicos del envío.

## Crea y Extiende Tu Arquitectura

La magia del patrón Abstract Factory reside en su capacidad de ser extensible. Si deseas agregar un nuevo tipo de notificación, como Email, simplemente implementas nuevas estructuras y métodos que cumplan las interfaces ya definidas.

## ¿Por qué es útil entender y aplicar Abstract Factory?

- **Flexibilidad**: Aumenta la capacidad de tu software de adaptarse y ampliar sus funcionalidades.
- **Simplicidad y Organización**: Organiza el código a través de interfaces claras y responsabilidades definidas.
- **Reutilización de Código**: Al diseñar las interfaces, permitimos que implementaciones futuras se basen sobre este marco.

Estudiar y aplicar patrones de diseño como Abstract Factory te prepara para enfrentar desafíos en el desarrollo de software de manera creativa y efectiva. ¡Sigue explorando para convertir nubes de incertidumbre en claros caminos de soluciones!
