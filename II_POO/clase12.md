# Abstact factory 2

## ¿Cómo implementar el patrón de diseño de notificación para email y SMS?

En el desarrollo de software, los patrones de diseño son cruciales para crear código reutilizable y escalable. Un ejemplo es el patrón de diseño de notificaciones, que permite enviar mensajes de manera abstracta y concreta. Este enfoque es útil para integrar diferentes métodos de envío, como SMS y Email, en una única arquitectura.

### ¿Cómo crear un tipo de notificación para email?

Primero, crea un tipo de notificación para email similar al SMS. Aquí, un tipo 'struct' de notificación por email será la base. Es importante que todas las clases concretas sigan una interfaz común, la cual no define una forma específica en el código.

```go
type EmailNotification struct{}

func (en EmailNotification) SendNotification() {
    fmt.Println("Sending Notification via Email")
}
```

### ¿Cuál es la función del 'iSender' en las notificaciones?

El siguiente paso es establecer un 'iSender' específico para el email, creando un nuevo 'struct' llamado EmailNotificationSender. Este debe implementar ciertos métodos para integrar sus funcionalidades adecuadamente.

```go
type EmailNotificationSender struct{}

func (ens EmailNotificationSender) GetSenderMethod() string {
    return "Email"
}

func (ens EmailNotificationSender) GetSenderChannel() string {
    return "SES" // Asumiendo que usamos Amazon Simple Email Service
}
```

### ¿Cómo se implementa la interfaz INotificationFactory para gestionar errores?

La interfaz INotificationFactory se utiliza para retornar la clase concreta que se encargará de enviar notificaciones. Implementar la gestión de errores en esta interfaz asegura que manejemos adecuadamente los tipos incorrectos de notificaciones.

```go
func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
    if notificationType == "SMS" {
        return SMSNotification{}, nil
    } else if notificationType == "Email" {
        return EmailNotification{}, nil
    }
    return nil, fmt.Errorf("No Notification Type")
}
```

### ¿Cómo usar el patrón de diseño en el código sin conocer detalles internos?

Para finalizar la implementación, necesitarás dos funciones: sendNotification y getMethod. Estas recibirán como parámetro una interfaz INotificationFactory, permitiendo que el código no necesite saber qué tipo de notificación específica se está utilizando. A continuación, la función main donde se ejecuta el patrón de diseño.

```go
func main() {
    smsFactory, _ := GetNotificationFactory("SMS")
    emailFactory, _ := GetNotificationFactory("Email")
    
    SendNotification(smsFactory)
    SendNotification(emailFactory)
}

func SendNotification(nf INotificationFactory) {
    nf.SendNotification()
}

func GetMethod(nf INotificationFactory) {
    fmt.Println(nf.GetSender().GetSenderMethod())
}
```

Implementar el patrón de diseño de notificaciones utilizando Go y interfaces permite la reutilización del código y facilita futuras ampliaciones de funcionalidades. Esta estructura modular es fundamental para mantener la calidad de las aplicaciones en crecimiento. Te animo a seguir explorando patrones de diseño para mejorar tu habilidad en programación.
