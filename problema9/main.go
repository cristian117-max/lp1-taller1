package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar una versión del problema de los Filósofos Comensales.
// Hay 5 filósofos y 5 tenedores (recursos). Cada filósofo necesita 2 tenedores para comer.
// Estrategia segura: imponer un **orden global** al tomar los tenedores (primero el menor ID, luego el mayor)
// para evitar deadlock. También puedes limitar concurrencia (ej. mayordomo).
//  completa la lógica de toma/soltado de tenedores y bucle de pensar/comer.

type tenedor struct{ mu sync.Mutex }

func filosofo(id int, izq, der *tenedor, wg *sync.WaitGroup) {
	//  desarrolla el código para el filósofo
	defer wg.Done()

	for i := 0; i < 3; i++ {
		pensar(id)

		if id < (id+1)%5 {
			izq.mu.Lock()
			der.mu.Lock()
		} else {
			der.mu.Lock()
			izq.mu.Lock()
		}
		comer(id)
		izq.mu.Unlock()
		der.mu.Unlock()
	}
	fmt.Printf("[filósofo %d] satisfecho\n", id)
}

func pensar(id int) {
	fmt.Printf("[filósofo %d] pensando...\n", id)
	//  simular tiempo de pensar
	time.Sleep(200 * time.Millisecond)

}

func comer(id int) {
	fmt.Printf("[filósofo %d] COMIENDO\n", id)
	// simular tiempo de pensar
	time.Sleep(200 * time.Millisecond)

}

func main() {
	const n = 5
	var wg sync.WaitGroup
	wg.Add(n)

	// crear tenedores
	forks := make([]*tenedor, n)
	for i := 0; i < n; i++ {
		// inicializar cada tenedor i
		forks[i] = &tenedor{}

	}

	// lanzar filósofos
	for i := 0; i < n; i++ {
		izq := forks[i]
		der := forks[(i+1)%n]
		//  lanzar goroutine para el filósofo i
		go filosofo(i, izq, der, &wg)

	}

	wg.Wait()
	fmt.Println("Todos los filósofos han comido sin deadlock.")
}
