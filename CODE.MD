# Functions

As this project evolves some functions are no longer necessary, however may be useful in the future. These functions can be stored within this file.

```go
func credentials() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)
	return strings.TrimSpace(password)
}
```