main_package_path = .
binary_name = website

export DB_FILENAME = /tmp/website.db

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)" || (echo "ERROR: there is unstaged/uncommitted changes." && exit 1)


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test.cover: run all tests and display coverage
.PHONY: test.cover
test.cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tools: install required tools for this project
.PHONY: tools
tools:
	go install github.com/a-h/templ/cmd/templ@latest
	brew install sqlite3
	brew install sqlite-utils

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## db.init: initialize sqlite db file if it doesn't exist yet.
.PHONY: db.init
db.init:
	@([ ! -f $(DB_FILENAME) ] && sqlite3 $(DB_FILENAME) "VACUUM;") || echo "database already exists"

## gen: generate code
.PHONY: gen
gen:
	go generate ./...

## build: build the application
.PHONY: build
build: gen
	CGO_ENABLED=1 go build -o=/tmp/bin/${binary_name} ${main_package_path}

## run: run the application @ localhost
.PHONY: run
run: build db.init
	/tmp/bin/${binary_name} --addr localhost:8080

## run.live: run the application @ localhost with reloading on file changes
.PHONY: run.live
run.live: db.init
	go run github.com/air-verse/air@v1.52.3 \
		--build.cmd "make build" --build.bin "/tmp/bin/${binary_name}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, templ, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true" \
		-- --addr localhost:8080


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: confirm audit no-dirty
	git push
