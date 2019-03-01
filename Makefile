
SERVICES := echo time
all: services

services: $(SERVICES)

tidy:
	$(foreach s,$(SERVICES),cd service/$(s);go mod tidy;cd ../..;)

%: service/%
	mkdir -p bin
	cd service/$@; go build -o ../../bin/$@

