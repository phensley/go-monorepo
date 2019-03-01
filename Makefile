TOP := $(shell pwd)
SERVICES := echo time custom lagged
TOOLS := \
	golang.org/x/tools/cmd/stringer \
	golang.org/x/tools/cmd/cover

BIN := $(TOP)/bin
TOOLBIN := $(TOP)/tools

all: services

services: $(SERVICES)

%: service/%
	cd $(TOP)/service/$@ ; \
	GOBIN=$(TOOLBIN) go generate ; \
	GOBIN=$(BIN) go install -v ... ;

clean: $(BIN)
	rm -f $(BIN)/*
	rmdir $(BIN)

bootstrap:
	$(foreach t,$(TOOLS),GO111MODULE=off GOBIN=$(TOOLBIN) go get $(t);)

tidy:
	$(foreach s,$(SERVICES),cd $(TOP)/service/$(s);go mod tidy;cd $(TOP);)
