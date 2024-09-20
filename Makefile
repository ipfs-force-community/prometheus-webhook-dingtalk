ldflags=-X=github.com/timonwong/prometheus-webhook-dingtalk/version.CurrentCommit=+git.$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif

GOFLAGS+=-ldflags="$(ldflags)"

build:
	go build -o dingtalk $(GOFLAGS) cmd/main.go
.PHONY: build
