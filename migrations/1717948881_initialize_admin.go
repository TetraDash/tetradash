package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		admin := &models.Admin{}
		admin.Email = "admin@tetradash.com"
		admin.SetPassword("password")

		return dao.SaveAdmin(admin)
	}, func(db dbx.Builder) error {

		dao := daos.New(db)

		admin, err := dao.FindAdminByEmail("admin@tetradash.com")
		if admin != nil {
			return dao.DeleteAdmin(admin)
		}

		// already deleted
		return err
	})
}
