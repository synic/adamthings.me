package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/synic/blog/internal/model"
	"github.com/synic/blog/internal/repo"
	"github.com/synic/blog/internal/store"
)

var (
	ErrNotFound = repo.ErrNotFound
)

type CreateCommentParams = store.CreateCommentParams

type ArticleService struct {
	repo    *repo.ArticleRepository
	queries *store.Queries
}

func NewArticleService(queries *store.Queries, repo *repo.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo, queries: queries}
}

func (s *ArticleService) TagInfo(ctx context.Context) map[string]int {
	return s.repo.TagInfo(ctx)
}

func (s *ArticleService) Tags(ctx context.Context) []string {
	return s.repo.Tags(ctx)
}

func (s *ArticleService) FindAll(ctx context.Context) ([]*model.Article, error) {
	return s.repo.FindAll(ctx)
}

func (s *ArticleService) Search(ctx context.Context, terms string) ([]*model.Article, error) {
	return s.repo.Search(ctx, terms)
}

func (s *ArticleService) FindOneBySlug(ctx context.Context, slug string) (*model.Article, error) {
	article, err := s.repo.FindOneBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) FindAllTags(ctx context.Context) ([]string, error) {
	return s.repo.FindAllTags(ctx)
}

func (s *ArticleService) FindByTag(ctx context.Context, tag string) ([]*model.Article, error) {
	return s.repo.FindByTag(ctx, tag)
}

func (s *ArticleService) CreateComment(
	ctx context.Context,
	params CreateCommentParams,
) (*store.Comment, error) {
	if params.ApprovalCode == "" {
		code, err := randomHex(32)

		if err != nil {
			return nil, err
		}
		params.ApprovalCode = code
	}

	comment, err := s.queries.CreateComment(ctx, params)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *ArticleService) FindCommentsForArticle(
	ctx context.Context,
	slug string,
) ([]store.FindCommentsForArticleRow, error) {
	return s.queries.FindCommentsForArticle(ctx, slug)
}

func (s *ArticleService) SendCommentNotifications(
	article *model.Article,
	user store.User,
	comment store.Comment,
) error {
	username := os.Getenv("SMTP_USERNAME")
	pass := os.Getenv("SMTP_PASSWORD")
	adminEmail := os.Getenv("ADMIN_EMAIL")

	if username == "" || pass == "" || adminEmail == "" {
		err := errors.New("Cannot send email, credentials not set")
		log.Printf("error sending comment approval email: %v", err)
		return err
	}

	auth := smtp.PlainAuth("", username, pass, "smtp.gmail.com")
	to := []string{adminEmail}

	subject := fmt.Sprintf(
		"New comment on blog article %s from %s",
		article.Slug,
		user.Email,
	)
	content := fmt.Sprintf(
		"Slug: %s\nEmail: %s\nName: %s\nContent: %s\n\nApproval URL: %s",
		article.Slug,
		user.Email,
		user.Name,
		comment.Content,
		comment.ApprovalURL(),
	)

	msg := []byte(
		fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", strings.Join(to, ", "), subject, content),
	)

	err := smtp.SendMail("smtp.gmail.com:587", auth, "no-reply@blog", to, msg)

	if err != nil {
		log.Printf("error sending comment approval email: %v", err)
	}

	return err
}

func (s *ArticleService) ApproveComment(
	ctx context.Context,
	approvalCode string,
) (store.Comment, error) {
	return s.queries.ApproveComment(ctx, approvalCode)
}

func randomHex(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
