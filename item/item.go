/* OpenBillsAPI - API for libre billing software OpenBills
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@vidhukant.xyz>

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package item

import (
	"github.com/MikunoNaka/OpenBillsAPI/brand"
)

/* each invoice must contain at least one item
 * you should be able to:
 * - add, modify, delete an item
 * - add item to invoice
 */

type Item struct {
	Unit        string // TODO: probably shouldn't be string
	Name        string
	Description string
	Brand       brand.Brand
	HSN         string

    // just the defaults, can be overridden in an invoice
	UnitPriceCurrency string
	UnitPrice         float64
	TaxPercentage     float64
	MaxQuantity       int // idk if quantities should be int or float
	MinQuantity       int
	Quantity          int // specified when creating invoice
}
