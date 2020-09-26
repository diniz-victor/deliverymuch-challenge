# Set project's name
PROJECT = deliverymuch-challenge
export PROJECT

PKG = github.com/$(PROJECT)
BASE_IMAGE = $(PROJECT)

# Basic information regarding the repository
VERSION ?= $(shell git describe --tags --always)
export VERSION

# Our minimum test coverage target, in %
MINIMUM_COVERAGE = 70

# Images that will be automatically built and published
IMAGES = consumer

vendor-build:
	go mod vendor && go build

setup: vendor-build run

run:
	@docker-compose up

include tooling/build/make/*.mk
