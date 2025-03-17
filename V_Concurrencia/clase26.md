# Woker Pools

## ¿Qué es Workerpool en Go y cómo puede transformar tus proyectos?

El ámbito de la programación es extenso y está lleno de conceptos interesantes. Uno de los más poderosos en Go es el workerpool, que permite la ejecución concurrente de tareas específicas por múltiples "trabajadores" o workers. Es crucial entender cómo implementar esto en tus proyectos para optimizar tareas que consumen tiempo. ¡Veamos cómo funciona en Go!

### ¿Cómo se inicia un proyecto de workerpool en Go?

Primero, crea un archivo en Go, por ejemplo, wp.go. Este archivo comenzará como la mayoría de los programas en Go, definiendo el paquete con `package main`. Luego, define una función para calcular la serie de Fibonacci. Esta función, que utiliza una implementación recursiva, será la pieza central del trabajo que realizarán los workers.

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

### ¿Cómo se crea un worker en Go?

Un worker es una función que recibe tareas (jobs) a través de un canal, las procesa y envía el resultado a través de otro canal. Es importante asignar un ID a cada worker para identificar qué trabajador realiza cada tarea.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, job)
        results <- fibonacci(job)
        fmt.Printf("Worker %d finished job %d\n", id, job)
    }
}
```

### ¿Cómo implementar el workerpool en la función principal?

La función `main` coordina todo el proceso: definir las tareas, crear los canales de comunicación, iniciar los workers y manejar los resultados.

- **Definir las tareas:** Utiliza un slice para almacenar los números a los que se les aplicará la serie de Fibonacci.
- **Configurar los workers:** Crea varios workers, asignándoles un ID y pasando los canales de tareas y resultados.
- **Manejar los canales:** Asigna las tareas a los workers a través de un canal y recibe los resultados por otro.

```go
func main() {
    tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
    numOfWorkers := 3

    jobs := make(chan int, len(tasks))
    results := make(chan int, len(tasks))

    for w := 1; w <= numOfWorkers; w++ {
        go worker(w, jobs, results)
    }

    for _, task := range tasks {
        jobs <- task
    }
    close(jobs)

    for i := 0; i < len(tasks); i++ {
        <-results
    }
}
```

### ¿Cómo optimiza el uso de gorutinas y canales en Go?

La magia de Go reside en su capacidad para manejar la concurrencia de manera eficiente y sencilla. Al utilizar gorutinas y canales de comunicación, es posible crear programas altamente concurrentes sin la complejidad que otros lenguajes requieren. Cada worker es manejado como una gorutina, procesando tareas de manera concurrente y colaborando mediante canales de comunicación.

Esta estructura no solo mejora el tiempo de ejecución de tareas pesadas, como el cálculo de la serie de Fibonacci para grandes números, sino que también garantiza que los recursos del sistema se usen de manera eficiente.

Con esta implementación, aumentamos nuestra capacidad de procesar tareas concurrentemente, algo vital para aplicaciones que requieren manejar cientos o miles de acciones simultáneamente.
