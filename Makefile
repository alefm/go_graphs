build:
	@go build .
	@./go_graphs
	@dot -Tpng output.dot -o sample.png

run:
	@./go_graphs
	@dot -Tpng output.dot -o sample.png

dot:
	@dot -Tpng output.dot -o sample.png

clean:
	@rm *.png
	@rm *.dot
	@rm go_graphs
