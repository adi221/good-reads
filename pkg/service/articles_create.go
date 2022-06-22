package service

import (
	"context"

	"github.com/adi221/good-reads/pkg/model"
)

func (reg *Registry) CreateArticle(ctx context.Context, form model.ArticleCreateForm) (*model.Article, error) {
	// uid := getCurrentUserIDFromContext(ctx)
	// For now
	var uid uint = 1232122

	var category *model.Category
	if form.CategoryID != nil {
		// TODO: check if category actually exists later on
		category = &model.Category{}
	}

	if category == nil {
		reg.logger.Error().Msg("Category ID must be provided")
	}

	if form.URL != nil && !form.IsComplete() {
		// Fetch original article in order to extract missing attributes
		if err := reg.scrapOriginalArticle(ctx, &form); err != nil {
			reg.logger.Info().Err(err).Uint(
				"uid", uid,
			).Str("title", form.TruncatedTitle()).Msg("unable to fetch original article")
		}
	}

	// Sanitize HTML content
	if form.HTML != nil {
		content := reg.sanitizer.Sanitize(*form.HTML)
		form.HTML = &content
	}

	reg.logger.Debug().Uint(
		"uid", uid,
	).Str("title", form.TruncatedTitle()).Msg("creating article...")
	article, err := reg.db.CreateArticleForUser(uid, form)
	if err != nil {
		reg.logger.Info().Err(err).Uint(
			"uid", uid,
		).Str("title", form.TruncatedTitle()).Msg("unable to create article")
		return nil, err
	}
	reg.logger.Info().Uint(
		"uid", uid,
	).Str("title", form.TruncatedTitle()).Uint("id", article.ID).Msg("article created")

	return article, nil
}

// scrapOriginalArticle add missing attributes form original article
func (reg *Registry) scrapOriginalArticle(ctx context.Context, article *model.ArticleCreateForm) error {
	page, err := reg.webScraper.Scrap(ctx, *article.URL)
	if page == nil {
		return err
	}
	article.URL = &page.URL
	if article.Title == "" {
		article.Title = page.Title
	}
	if article.HTML == nil {
		article.HTML = &page.HTML
	}
	if article.Text == nil {
		article.Text = &page.Text
	}
	if article.Image == nil {
		article.Image = &page.Image
	}

	return err
}
