// Service
package feedback

import (
    "context"
    "errors"
    "service/internal/domains/feedback/models"
)

var (
    ErrFeedbackNotFound = errors.New("feedback not found")
    ErrUnauthorized     = errors.New("unauthorized access")
)

type Service struct {
    repo *Repository
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) ListAllFeedback(ctx context.Context) ([]models.Feedback, error) {
    return s.repo.GetAllFeedback(ctx)
}

func (s *Service) ListUserFeedback(ctx context.Context, userID int64) ([]models.Feedback, error) {
    return s.repo.GetUserFeedback(ctx, userID)
}

func (s *Service) SendFeedback(ctx context.Context, feedback models.Feedback) error {
    return s.repo.CreateFeedback(ctx, &feedback)
}

func (s *Service) AnswerFeedback(ctx context.Context, id int64, answer string) error {
    // Check if feedback exists
    feedback, err := s.repo.GetFeedbackByID(ctx, id)
    if err != nil {
        return ErrFeedbackNotFound
    }
    
    if feedback == nil {
        return ErrFeedbackNotFound
    }
    
    return s.repo.UpdateFeedbackAnswer(ctx, id, answer)
}