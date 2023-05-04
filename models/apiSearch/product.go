package models

import (
	"api_fiber_warehouse/initializers"
	modelsStruc "api_fiber_warehouse/models/struc/search"
	"log"

	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
)

func Prod(c *fiber.Ctx) []modelsStruc.Searchprod {
	var ks modelsStruc.Keyword
	var listsearch []modelsStruc.Searchprod
	if err := c.BodyParser(&ks); err != nil {
		return listsearch
	}
	var keysearch = ks.Keysearch

	if utf8.RuneCountInString(keysearch) > 2 {
		var err error
		sql_statement := "SELECT id_product AS 'value', CONCAT( key_search ) AS title, IFNULL(sale_price_per_unit,0) AS sale_price_per_unit, IFNULL(id_supplier,0) AS id_supplier, IFNULL(id_prod_category,0) AS  id_prod_category, IFNULL(qty_per_pack,0) AS qty_per_pack, IFNULL(unit_name,'') AS unit_name FROM product WHERE key_search LIKE '%" + keysearch + "%'"
		list, err := initializers.DB.Query(sql_statement)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer list.Close()
		for list.Next() {
			var row modelsStruc.Searchprod
			err = list.Scan(
				&row.Value,
				&row.Title,
				&row.Sale_price_per_unit,
				&row.Id_supplier,
				&row.Id_prod_category,
				&row.Qty_per_pack,
				&row.Unit_name,
			)
			if err != nil {
				log.Fatal(err.Error())
			}
			listsearch = append(listsearch, row)
		}
		err = list.Err()
		if err != nil {
			log.Fatal(err.Error())
		}
		return listsearch
	}
	return listsearch
}
