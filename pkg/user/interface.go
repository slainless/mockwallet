package user

import "github.com/slainless/mock-fintech-platform/pkg/platform"

// based on the requirements, we need to attach
// - user manager
// - monetary account manager
// - transaction history manager
type IUserService interface {
	UserManager() platform.UserManager
	TransactionHistoryManager() platform.TransactionHistoryManager
	AccountManager() platform.MonetaryAccountManager
}
