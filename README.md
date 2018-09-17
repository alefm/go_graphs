# go_graphs

#### Sobre o projeto
Explicar o projeto.
#### Visualização do Grafo

Todos os grafos gerados neste projetos, serão exportados para a linguagem de  descrição `dot` em um arquivo chamado `output.dot`.

Este arquivo deverá ser lido por um software de visualização gráfica de grafos chamado [Graphviz](http://www.graphviz.org).

[Linguagem Dot](https://www.graphviz.org/doc/info/lang.html) <br>
[Exemplos](https://github.com/gyuho/learn/tree/master/doc/go_graph_interface#graph-visualization-with-dot)

##### [Instalação Graphviz](https://graphviz.gitlab.io/download/)

* Linux:

    ```sudo apt-get -y install graphviz```
    <br>```yum install graphviz*```
    
* MacOS:

    ```brew install graphviz```
    
##### Utilização

Após a geração do arquivo `output.dot`, você poderá executar um destes comandos no terminal para visualização dos grafos:
```shell
dot -Tpng output.dot -o sample.png
dot -Tpdf output.dot -o sample.pdf
```


Using Makefile:

```
# All
make build

# Create output.dot
make run

# Create sample.png
make dot
```
