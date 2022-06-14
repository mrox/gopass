CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build -a -ldflags '-linkmode external -extldflags "-static"' cmd/acp/main.go


#!/usr/bin/bash
# archs=(amd64 arm64 ppc64le ppc64 s390x)

# for arch in ${archs[@]}
# do
#         env GOOS=linux GOARCH=${arch} go build -o prepnode_${arch}
# done