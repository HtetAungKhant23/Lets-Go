package main

import "fmt"

func (a *analytics) handleEmailBounce(e email) error {
	updStatusErr := e.recipient.updateStatus(e.status)
	analyticErr := a.track(e.status)

	if updStatusErr != nil {
		return fmt.Errorf("error updating user status: %w", updStatusErr)
	}

	if analyticErr != nil {
		return fmt.Errorf("error tracking user bounce: %w", analyticErr)
	}

	return nil
}

type perm string

const (
	Read  perm = "read"
	Write perm = "write"
	Exec  perm = "exec"
)

func checkPermission(p perm) {
	fmt.Println(p)
}

func main() {
	var Admin = "admin"
	var User = perm("user")

	checkPermission(perm(Admin))
	checkPermission(User)

}
