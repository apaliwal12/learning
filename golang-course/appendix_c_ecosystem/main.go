package main

import "fmt"

// ─────── APPENDIX C: THE GO ECOSYSTEM GUIDE ───────
// Go's standard library is incredible, but for production systems,
// you will likely use some of these battle-tested third-party libraries.

func demonstrateEcosystem() {
	fmt.Println("\n--- The Go Ecosystem (Top Packages) ---")
	
	fmt.Println("1. Routing & Web Frameworks:")
	fmt.Println("   - net/http 1.22+ (Built-in! Now supports path parameters!)")
	fmt.Println("   - github.com/go-chi/chi (Lightweight, 100% stdlib compatible router)")
	fmt.Println("   - github.com/gin-gonic/gin (Fast, popular framework with lots of middleware)")
	fmt.Println("   - github.com/gofiber/fiber (Express-like, built on fasthttp - extremely fast)")
	
	fmt.Println("\n2. Databases & ORMs:")
	fmt.Println("   - database/sql (Built-in, safe, but verbose)")
	fmt.Println("   - github.com/jmoiron/sqlx (A lightweight wrapper over database/sql to scan rows into Structs)")
	fmt.Println("   - github.com/jackc/pgx (The absolute best Postgres driver and toolkit for Go)")
	fmt.Println("   - gorm.io/gorm (The most popular ORM for Go. Heavily uses reflection. Slower, but easy to use)")
	fmt.Println("   - sqlc.dev (Compiles raw SQL queries into type-safe Go code! Highly recommended!)")
	
	fmt.Println("\n3. Configuration:")
	fmt.Println("   - github.com/spf13/viper (The industry standard for reading configs from JSON/YAML/ENV/Consul)")
	fmt.Println("   - github.com/joho/godotenv (Simple tool to load .env files)")
	
	fmt.Println("\n4. CLIs (Command Line Interfaces):")
	fmt.Println("   - github.com/spf13/cobra (Used to build Kubernetes, Docker, GitHub CLI, and more!)")
	fmt.Println("   - github.com/urfave/cli (A simpler alternative to Cobra)")
	
	fmt.Println("\n5. Testing:")
	fmt.Println("   - testing (Built-in, generally all you need)")
	fmt.Println("   - github.com/stretchr/testify (Adds assertions like `assert.Equal(t, a, b)` and mocking)")
	fmt.Println("   - github.com/DATA-DOG/go-sqlmock (For mocking database/sql)")
	
	fmt.Println("\n6. Logging & Observability:")
	fmt.Println("   - log/slog (Built-in structured logging)")
	fmt.Println("   - go.opentelemetry.io/otel (The standard for distributed tracing and metrics)")
	fmt.Println("   - github.com/uber-go/zap (Blazing fast structured logging, older than slog)")
}

func main() {
	demonstrateEcosystem()
}
