package main

// TODO 1: Implement MultipleReturns
// It should take a string and return its length and the string itself.
func MultipleReturns(s string) (int, string) {
	return 0, "" // Fix me
}

// TODO 2: Implement SumVariadic
// It should take any number of integers and return their sum.
func SumVariadic(nums ...int) int {
	return 0 // Fix me
}

// TODO 3: Implement MultiplyBy
// It should return a closure that multiplies its input by the given factor.
func MultiplyBy(factor int) func(int) int {
	return func(n int) int {
		return 0 // Fix me
	}
}

// TODO 4: Implement ExecuteWithCleanup
// It should execute the `mainFunc`, but ensure that `cleanupFunc` is ALWAYS
// called after `mainFunc` finishes, even if `mainFunc` panics.
// Hint: use defer.
func ExecuteWithCleanup(mainFunc func(), cleanupFunc func()) {
	// Fix me
}

// TODO 5: Implement FilterWords
// It should return a new slice of strings containing only the words that
// satisfy the predicate.
func FilterWords(words []string, predicate func(string) bool) []string {
	return nil // Fix me
}

// TODO 6: Implement SafeExecute
// It should execute `f`. If `f` panics, it should recover the panic and return
// the panic message as an error. If `f` completes normally, return nil.
// Hint: use defer, recover(), and fmt.Errorf.
func SafeExecute(f func()) error {
	return nil // Fix me
}

// TODO 7: Implement MemoizeFibonacci
// It should return a closure that calculates the nth Fibonacci number.
// The closure should maintain a cache (map or slice) of previously calculated
// Fibonacci numbers to avoid redundant calculations.
// Fib(0) = 0, Fib(1) = 1, Fib(n) = Fib(n-1) + Fib(n-2).
func MemoizeFibonacci() func(int) int {
	return func(n int) int {
		return 0 // Fix me
	}
}
