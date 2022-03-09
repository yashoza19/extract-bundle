package bundle

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/yashoza19/extract-bundles/pkg/models"
)

type Data struct {
	ExtractBundle []models.ExtractBundle
	Flags         Flags
}

func BuildBundlesQuery() (string, error) {
	query := sq.Select("o.name, o.bundlepath").From(
		"operatorbundle o")

	query.OrderBy("o.name")

	sql, _, err := query.ToSql()
	if err != nil {
		return "", fmt.Errorf("unable to create sql : %s", err)
	}
	return sql, nil
}
