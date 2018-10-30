build:
	@go build .
	@./go_graphs

run:
	@./go_graphs

clean:
	@rm *.png
	@rm *.dot
	@rm go_graphs
