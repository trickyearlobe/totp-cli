BUILD_TIMESTAMP = `date '+%Y%m%d%H%M%S'`
GIT_COMMIT = `git rev-list -1 HEAD`
GIT_REPO = github.com/trickyearlobe/totp-cli
BUILT_BY = `whoami`
GOPATH = `go env GOPATH`

LDFLAGS = "-X github.com/trickyearlobe.com/totp-cli/cmd.GitCommit=$(GIT_COMMIT) -X github.com/trickyearlobe.com/totp-cli/cmd.BuildTimestamp=$(BUILD_TIMESTAMP) -X github.com/trickyearlobe.com/totp-cli/cmd.GitRepo=$(GIT_REPO) -X github.com/trickyearlobe.com/totp-cli/cmd.BuiltBy=$(BUILT_BY)"

zipfile: build
	tar -zcvf totp-cli-$(BUILD_TIMESTAMP).tgz target

build:
	@echo using LDFLAGS $(LDFLAGS)
	GOOS=darwin   GOARCH=amd64 go build -o target/mac/amd64/totp-cli         -ldflags=$(LDFLAGS)
	GOOS=darwin   GOARCH=arm64 go build -o target/mac/arm64/totp-cli         -ldflags=$(LDFLAGS)
	GOOS=linux    GOARCH=amd64 go build -o target/linux/amd64/totp-cli       -ldflags=$(LDFLAGS)
	GOOS=linux    GOARCH=arm64 go build -o target/linux/arm64/totp-cli       -ldflags=$(LDFLAGS)
	GOOS=windows  GOARCH=amd64 go build -o target/windows/amd64/totp-cli.exe -ldflags=$(LDFLAGS)
	GOOS=windows  GOARCH=arm64 go build -o target/windows/arm64/totp-cli.exe -ldflags=$(LDFLAGS)
	GOOS=openbsd  GOARCH=amd64 go build -o target/openbsd/amd64/totp-cli     -ldflags=$(LDFLAGS)
	GOOS=openbsd  GOARCH=amd64 go build -o target/openbsd/arm64/totp-cli     -ldflags=$(LDFLAGS)
	GOOS=netbsd   GOARCH=amd64 go build -o target/netbsd/amd64/totp-cli      -ldflags=$(LDFLAGS)
	GOOS=netbsd   GOARCH=amd64 go build -o target/netbsd/arm64/totp-cli      -ldflags=$(LDFLAGS)

install:
	go install -ldflags=$(LDFLAGS)

clean:
	rm -rf target
	rm -f *.tgz
	rm -f $(GOPATH)/bin/totp-cli