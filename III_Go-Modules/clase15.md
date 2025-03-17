# Go modulos

## ¿Cómo gestionar dependencias en Go?

Comprender la gestión de dependencias en Go es crucial para integrar módulos de otros proyectos de manera eficiente. Similar a cómo se usan herramientas como NPM en Node.js o PIP en Python, Go tiene su propio sistema para facilitar este proceso. ¡Vamos a explorarlo!

### ¿Cómo crear un módulo en Go?

El proceso es simple y se realiza desde la terminal:

1. Navega a la carpeta donde deseas crear el módulo.
2. Ejecuta el siguiente comando para inicializar un nuevo módulo:

```sh
go mod init github.com/tu-usuario/nombre-modulo
```

Este comando no solo crea el módulo, sino que también prepara el entorno para futuras dependencias.

### ¿Cómo instalar una dependencia?

Una vez que tu módulo esté listo, instalar dependencias externas es sencillo:

Usa el comando `go get` seguido del repositorio y módulo que necesitas:

```sh
go get github.com/usuario/modulo
```

Esto actualizará automáticamente el archivo `go.mod` con las dependencias utilizadas y generará un archivo `go.sum` para verificar la integridad del código descargado.

### ¿Qué hacer si no usas una dependencia?

Para eliminar dependencias no utilizadas, utiliza el siguiente comando:

```sh
go mod tidy
```

Este comando limpia las referencias innecesarias, asegurando que solo mantengas lo esencial para tu proyecto.

### ¿Cómo usar caché para mejorar la gestión de dependencias?

Cuando descargas módulos, Go los almacena en un caché para optimizar futuras instalaciones. Para visualizar este caché, usa:

```sh
go mod download -json
```

Esto mostrará en detalle dónde y cómo se almacenan localmente tus dependencias.

### ¿Cómo utilizar las dependencias en tu proyecto?

Para usar un módulo instalado, importa el paquete en tu archivo `main.go` y llama a las funciones necesarias. Por ejemplo, si ya instalaste `hello-mode`, podrías hacerlo así:

```go
package main

import (
    "github.com/usuario/hello-mode"
)

func main() {
    hello-mode.SayHello("Platzi")
}
```

La herramienta de autocompletar de tu editor te ayudará sugiriendo funciones de los módulos importados.

### ¿Cómo manejar múltiples versiones de un módulo?

Si necesitas diferentes versiones de un mismo módulo, puedes hacerlo de la siguiente manera:

1. Instala la versión específica:

    ```sh
    go get github.com/usuario/modulo/v2
    ```

2. Usa alias en tu importación para diferenciarlas:

    ```go
    import (
        moduloV1 "github.com/usuario/modulo/v1"
        moduloV2 "github.com/usuario/modulo/v2"
    )
    ```

Cuando invoques sus funciones, usa estos alias para especificar qué versión estás utilizando.

### Consejos prácticos para manejar dependencias

- **Verificar la documentación:** Antes de actualizar o instalar un nuevo módulo, revisa la documentación para prever cambios que podrían afectar tu código.
- **Compatibilidad de versiones:** Ten cuidado al usar nuevas versiones, ya que podrían incluir cambios incompatibles. Usar alias puede ayudarte a manejar distintas versiones simultáneamente.
- **Uso del archivo `go.sum`:** Este archivo comprueba la integridad del código descargado. Si ves discrepancias en los checksums, revisa el código para asegurarte de que no hay cambios sospechosos.
