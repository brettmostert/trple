fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

staticcheck:
	@sh -c "'$(CURDIR)/scripts/staticcheck.sh'"

build:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

docker-build:
	@sh -c "'$(CURDIR)/scripts/docker-build.sh'"

.NOTPARALLEL:

.PHONY: fmtcheck staticcheck build docker-build