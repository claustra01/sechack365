package repository

import (
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
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
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS "user.username",
		users.protocol AS "user.protocol",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
		FROM articles
		JOIN users ON articles.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE articles.id = $1;
	`
	if err := r.SqlHandler.Get(article, query, id); err != nil {
		return nil, cerror.Wrap(err, "failed to get article by id")
	}
	return article, nil
}

func (r *ArticleRepository) CreateArticleComment(articleId, userId, content string) error {
	var count int
	query := "SELECT COUNT(*) FROM article_comments WHERE article_id = $1 AND user_id = $2 AND content = $3;"
	if err := r.SqlHandler.Get(&count, query, articleId, userId, content); err != nil {
		return cerror.Wrap(err, "failed to check article comment")
	}
	if count > 0 {
		return nil // already exists
	}

	uuid := util.NewUuid()
	query = `
		INSERT INTO article_comments (id, article_id, user_id, content)
		VALUES ($1, $2, $3, $4);
	`
	if _, err := r.SqlHandler.Exec(query, uuid, articleId, userId, content); err != nil {
		return cerror.Wrap(err, "failed to create article comment")
	}
	return nil
}

func (r *ArticleRepository) FindCommentsByArticleId(articleId string) ([]*model.ArticleCommentWithUser, error) {
	var comments []*model.ArticleCommentWithUser
	query := `
		SELECT
		article_comments.id,
		article_comments.content,
		article_comments.created_at,
		article_comments.updated_at,
		users.id AS "user.id",
		CASE
			WHEN users.protocol = 'local' THEN '@' || users.username
			WHEN users.protocol = 'activitypub' THEN '@' || ap_user_identifiers.local_username || '@' || ap_user_identifiers.host
			WHEN users.protocol = 'nostr' THEN nostr_user_identifiers.npub
		END AS "user.username",
		users.protocol AS "user.protocol",
		users.display_name AS "user.display_name",
		users.icon AS "user.icon"
		FROM article_comments
		JOIN users ON article_comments.user_id = users.id
		LEFT JOIN ap_user_identifiers ON users.id = ap_user_identifiers.user_id
		LEFT JOIN nostr_user_identifiers ON users.id = nostr_user_identifiers.user_id
		WHERE article_comments.article_id = $1
		ORDER BY article_comments.created_at DESC;
	`
	if err := r.SqlHandler.Select(&comments, query, articleId); err != nil {
		return nil, cerror.Wrap(err, "failed to get comments by article id")
	}
	return comments, nil
}

func (r *ArticleRepository) CreateArticlePostRelation(articleId, postId string) error {
	query := `
		INSERT INTO article_post_relations (article_id, post_id)
		VALUES ($1, $2);
	`
	if _, err := r.SqlHandler.Exec(query, articleId, postId); err != nil {
		return cerror.Wrap(err, "failed to create article post relation")
	}
	return nil
}

func (r *ArticleRepository) FindArticlePostRelation(postId string) (*model.ArticlePostRelation, error) {
	relation := new(model.ArticlePostRelation)
	query := "SELECT * FROM article_post_relations WHERE post_id = $1;"
	if err := r.SqlHandler.Get(relation, query, postId); err != nil {
		return nil, cerror.Wrap(err, "failed to get article post relation")
	}
	return relation, nil
}
