binary_name=jean-claude
build_dir=build
src_dir=src

help:
	@echo "Usage:"
	@echo " build - builds ${binary_name} binaries."
	@echo " clean - removes temp files, binaries and other generated files."

.PHONY: build
build:
	GOARCH=amd64 GOOS=linux cd ${src_dir} && go build -o ../${build_dir}/${binary_name}

run: build
	${build_dir}/${binary_name}

clean:
	@rm -Rf ${build_dir}