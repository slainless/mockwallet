package query

import (
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/slainless/mock-fintech-platform/pkg/internal/artifact/database/mock_fintech/public/table"
	"github.com/slainless/mock-fintech-platform/pkg/platform"
)

func GetAllAccounts(ctx context.Context, db *sql.DB, userUUID uuid.UUID) ([]platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(table.PaymentAccounts.UserUUID.EQ(UUID(userUUID)))

	accounts := make([]platform.PaymentAccount, 0)
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccount(ctx context.Context, db *sql.DB, accountUUID uuid.UUID) (*platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(table.PaymentAccounts.UUID.EQ(UUID(accountUUID)))

	var account platform.PaymentAccount
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetAccountWhereUser(ctx context.Context, db *sql.DB, userUUID, accountUUID uuid.UUID) (*platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(
			table.PaymentAccounts.UserUUID.EQ(UUID(userUUID)).
				AND(table.PaymentAccounts.UUID.EQ(UUID(accountUUID))),
		)

	var account platform.PaymentAccount
	err := stmt.QueryContext(ctx, db, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetTwoAccounts(ctx context.Context, db *sql.DB, FirstUUID, SecondUUID uuid.UUID) (*platform.PaymentAccount, *platform.PaymentAccount, error) {
	stmt := SELECT(table.PaymentAccounts.AllColumns).
		FROM(table.PaymentAccounts).
		WHERE(
			table.PaymentAccounts.UUID.EQ(UUID(FirstUUID)).
				OR(table.PaymentAccounts.UUID.EQ(UUID(SecondUUID))),
		)

	var accounts []platform.PaymentAccount
	err := stmt.QueryContext(ctx, db, &accounts)
	if err != nil {
		return nil, nil, err
	}

	if len(accounts) == 1 {
		if accounts[0].UUID == FirstUUID {
			return &accounts[0], nil, nil
		}
		return nil, &accounts[0], nil
	}

	var first, second *platform.PaymentAccount
	if accounts[0].UUID == FirstUUID {
		first = &accounts[0]
		second = &accounts[1]
	} else {
		first = &accounts[1]
		second = &accounts[0]
	}

	return first, second, nil
}

func CheckOwner(ctx context.Context, db *sql.DB, userUUID, accountUUID uuid.UUID) error {
	stmt := SELECT(Bool(true).AS("exists")).
		FROM(table.PaymentAccounts).
		WHERE(
			table.PaymentAccounts.UUID.EQ(UUID(accountUUID)).
				AND(table.PaymentAccounts.UserUUID.EQ(UUID(userUUID))),
		)

	var exists struct {
		exists bool
	}
	err := stmt.QueryContext(ctx, db, &exists)
	if err != nil {
		return err
	}

	return nil
}

func InsertAccount(ctx context.Context, db *sql.DB, account *platform.PaymentAccount) error {
	stmt := table.PaymentAccounts.INSERT(table.PaymentAccounts.MutableColumns).
		MODEL(account).
		RETURNING(table.PaymentAccounts.AllColumns)

	err := stmt.QueryContext(ctx, db, account)
	if err != nil {
		return err
	}

	return nil
}
