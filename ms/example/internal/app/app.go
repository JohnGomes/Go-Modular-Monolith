//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE

// Package app provides business logic.

package app

// ServiceName provides name of this microservice for logs/metrics.
const ServiceName = "example"
