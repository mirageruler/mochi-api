package producttheme

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Get() ([]model.ProductTheme, error)
}
