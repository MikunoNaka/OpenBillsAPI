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

package invoice

import (
	"github.com/MikunoNaka/OpenBills/client"
)

/* you should be able to:
 * - add, modify, delete an invoice
 * - add client to invoice
 * - add items to invoice
 */


/* NOTE: not sure if I should store vehicle numbers
 * one vehicle may be owned by two transporters
 * like maybe it's rented, etc
 */
// maybe store transporter details in DB
type Transporter struct {
	Id            int
	Name          string
	GSTIN         string
	TransporterId string
}

// transport details
type Transport struct {
	Transporter Transporter
	VerhicleNum string
	Note        string
}

// TODO: check if IGST/CGST Apply
// TODO: add place of supply
// TODO: store payment method and related details
// TODO: if International then store details like currency
type Invoice struct {
	Id            int // not the same as invoice number
	InvoiceNumber int
	Recipient     client.Client

	// TODO: handle dates probably shouldn't be stored as string
	CreatedAt     string
	LastUpdated   string
	Paid          bool

	International bool

	ShippingAddress client.Address
	Transport     Transport
}
