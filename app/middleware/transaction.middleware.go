package middleware

import (
	"github.com/bantawao4/gofiber-boilerplate/config"
	"github.com/gofiber/fiber/v2"
)

const DBTransaction = "db_trx"

func DBTransactionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		txHandle := config.DB.Db.Begin()

		c.Locals(DBTransaction, txHandle)

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
				panic(r)
			}
		}()

		err := c.Next()

		if err != nil {
			txHandle.Rollback()
			return err
		}

		if c.Response().StatusCode() >= 200 && c.Response().StatusCode() < 300 {
			if err := txHandle.Commit().Error; err != nil {
				txHandle.Rollback()
				return err
			}
		} else {
			txHandle.Rollback()
		}

		return nil
	}
}
