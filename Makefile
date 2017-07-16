##### Example ######

example: depend bootstrap delete context

delete:
	rm -f ./hello.go
	rm -f ./main.go

##### Convenience targets ######

REPO:=github.com/pei0804/goa-oauth2-practice
GAE_PROJECT:=projectName

init: depend bootstrap context
gen: clean generate context

depend:
	@which glide || go get -v github.com/Masterminds/glide
	@glide install

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design
