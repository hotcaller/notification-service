// Repository
package feedback

import (
    "context"
    "fmt"
    "time"

    "github.com/Arlandaren/pgxWrappy/pkg/postgres"
    "service/internal/domains/feedback/models"
)

type Repository struct {
    db *postgres.Wrapper
}

func NewRepository(db *postgres.Wrapper) *Repository {
    return &Repository{
        db: db,
    }
}

func (r *Repository) GetAllFeedback(ctx context.Context) ([]models.Feedback, error) {
    var feedbacks []models.Feedback
    query := `
        SELECT id, header, content, answer, user_id, created_at, answered_at
        FROM feedback
        ORDER BY created_at DESC
    `
    if err := r.db.Select(ctx, &feedbacks, query); err != nil {
        return nil, fmt.Errorf("failed to retrieve feedback: %w", err)
    }
    return feedbacks, nil
}

func (r *Repository) GetUserFeedback(ctx context.Context, userID int64) ([]models.Feedback, error) {
    var feedbacks []models.Feedback
    query := `
        SELECT id, header, content, answer, user_id, created_at, answered_at
        FROM feedback
        WHERE user_id = $1
        ORDER BY created_at DESC
    `
    if err := r.db.Select(ctx, &feedbacks, query, userID); err != nil {
        return nil, fmt.Errorf("failed to retrieve user feedback: %w", err)
    }
    return feedbacks, nil
}

func (r *Repository) CreateFeedback(ctx context.Context, feedback *models.Feedback) error {
    // Format current time as ISO string
    feedback.CreatedAt = time.Now().Format(time.RFC3339)
    
    query := `
        INSERT INTO feedback (header, content, user_id, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
    
    return r.db.QueryRow(ctx, query, 
        feedback.Header, 
        feedback.Content, 
        feedback.UserID, 
        feedback.CreatedAt).Scan(&feedback.ID)
}

func (r *Repository) GetFeedbackByID(ctx context.Context, id int64) (*models.Feedback, error) {
    var feedback models.Feedback
    query := `
        SELECT id, header, content, answer, user_id, created_at, answered_at
        FROM feedback
        WHERE id = $1
    `
    if err := r.db.Get(ctx, &feedback, query, id); err != nil {
        return nil, fmt.Errorf("failed to retrieve feedback with ID %d: %w", id, err)
    }
    return &feedback, nil
}

func (r *Repository) UpdateFeedbackAnswer(ctx context.Context, id int64, answer string) error {
    // Format current time as ISO string
    now := time.Now().Format(time.RFC3339)
    
    query := `
        UPDATE feedback
        SET answer = $1, answered_at = $2
        WHERE id = $3
    `
    _, err := r.db.Exec(ctx, query, answer, now, id)
    if err != nil {
        return fmt.Errorf("failed to update feedback answer: %w", err)
    }
    return nil
}