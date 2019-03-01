TOP := $(shell pwd)
SERVICES := echo time custom lagged
all: services

services: $(SERVICES)

tidy:
	$(foreach s,$(SERVICES),cd $(TOP)/service/$(s);go mod tidy;cd $(TOP);)

%: service/%
	mkdir -p bin
	cd $(TOP)/service/$@; go build -o $(TOP)/bin/$@

clean:
	$(foreach s,$(SERVICES),rm -f $(TOP)/bin/$(s);)
