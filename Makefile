build:
	@go build .
	@./go_graphs -> output.dot
	@dot -Tpng output.dot -o sample.png


run:
	@./go_graphs -> output.dot
	@dot -Tpng output.dot -o sample.png


dot:
	@dot -Tpng output.dot -o sample.png
