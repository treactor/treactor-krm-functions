include ./Makefile.Common

ALL_MODULES := $(shell find . -type f -name "go.mod" -exec dirname {} \; | sort | egrep  '^./' )
all-modules:
	@echo $(ALL_MODULES) | tr ' ' '\n' | sort
MODULEDIRS = $(ALL_MODULES:%=for-all-target-%)
for-all-target: $(MODULEDIRS)
$(MODULEDIRS):
	$(MAKE) -C $(@:for-all-target-%=%) $(TARGET)
.PHONY: for-all-target

ALL_FUNCTIONS := $(shell find . -type f -name "Dockerfile" -exec dirname {} \; | sort | egrep  '^./' )
all-functions:
	@echo $(ALL_MODULES) | tr ' ' '\n' | sort
FUNCTIONDIRS = $(ALL_FUNCTIONS:%=for-all-function-%)
for-all-function: $(FUNCTIONDIRS)
$(FUNCTIONDIRS):
	$(MAKE) -C $(@:for-all-function-%=%) $(TARGET)
.PHONY: for-all-function


.PHONY: gotidy
gotidy:
	$(MAKE) for-all-target TARGET="tidy"

.PHONY: unstable
docker-unstable:
	$(MAKE) for-all-function TARGET="unstable"

.PHONY: gotest
gotest:
	$(MAKE) for-all-target TARGET="test"
