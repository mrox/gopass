## Compline arm64
```
CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build -a -ldflags '-linkmode external -extldflags "-static"' cmd/acp/main.go
```

## Cross arm64

### Tutorial
```
https://lucor.dev/post/cross-compile-golang-fyne-project-using-zig/
```

### ARM64

```
# install the GUI and CLI dependencies
dpkg --add-architecture arm64
apt-get update
apt-get install -y -q --no-install-recommends \
    libgl-dev:arm64 \
    libx11-dev:arm64 \
    libxrandr-dev:arm64 \
    libxxf86vm-dev:arm64 \
    libxi-dev:arm64 \
    libxcursor-dev:arm64 \
    libxinerama-dev:arm64

# create dist dir
mkdir -p dist/linux-arm64

# build
CGO_ENABLED=1 \
GOOS=linux \
GOARCH=arm64 \
PKG_CONFIG_LIBDIR=/usr/lib/aarch64-linux-gnu/pkgconfig \
CC="zig cc -target aarch64-linux-gnu -isystem /usr/include -L/usr/lib/aarch64-linux-gnu" \
CXX="zig c++ -target aarch64-linux-gnu -isystem /usr/include -L/usr/lib/aarch64-linux-gnu" \
go build -trimpath -o dist/linux-arm64 ./cmd/...
```