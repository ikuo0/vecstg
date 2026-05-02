# vecstg

# environment

```bash
sudo apt-get update
sudo apt-get install -y \
  pkg-config \
  libx11-dev \
  libasound2-dev \
  libgl1-mesa-dev \
  libxcursor-dev \
  libxi-dev \
  libxinerama-dev \
  libxrandr-dev \
  libxxf86vm-dev
```

```bash
go env GOMODCACHE
# /go/pkg/mod
```

# modules
## ebiten
```bash
go mod init github.com/hajimehoshi/ebiten/v2 latest
go mod tidy
```