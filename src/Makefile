binary_name=jean-claude
build_dir=../build

help:
	@echo "Usage:"
	@echo " build - builds ${binary_name} binaries."
	@echo " clean - removes temp files, binaries and other generated files."

.PHONY: build
build:
	GOARCH=amd64 GOOS=linux go build -o ${build_dir}/${binary_name}

run: build
	${build_dir}/${binary_name}

clean:
	@rm -Rf ${build_dir}