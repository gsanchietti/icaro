/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Icaro project.
 *
 * Icaro is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Icaro is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Icaro.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package models

import "time"

type HotspotVoucher struct {
	Id            int       `db:"id" json:"id"`
	HotspotId     int       `db:"hotspot_id" json:"hotspot_id"`
	Code          string    `db:"code" json:"code"`
	AutoLogin     bool      `db:"auto_login" json:"auto_login"`
	BandwidthUp   int       `db:"bandwidth_up" json:"bandwidth_up"`
	BandwidthDown int       `db:"bandwidth_down" json:"bandwidth_down"`
	Duration      int       `db:"duration" json:"duration"`
	Expires       time.Time `db:"expires" json:"expires"`
}
