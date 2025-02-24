package db_driver

const (
	queryFindAccount = `
	SELECT 
	    account_id,
		account_pin,
		account_balance::numeric
	FROM accounts
	WHERE account_number = $1
`
	queryGetBalance = `
	SELECT account_balance::numeric
	FROM accounts
	WHERE account_number = $1
`
	queryUpdateBalance = `
	UPDATE accounts
	SET account_balance = $2
	WHERE account_number = $1;
`
	queryAddTransaction = `
	INSERT INTO transactions(transaction_id, account_id, transaction_type, transaction_value, transaction_date)
	VALUES ($2, (SELECT id FROM accounts WHERE account_number = $1), $3, $4, $5)
`
	queryGetTransactions = `
	SELECT 
		transaction_id, 
		transaction_type,
		transaction_value, 
		transaction_date
	FROM transactions
	WHERE account_id = (SELECT id FROM accounts WHERE account_number = $1)
`
	queryCreateAccount = `
	INSERT INTO accounts(account_id, account_number, account_pin, account_balance)
	Values ($1, $2, $3, $4)
`
	queryDeleteAccount = `
	DELETE FROM accounts
	WHERE account_number = $1
`
	queryFindAdmin = `
	SELECT admin_id
	FROM admins
	WHERE admin_number = $1 and admin_password = $2
	`
)
