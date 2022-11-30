# Alura - Go: Orientação a Objetos

Exercícios do curso Go: Orientação a Objetos da plataforma de estudos Alura

# Particularidades
- Zero-initialization - Conceito que determina a forma com que o GO atua na inicialização de variáveis <br/>
  <img src="https://caelum-online-public.s3.amazonaws.com/1365-golang-oo/aula_01/smais_aula01.png" width="300px">

- Variadic function - Função com quantidade indeterminada de parâmetros

```golang 
package main

import (
    "fmt"
)

func Somando(numeros ...int) int {
    resultadoDaSoma := 0
    for _, numero := range numeros {
        resultadoDaSoma += numero
    }
    return resultadoDaSoma
}

func main() {
    fmt.Println(Somando(1))
    fmt.Println(Somando(1,1))
    fmt.Println(Somando(1,1,1))
    fmt.Println(Somando(1,1,2,4))
}
```

- Visibilidade - Na linguagem Go, se escrevermos nomes de variáveis deixando a primeira letra minúscula, elas ficarão visíveis apenas no arquivo em que foram declaradas (private).
Primeira letra maiúscula altera a visibilidade para fora do pacote (public)
- Não existe o conceito de Classe e consequentemente herança na linguagem GO.
- É possível fazer a composição de structs, onde um variável de um struct A, pode ser do tipo struct B

# Structs
Structs são tipos que contem outros tipos. Quando não passamos nenhuma informação para a struct, ela é criada automaticamente com os valores padrão. <br/>

Existem 3 formas de utilizar/inicializar uma struct:

```golang
package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
  /*1 - Dessa forma você não é obrigado a enviar todos as propriedades declaradas na struct, porém é necessário especificar com chave e valor*/
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

  //2 - Dessa forma você é obrigado a enviar todos as propriedades declaradas na struct
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

  //3 -Trabalhando com ponteiro de forma explícita.
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)
	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
}

```


# Estruturas de dados
## Array
Tipados e com quantidade pré determinada. 

### Funções
- len - Tamanho do array
- range - Retorna o item e o index do array
- cap - Retorna a capacidade do array

## Slices
Abstrações de array onde a própria linguagem gerencia entre outras coisas, o tamanho do array. 
Quando estouramos o tamanho do slice, ele dobra o tamanho. 

Exemplo: 
```golang
valores: = [ ]string{"Paloma", "Arthur"}
fmt.Println(len(valores)) -> retorna tamanho 2
fmt.Println(cap(valores)) -> retorna capacidade 2
```

```golang
valores := append(valores, "Vera")
fmt.Println(len(valores)) -> retorna tamanho 3
fmt.Println(cap(valores)) -> retorna capacidade 4
```
# Pacotes
## fmt 
Equivalente ao console.log e suas variações do js. Serve para imprimir dados em tela
## reflect 
Extrai informações em tempo de execução.
## os 
O Pacote OS fornece uma interface independente de plataforma para a funcionalidade do sistema operacional.
 - Flags para abertura de arquivos: https://pkg.go.dev/os#pkg-constants
  
## net 
Pacote net fornece uma interface portátil para E/S de rede, incluindo TCP/IP, UDP, resolução de nome de domínio e soquetes de domínio Unix.
## time
Pacote fornece funcionalidade para medir e exibir o tempo.
- Formatando exibição de tempo: https://pkg.go.dev/time
## io
O pacote io fornece interfaces básicas para primitivas de E/S.
## bufio
O pacote bufio implementa E/S em buffer. Ele envolve um objeto io.Reader ou io.Writer, criando outro objeto (Reader ou Writer) que também implementa a interface, mas fornece buffer e alguma ajuda para E/S textual.
## strconv 
O pacote strconv implementa conversões de e para representações de string de tipos de dados básicos.