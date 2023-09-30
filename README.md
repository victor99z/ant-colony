# Ant Colony

### Instalação e Requisitos

- golang 1.21

### Executar

Se se o parametro `debug` for passado na inicilização, o programa ira inicializar a engine de
jogos 2d chamada `ebiten` para exibir o comportamento da formiga ao longo do
tempo de iteração estipulado no arquivo /utils/constraints.go

Caso parametro `print` for recebido, ira apenas mostrar o resultado final das iterações.

```bash
$ go mod tidy
$ go run main.go debug
$ go run main.go print
```

### Verificar as race condition das formigas

```bash
$ go run -race main.go
```

### Benchmark??

- Dar uma olhada no package `testing`

```
$ go test -bench=.
```

### TODO:

- [x] Formigas paralelas
- [x] Debug usando alguma engine 2D simples (ebiten)

### FIXME:

- [x] Quando termina a iteração, se a formiga estiver com o item ele sera
      perdido.
- [x] Inconsistencia encontrada em alguns casos, formigas se sobrescrevendo?

### Testes realizados:

```
    MATRIZ_SIZE       = 100
	NUMBER_OF_ANTS    = 20
	NUMBER_OF_ITEMS   = 100
	ANT_RANGE         = 2
	NUMBER_ITERATIONS = 100000
```

### Referencias

```
Julia Handl, Joshua D. Knowles, & Marco Dorigo (2003). Ant-based clustering: a comparative study of its relative performance with respect to k-means, average link and 1d-som.

O. A. Mohamed Jafar, & R. Sivakumar (2010). Ant-based Clustering Algorithms: A Brief Survey. International Journal of Computer Theory and Engineering, 787-796.

Zahra Sadeghi, Mohammad Teshnehlab, & Mir Mohsen Pedram (2008). K-Ants Clustering - A new Strategy Based on Ant Clustering.
```