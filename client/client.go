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

package client

/* each invoice has a client
 * you should be able to:
 * - add, modify, delete a client
 * - add client to invoices
 * - get all invoices associated with client, etc
 */

/* each contact has one name
 * but multiple contact addresses
 * it is assumed that the first one
 * has the highest priority
 */
// TODO: contact should have an "other" field where key and values
// can be specified i.e Twitter: @ClientTwitterAccount
type Contact struct {
	Name    string
	Phone   []string
	Email   []string
	Website string
}

type Address struct {
	Contact Contact
	Address string
	City    string
	State   string
	ZIPCode string
	Country string
}

type Client struct {
	Id      int
	Name    string
	Contact Contact
	GSTIN   string

	BillingAddress    Address
	ShippingAddresses []Address
}
