sca:
	@sh -c "'$(CURDIR)/scripts/sca.sh'"

static-code-analysis:
	@sh -c "'$(CURDIR)/scripts/sca.sh'"

build: sca
	@sh -c "'$(CURDIR)/scripts/build.sh'"

docker-build: sca
	@sh -c "'$(CURDIR)/scripts/docker-build.sh'"

docker-build-release: sca
	@sh -c "'$(CURDIR)/scripts/docker-build-release.sh'"

docker-clean-all:
	@sh -c "'$(CURDIR)/scripts/docker-clean-all.sh'"

docker-clean:
	@sh -c "'$(CURDIR)/scripts/docker-clean.sh'"

proto:
	@sh -c "'$(CURDIR)/scripts/proto.sh'"

.NOTPARALLEL:

.PHONY: sca static-code-analysis build docker-build docker-clean-all docker-clean docker-build-release proto