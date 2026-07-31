package main

import (
	"net"
	"net/http"
	"regexp"
	"strings"
)

type ValidationResult struct {
	Email       string
	SyntaxValid bool
	MXValid     bool
	Deliverable bool
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{
		"message": "server is running",
		"status":  "ok",
	})
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	result := validateEmail(email)

	type EmailResponse struct {
		Email       string `json:"email"`
		Syntax      bool   `json:"syntax"`
		Domain      bool   `json:"domain"`
		Deliverable bool   `json:"deliverable"`
	}

	response := EmailResponse{
		Email:       result.Email,
		Syntax:      result.SyntaxValid,
		Domain:      result.MXValid,
		Deliverable: result.Deliverable,
	}

	WriteJSON(w, http.StatusOK, response)
}

// checkSyntax uses Regex to ensure standard email formatting.
func checkSyntax(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).
		MatchString(email)
}

// checkDomain extracts the domain and queries the network for Mail Exchange (MX) records.
func checkDomain(email string) bool {
	b := strings.Split(email, "@")

	if len(b) != 2 {
		return false
	}

	domain := b[1]

	dns_rec, err := net.LookupMX(domain)
	if err != nil || len(dns_rec) == 0 {
		return false
	}

	return true
}

// validateEmail orchestrates the syntax and domain checks.
func validateEmail(email string) ValidationResult {
	result := ValidationResult{
		Email:       email,
		SyntaxValid: false,
		MXValid:     false,
		Deliverable: false,
	}

	// 1. Call checkSyntax(email). If it returns false, return the 'result' struct immediately.
	if !checkSyntax(email) {
		return result
	}

	// 2. Set syntax valid to true.
	result.SyntaxValid = true

	// 3. Call checkDomain(email). If it returns false, return the 'result' struct immediately.
	if !checkDomain(email) {
		return result
	}

	// 4. Set MX valid to true.
	result.MXValid = true

	// 5. If both passed, set deliverable to true and return the struct.
	result.Deliverable = true

	return result
}

// // processEmailsConcurrent uses Goroutines to validate emails in parallel.
// func processEmailsConcurrent(emails []string) []ValidationResult {
// 	var results []ValidationResult
//
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
//
// 	for _, email := range emails {
// 		// TODO 4: Add 1 to the WaitGroup counter.
// 		wg.Add(1)
//
// 		// TODO 5: Launch a Goroutine
// 		go func(email string) { // Inside the Goroutine:
// 			// 1. Defer the WaitGroup Done() method to signal completion.
// 			defer wg.Done()
//
// 			// 2. Call validateEmail(email) and store the result.
// 			result := validateEmail(email)
//
// 			// TODO 6: Prevent Race Conditions
// 			// 3. Lock the Mutex
// 			mu.Lock()
//
// 			// 4. Append the result to the 'results' slice
// 			results = append(results, result)
//
// 			// 5. Unlock the Mutex
// 			mu.Unlock()
// 		}(email)
// 	}
//
// 	// TODO 7: Wait for all Goroutines to finish before continuing
// 	wg.Wait()
//
// 	return results
// }
