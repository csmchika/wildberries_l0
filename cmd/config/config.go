package config

import (
	"os"
)

func ConfigSetup() {
	// Database settings
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "admin")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PORT", "5432")

	os.Setenv("DB_POOL_MAXCONN", "5")
	os.Setenv("DB_POOL_MAXCONN_LIFETIME", "300")

	// NATS-Streaming settings
	os.Setenv("NATS_HOSTS", "localhost")
	os.Setenv("NATS_CLUSTER_ID", "l0-cluster")
	os.Setenv("NATS_CLIENT_ID", "l0-client")
	os.Setenv("NATS_SUBJECT", "wl")
	os.Setenv("NATS_DURABLE_NAME", "Replica-1")
	os.Setenv("NATS_ACK_WAIT_SECONDS", "30")

	// Cache settings
	os.Setenv("CACHE_SIZE", "10")
	os.Setenv("APP_KEY", "WB-1")
}
