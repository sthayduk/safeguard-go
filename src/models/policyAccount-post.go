package models

// LinkToUser links the current PolicyAccount to the specified User.
// It returns a slice of PolicyAccount and an error if any occurs during the linking process.
//
// Parameters:
//
//	user - The User to link the PolicyAccount to.
//
// Returns:
//
//	[]PolicyAccount - A slice containing the linked PolicyAccount.
//	error - An error if any issues occur during the linking process.
func (p PolicyAccount) LinkToUser(user User) ([]PolicyAccount, error) {
	return AddLinkedAccounts(p.client, user, []PolicyAccount{p})
}

// UnlinkFromUser removes the association between the given PolicyAccount and the specified User.
// It returns a slice of PolicyAccount and an error if the operation fails.
//
// Parameters:
//
//	user - The User from whom the PolicyAccount should be unlinked.
//
// Returns:
//
//	[]PolicyAccount - A slice containing the PolicyAccount after unlinking.
//	error - An error if the unlinking operation fails.
func (p PolicyAccount) UnlinkFromUser(user User) ([]PolicyAccount, error) {
	return RemoveLinkedAccounts(p.client, user, []PolicyAccount{p})
}
