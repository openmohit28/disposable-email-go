// Package disposableemail provides functionality to check if an email address
// uses a disposable domain.
//
// Example usage:
//
//	import "github.com/yourusername/disposable-email-checker"
//
//	func main() {
//		email := "test@example.com"
//		if disposableemail.IsDisposable(email) {
//			fmt.Println("This is a disposable email address")
//		} else {
//			fmt.Println("This is not a disposable email address")
//		}
//	}
package disposableemail