package repository

import (
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
)

type ArticleRepository struct {
	SqlHandler model.ISqlHandler
}

func (r *ArticleRepository) Create(id, userId, title, content string) error {
	query := `
		INSERT INTO articles (id, user_id, title, content)
		VALUES ($1, $2, $3, $4);
	`
	if _, err := r.SqlHandler.Exec(query, id, userId, title, content); err != nil {
		return cerror.Wrap(err, "failed to create article")
	}
	return nil
}

func (r *ArticleRepository) FindById(id string) (*model.ArticleWithUser, error) {
	article := new(model.ArticleWithUser)
	query := `
		SELECT
		articles.id,
		articles.title,
		articles.content,
		articles.created_at,
		articles.updated_at,
		users.id AS "user.id",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
		FROM articles
		JOIN users ON articles.user_id = users.id
		WHERE articles.id = $1;
	`
	if err := r.SqlHandler.Get(article, query, id); err != nil {
		return nil, cerror.Wrap(err, "failed to get article by id")
	}
	return article, nil
}

func (r *ArticleRepository) FindCommentsByArticleId(articleId string) ([]*model.ArticleCommentWithUser, error) {
	var comments []*model.ArticleCommentWithUser
	query := `
		SELECT article_comments.*,
		users.id AS "user.id",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
		FROM article_comments
		JOIN users ON article_comments.user_id = users.id
		WHERE article_comments.article_id = $1
		ORDER BY article_comments.created_at DESC;
	`
	if err := r.SqlHandler.Select(&comments, query, articleId); err != nil {
		return nil, cerror.Wrap(err, "failed to get comments by article id")
	}
	return comments, nil
}
