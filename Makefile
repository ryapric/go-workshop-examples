MODULES := $(shell find . -maxdepth 1 -type d -not -path '*.git' -not -path '.')
HERE := $(shell pwd)
MOD_LOOP_START := for mod in $(MODULES); do echo "processing $${mod}" && cd "$${mod}" && 
MOD_LOOP_END := && echo "done with $${mod}" && cd $(HERE) ; done

tidy:
	@$(MOD_LOOP_START) go mod tidy $(MOD_LOOP_END)

test:
	@$(MOD_LOOP_START) go test -v -cover $(MOD_LOOP_END)
