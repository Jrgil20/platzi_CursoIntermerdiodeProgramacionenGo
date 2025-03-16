# Curso de Go Intermedio: Programación Orientada a Objetos y Concurrencia

## [Slide](https://static.platzi.com/media/public/uploads/slides-del-curso-de-go-intermedio-programacion-orientada-a-objetos-y-concurrencia_80748e0d-0da8-41fe-ada4-ad5b6121863e.pdf)

### Introducción al curso

#### Golang

- Lenguaje compilado.
- Potente librería estándar.
- Manejo de concurrencia a través de GoRoutines y Channels.
- Diseñado por Google, Ken Thompson (UNIX) es parte del equipo de diseño.

- Muy popular en aplicaciones de CLI y Backend.
- Docker, Kubernetes y Terraform están escritos en Golang.
- Muchos proyectos de la CNCF están escritos en Go.
- Muy utilizado para escribir malware.

- Según StackOverflow, es el tercer lenguaje mejor pagado a nivel global.
- Según StackOverflow, es la quinta tecnología más amada por los desarrolladores y la tercera más deseada para trabajar.

### Cosas que debes saber

#### Conocimientos requeridos

- Declaración de variables
- Condicionales
- Sintaxis básica
- Declaración de GoRoutines y Channels
- Slices y maps
- Apuntadores

#### Qué cosas terminaré sabiendo

##### Conocimientos adquiridos

- ¿Es Go orientado a objetos?
- Cómo aplicar los conceptos de POO en Golang.
- Crear Test Unitarios en Go para aplicar TDD.
- Calcular el Code Coverage de mi proyecto.

- Análisis del Profiling en tus programas.
- Cómo multiplexar mensajes de canales.
- Técnicas para la creación de Worker Pools concurrentes.
- Crear Test Unitarios en Go para aplicar TDD.
- Crear un Job Queue concurrente.

### Repaso general: variables, condicionales, slices y map

#### Repaso general

- Repasar los conceptos importantes del curso anterior que serán cruciales para este curso.
- Definir structs y darles funcionalidad utilizando receiver functions.
- Apuntadores y cómo estos nos ayudan a utilizar/modificar las variables correctas y no generar copias.

- Go utiliza las GoRoutines como mecanismo para aplicar concurrencia.
- Channels como medios de comunicación entre diferentes rutinas.

### Repaso general: GoRoutines y apuntadores

### ¿Es Go orientado a objetos?

- POO se ha convertido en uno de los paradigmas de programación predominante en la industria.
- POO puede llegar a ser muy riguroso, pero a cambio permite una alta reutilización de código y la aplicación de un sinnúmero de patrones de diseño.

- Go puede alcanzar la aplicación de los conceptos de POO, pero de una forma diferente a lenguajes como Java o Python.

### Structs vs. clases

- El objetivo de una clase es definir una serie de propiedades y métodos que un objeto puede usar para alcanzar diferentes objetivos.
- Go utiliza Structs para generar “nuevos tipos” de datos que se pueden utilizar para cumplir tareas en específico.

### Métodos y funciones

- Algunos lenguajes de programación implementan la filosofía que TODO debe de ser un objeto, sin embargo, no siempre es algo aplicable, por ejemplo, en una librería con utilidades que no pertenecen a un dominio específico.

- Go permite que este tipo de “utilidades” no tengan que pertenecer obligatoriamente a un struct.

### Constructores

- Los constructores permiten la instanciación de una clase a un objeto, asimismo permite definir propiedades predefinidas.
- En Go podemos utilizar funciones que puedan crear structs con propiedades que nosotros pasamos como parámetros.

### Herencia

- Go no permite el concepto de herencia tal y como lo conocemos de lenguajes como TypeScript; en otros lenguajes de programación se utiliza la herencia como una forma de alcanzar reutilización de código y polimorfismo.

- Para alcanzar este mismo objetivo Go aplica un principio llamado Composition Over Inheritance que utilizando un campo anónimo en un struct puede “heredar” todas las propiedades y métodos.

### Interfaces

- Diferentes lenguajes de programación utilizan sintaxis explícita para decir que una clase implementa una interfaz.
- Go lo hace de manera implícita lo que permite la reutilización de código y el polimorfismo.

### Abstract factory

- Patrón de diseño de tipo creacional que permite la producción de objetos de la misma familia o tipo sin especificar su clase concreta, permitiendo esa determinación en tiempo de ejecución.

### Implementación final de Abstract Factory

- Go te permite implementar muchos de los patrones de diseño que están basados en polimorfismo mediante la utilización de interfaces.

### Funciones anónimas

- Múltiples lenguajes de programación permiten la creación de funciones que no tienen un nombre, JavaScript es quizá uno en dónde es más común ver este patrón.
- Go permite la creación de estas funciones anónimas, aunque deben ser usadas con cuidado para evitar romper el principio de DRY.

### Funciones variádicas y retornos con nombre

- Las funciones variádicas nos permiten utilizar como slices los argumentos de funciones de los cuales no sabemos su longitud exacta.
- Los retornos con nombre nos permitirán definir variables antes de definir el cuerpo de la función, por lo cual solo utilizaremos return para devolverlos.

### Go Modules

#### ¿Cómo los puedo utilizar?

- En Node podemos usar npm para generar un árbol de dependencias de nuestro proyectos, en Golang existen los modules con los cuales podremos definir en qué paquetes tenemos dependencias para poder instalarlas, asimismo utiliza el comando `go get -u package.name` para instalarlos.

#### Creando nuestro módulo

- Para generar un módulo nuevo, utilizaremos `go mod init my.module`, a partir de aquí seremos capaces de generar un módulo que otras personas puedan utilizar.

### Testing

- Por lo general, muchos lenguajes de programación utilizan dependencias externas para la creación y ejecución de tests.
- Golang, haciendo realce de su magnífica librería estándar, nos permite crear los tests que necesitamos para tener un código más robusto.

### Code coverage

- El code coverage es una métrica que nos permitirá identificar las partes de nuestro código que no han sido probadas y que potencialmente tienen el riesgo de contener un bug.

### Profiling

- El profiling nos ayudará a encontrar en nuestro código aquellas partes críticas que influyen en una ejecución lenta o con mucho consumo de memoria. Con estas técnicas sabremos en qué enfocarnos a la hora de mejorar nuestro código.

### Testing usando mocks

- Cuando escribimos un test unitario que tiene dependencias en diferentes servicios nos vemos en la necesidad de crear “mocks” que emulen comportamientos de esos servicios con el objetivo que nuestro test se encargue de probar la funcionalidad que nos interesa y no la de las dependencias.

### Implementando Mocks

- Es muy importante guardar los valores originales de las funciones que estamos utilizando para realizar nuestros tests en caso de que en test posteriores decidamos evaluar otras partes del código que también dependen de esas funciones.

### Unbuffered channels y buffered channels

- Go utiliza dos tipos de canales para las comunicaciones, los unbuffered channels son aquellos que tienen una cola de longitud 0, es decir, se utilizan para comunicación síncrona.

- También existen Buffered Channels que permiten definir una longitud de buffer, es decir, se pueden almacenar ciertos valores antes de empezar a leer.

### Waitgroup

- Cuando se trabaja con GoRoutines no siempre se planea que estas envíen datos a través de canales entre unas y otras, en estos casos se puede utilizar
