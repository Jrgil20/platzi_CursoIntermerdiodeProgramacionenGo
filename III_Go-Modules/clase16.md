# my module

## ¿Cómo crear y compartir un módulo en GitHub con Go?

Crear y compartir un módulo en Go a través de GitHub es una habilidad valiosa para cualquier desarrollador. Este proceso no solo mejora tus proyectos personales, sino que también te permite contribuir al mundo del open source, poniendo tus soluciones a disposición de otros programadores y formando parte de una comunidad global.

### Primer paso: crear un repositorio en GitHub

El primer paso es crear un repositorio en GitHub, que servirá como el lugar donde almacenarás tu proyecto de Go. Este repositorio es esencial para poder compartir tu módulo con otros y para gestionar versiones del mismo.

1. **Crea un repositorio en GitHub**: Recuerda la ruta y el nombre del repositorio, así como tu nombre de usuario. Esta información será crucial para configurar tu módulo más adelante.
2. **Configura Git en tu computadora**: Asegúrate de tener Git instalado y configurado para poder interactuar con GitHub desde tu computadora.

### Crear un módulo en Go

Una vez que tienes tu repositorio preparado, puedes proceder a crear el módulo en Go utilizando Visual Studio Code como el entorno de desarrollo.

1. **Inicializa el módulo**: Usa el comando `go mod init` seguido por la ruta de tu repositorio en GitHub para inicializar un nuevo módulo de Go. Por ejemplo:

    ```sh
    go mod init github.com/tu_usuario/tu_repositorio
    ```

2. **Crea el archivo de utilidades**: Crea un archivo llamado `utils.go` donde definirás las funciones de tu módulo.

3. **Define el paquete y las funciones**: El paquete no debe ser `main`, ya que este módulo servirá como una dependencia externa. Asegúrate de que las funciones que quieras exportar empiecen con una letra mayúscula.

    ```go
    package utiles

    import "fmt"

    func HelloWorld() {
        fmt.Println("Hello from Platzi")
    }
    ```

### Subir el módulo a GitHub

Una vez que tu módulo está listo, es momento de subirlo a GitHub para que esté disponible para su uso.

1. **Inicializa el repositorio local**: Usa `git init` para inicializar el repositorio.
2. **Agrega el repositorio remoto**: Usa `git remote add origin` seguido por la URL de tu repositorio en GitHub.
3. **Crea una rama principal**: Usa `git branch -M main` para establecer la rama principal.
4. **Agrega y comitea los archivos**: Usa los comandos `git add .` y `git commit -m "Initial commit"` para preparar y confirmar tus cambios.
5. **Haz push a GitHub**: Finalmente, usa `git push -u origin main` para subir tus archivos al repositorio.

### Utilizar el módulo en otro proyecto

Ahora que tu módulo está en GitHub, puedes utilizarlo en cualquier proyecto nuevo.

1. **Crea un nuevo módulo**: Dentro de un directorio nuevo, usa el comando `go mod init` para inicializar el módulo.
2. **Agrega la dependencia**: Usa `go get github.com/tu_usuario/tu_repositorio` para agregar tu módulo como una dependencia.
3. **Importa y utiliza el módulo en tu nuevo proyecto**:

    ```go
    package main

    import (
        "github.com/tu_usuario/tu_repositorio/utiles"
    )

    func main() {
        utiles.HelloWorld()
    }
    ```

Seguir estos pasos no solo optimiza tu flujo de trabajo, sino que también fomenta la colaboración y la innovación en proyectos de código abierto. ¡No olvides experimentar y compartir tus propias creaciones con el mundo!
