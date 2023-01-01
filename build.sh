#!/bin/bash

darwin_archs=(amd64 arm64)

for darwin_arch in ${darwin_archs[@]}
do
    env GOOS=darwin GOARCH=${darwin_arch} go build -tags free -o builds/free/darwin/${darwin_arch}/audiofile main.go
    env GOOS=darwin GOARCH=${darwin_arch} go build -tags pro -o builds/pro/darwin/${darwin_arch}/audiofile main.go
    env GOOS=darwin GOARCH=${darwin_arch} go build -tags profile -o builds/profile/darwin/${darwin_arch}/audiofile main.go
done

linux_archs=(386 amd64 arm arm64 loong64 mips mips64 mips64le mipsle ppc64 ppc64le riscv64 s390x)

for linux_arch in ${linux_archs[@]}
do
    env GOOS=linux GOARCH=${linux_arch} go build -tags free -o builds/free/linux/${linux_arch}/audiofile main.go
    env GOOS=linux GOARCH=${linux_arch} go build -tags pro -o builds/pro/linux/${linux_arch}/audiofile main.go
    env GOOS=linux GOARCH=${linux_arch} go build -tags profile -o builds/profile/linux/${linux_arch}/audiofile main.go
done

windows_arches=(386 amd64 arm arm64)

for windows_arch in ${windows_arches[@]}
do
    env GOOS=windows GOARCH=${windows_arch} go build -tags free -o builds/free/windows/${windows_arch}/audiofile.exe main.go
    env GOOS=windows GOARCH=${windows_arch} go build -tags pro -o builds/pro/windows/${windows_arch}/audiofile.exe main.go
    env GOOS=windows GOARCH=${windows_arch} go build -tags profile -o builds/profile/windows/${windows_arch}/audiofile.exe main.go
done