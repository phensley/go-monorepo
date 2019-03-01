TOP := $(shell pwd)
SERVICES := echo time custom lagged
GOBIN := $(TOP)/bin

all: services

services: $(SERVICES)

tidy:
	$(foreach s,$(SERVICES),cd $(TOP)/service/$(s);go mod tidy;cd $(TOP);)

%: service/%
	cd $(TOP)/service/$@; GOBIN=$(GOBIN) go install -v ...

clean: $(GOBIN)
	rm -f $(GOBIN)/*
	rmdir $(GOBIN)
