# 🚀 The Definitive Go Mastery Guide

Welcome to the definitive, comprehensive, end-to-end Go (Golang) mastery guide! This repository is designed as a **"Book-in-Code"**. Instead of reading a static textbook, you will learn the *why*, the *how*, and the historical context of Go directly through heavily commented, executable Go code.

This guide takes you from an absolute beginner to a master level, acting as a one-stop resource. Every single module is a standalone, runnable Go project with a corresponding assignment test suite to validate your learning.

## 📖 How to Use This Guide

Each folder in this repository represents a distinct topic. Inside most folders, you will find:
1. `go.mod`: Making the module standalone.
2. `main.go` / `<topic>.go`: The core reading material. These files contain rich, narrative-driven comments explaining the theory, rhetoric, and code examples. Run them using `go run .`.
3. `assignment.go`: A set of incomplete `TODO` functions for you to implement.
4. `assignment_test.go`: A robust, table-driven test suite. Run `go test -v` to check your work!

---

## 🗂️ Course Index

### Part 1: Foundations
The bedrock of the Go programming language. Learn how Go views memory, types, and the underlying philosophy of the language.

* **[Module 01: Fundamentals](./module01_fundamentals)**
  * **Concepts:** Basic Syntax, Primitive Types, The Philosophy of Memory (Zero Values), Pointers (Stack vs Heap basics), Constants, and Enums (`iota`).
* **[Module 02: Functions & Control Flow](./module02_functions)**
  * **Concepts:** First-Class Functions, Closures, Multiple Return Values, Named Returns, Error Handling Basics, `defer`, `panic`, and `recover`.
* **[Module 03: Data Structures](./module03_data_structures)**
  * **Concepts:** Arrays vs Slices, Slice Headers (Pointer, Length, Capacity), `append` internals, Maps, Structs, Struct Tags, and Memory Padding.
* **[Module 04: Strings & Runes](./module04_strings)**
  * **Concepts:** ASCII vs UTF-8, Strings as read-only byte slices, Runes (int32), and efficient string building (`strings.Builder`).
* **[Module 05: Interfaces & Composition](./module05_interfaces)**
  * **Concepts:** Implicit Interface Satisfaction, The Empty Interface (`any`), Type Assertions, Type Switches, and Composition over Inheritance (Struct Embedding).
* **[Module 06: Packages & Modules](./module06_packages)**
  * **Concepts:** `go.mod`, encapsulation (capital vs lowercase), initialization (`init()` functions), and standard project layout.
* **[Module 07: I/O & Interfaces](./module07_io)**
  * **Concepts:** `io.Reader`, `io.Writer`, Decorator pattern with standard streams (`os.Stdout`, `bufio`), file reading/writing.
* **[Module 08: Time & Scheduling](./module08_time)**
  * **Concepts:** `time.Time`, `time.Duration`, Monotonic Clocks, Timers, Tickers, and parsing/formatting (the magic `2006-01-02 15:04:05` layout).

### Part 2: Concurrency & Networking
Go's superpower. Understand how to design highly concurrent systems using CSP (Communicating Sequential Processes) and standard network protocols.

* **[Module 09: Concurrency Foundations](./module09_concurrency)**
  * **Concepts:** Goroutines (M:N scheduling), Channels (Unbuffered vs Buffered), `sync.WaitGroup`, `sync.Mutex`, `sync.Once`, and race conditions.
* **[Module 10: Advanced Concurrency](./module10_advanced_concurrency)**
  * **Concepts:** The `select` statement, `context.Context` (cancellations and timeouts), pipelines, and fan-out/fan-in patterns.
* **[Module 11: Networking & HTTP](./module11_http)**
  * **Concepts:** Building custom HTTP Clients with timeouts, `http.Handler` interfaces, RESTful JSON APIs, and writing HTTP Middleware.
* **[Module 12: Databases & Storage](./module12_databases)**
  * **Concepts:** The `database/sql` interface, driver registration (blank imports), connection pooling, Transactions, and the DAO/Repository Pattern.

### Part 3: Robustness, Modern Go & Advanced Topics
Level up to senior engineering. Learn to observe, test, containerize, and utilize the newest and most dangerous parts of the language.

* **[Module 13: Crypto & Security](./module13_crypto)**
  * **Concepts:** Fast Hashing (SHA-256) vs Slow Hashing (bcrypt), Symmetric Encryption (AES-GCM), and secure nonce generation.
* **[Module 14: Robust Engineering](./module14_robustness)**
  * **Concepts:** Go 1.13+ Error Wrapping (`fmt.Errorf("%w")`, `errors.Is`, `errors.As`), Dependency Injection for Mocking, and Graceful Shutdowns.
* **[Module 15: Observability](./module15_observability)**
  * **Concepts:** Standard `log` vs Modern Structured Logging (`log/slog` in Go 1.21+), JSON Handlers, and injecting Trace IDs using Context.
* **[Module 16: Modern Go (Go 1.18 - Go 1.23) ](./module16_modern_go)**
  * **Concepts:** Generics (Type Parameters and Constraints), and custom Iterators (`iter.Seq` in Go 1.23+).
* **[Module 17: Build System & Deployment](./module17_build)**
  * **Concepts:** Build constraints/tags (`//go:build`), Linker flags (`ldflags` for version injection), Cross-compilation, and Multi-stage Docker builds.
* **[Module 18: Advanced Internals](./module18_internals)**
  * **Concepts:** The Go Memory Model, Escape Analysis (Stack vs Heap), Garbage Collection (`GOGC`), Runtime Reflection (`reflect`), and bypassing the type system with `unsafe`.

---

## 🛠️ Capstone Projects
Practical projects that synthesize multiple modules into realistic, robust applications.

* **[Project A: Concurrent Web Scraper](./project_a_scraper)**
  * Uses HTTP Clients, Goroutines, Channels, WaitGroups, and Context for timeout-bound parallel web scraping.
* **[Project B: TCP Load Balancer](./project_b_load_balancer)**
  * A low-level networking project using `net.Listener`, `net.Conn`, Mutex-backed round-robin algorithms, and `io.Copy` for fast byte streaming.
* **[Project C: gRPC + REST Microservice Concept](./project_c_grpc)**
  * Demonstrates the usage of Protocol Buffers (`.proto`) and the conceptual layout of a dual gRPC and HTTP/JSON (grpc-gateway) microservice.
* **[Project D: Key-Value Store with Write-Ahead Log (WAL)](./project_d_kvstore)**
  * A database internals project combining `sync.RWMutex`, File I/O, and recovery mechanisms to persist data across crashes.

---

## 📚 Appendices
Quick reference guides and deep dives into the culture of Go engineering.

* **[Appendix A: Go Anti-Patterns](./appendix_a_antipatterns)**
  * What *not* to do: Goroutine leaks, pre-Go 1.22 loop variable capture bugs, error shadowing, and naked returns.
* **[Appendix B: Patterns Cookbook](./appendix_b_patterns)**
  * Classic Go idioms, focusing heavily on the highly-used Functional Options Pattern for configuring structs.
* **[Appendix C: The Go Ecosystem Guide](./appendix_c_ecosystem)**
  * A curated list of the industry-standard third-party libraries (e.g., Gin, SQLX, PGX, Viper, Cobra, Testify, Zap).
* **[Appendix D: Interview Prep Q&A](./appendix_d_interview)**
  * A rapid-fire Q&A covering the most commonly asked technical questions in Senior Go Engineering interviews.
