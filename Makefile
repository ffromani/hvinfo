all: binaries

binaries:
	./hack/build.sh ${VERSION}

prepare:
	mkdir -p binaries

clean:
	rm -rf binaries

release: clean prepare binaries

.PHONY: all release prepare clean binaries
