package repositories

import (
	"context"
	"database/sql"
	"log"

	"CyberMate_Back/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) models.UserRepository {
	return &userRepo{pool: pool}
}

func (r *userRepo) Upsert(ctx context.Context, chatID int64) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO users (chat_id) VALUES ($1) ON CONFLICT (chat_id) DO NOTHING`,
		chatID)
	if err != nil {
		log.Printf("❌ Upsert error for chat_id %d: %v", chatID, err)
		return err
	}
	return nil
}

func (r *userRepo) Get(ctx context.Context, chatID int64) (*models.User, error) {
	var u models.User

	// 🔥 Убрали notifications_enabled из запроса
	err := r.pool.QueryRow(ctx,
		`SELECT chat_id, language FROM users WHERE chat_id = $1`,
		chatID).Scan(&u.ChatID, &u.Language)

	if err != nil {
		if err == sql.ErrNoRows {
			return &models.User{
				ChatID:   chatID,
				Language: "ru",
			}, nil
		}
		log.Printf("❌ Get user error for chat_id %d: %v", chatID, err)
		return nil, err
	}

	return &u, nil
}

func (r *userRepo) UpdateLanguage(ctx context.Context, chatID int64, lang string) error {
	if lang != "ru" && lang != "en" {
		log.Printf("⚠️  Invalid language '%s' for chat_id %d", lang, chatID)
		lang = "ru"
	}

	result, err := r.pool.Exec(ctx,
		`UPDATE users SET language = $1, updated_at = NOW() WHERE chat_id = $2`,
		lang, chatID)

	if err != nil {
		log.Printf("❌ UpdateLanguage error for chat_id %d: %v", chatID, err)
		return err
	}

	if result.RowsAffected() == 0 {
		log.Printf("⚠️  No user found to update language for chat_id %d", chatID)
	}

	return nil
}
