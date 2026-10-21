package expense

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPool struct {
	db *pgxpool.Pool
}

func NewPostgresStore(db *pgxpool.Pool) *PostgresPool {
	return &PostgresPool{db: db}
}

func (p *PostgresPool) Add(ctx context.Context, value Expense) (Expense, error) {
	spentOn, err := time.Parse(time.DateOnly, value.SpentOn)
	if err != nil {
		return Expense{}, err
	}

	const query = `
		INSERT INTO expenses (categoryid, amount, spenton, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at	
	`
	err = p.db.QueryRow(ctx, query, value.CategoryID, value.Amount, spentOn, value.Description).Scan(&value.ID, &value.CreatedAt)

	if err != nil {
		return Expense{}, err
	}

	return value, nil
}

func (p *PostgresPool) Delete(ctx context.Context, id int64) (bool, error) {
	const query = `
		DELETE FROM expenses WHERE id = $1
	`
	tag, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}
	if affected := tag.RowsAffected(); affected == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (p *PostgresPool) List(ctx context.Context) ([]Expense, error) {
	const query = `
		SELECT id, categoryid, amount, spenton, description, created_at
		FROM expenses
		ORDER BY id
	`
	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]Expense, 0)

	for rows.Next() {
		var value Expense
		var spentOn time.Time

		err := rows.Scan(
			&value.ID,
			&value.CategoryID,
			&value.Amount,
			&spentOn,
			&value.Description,
			&value.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		value.SpentOn = spentOn.Format(time.DateOnly)
		result = append(result, value)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
