COMPANY_ROOT:=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

$(COMPANY_ROOT)/-swagger: swagger-to-go
	$(BIN)/swagger-to-go -pkg companypb -file $(COMPANY_ROOT)/proto/company.swagger.json > $(COMPANY_ROOT)/proto/company.swagger.pb.go


$(COMPANY_ROOT)/-migration: $(BIN)/go-bindata
	echo "Companies"
	cd $(COMPANY_ROOT)/migrations && $(BIN)/go-bindata -nometadata -o migration.gen.go -nomemcopy=true -pkg=migrations ./postgres/...

#$(COMPANY_ROOT)-lint: $(LINTER)
#	$(LINTERCMD) $(COMPANY_ROOT)/...

.PHONY: $(COMPANY_ROOT)/-swagger #$(COMPANY_ROOT)-migration $(COMPANY_ROOT)-lint