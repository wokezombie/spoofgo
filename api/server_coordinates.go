package api

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"net/http"
)

func (self *Server) Coordinates(w http.ResponseWriter, r *http.Request) {
	self.APICoordinatesRespond(nil, w, r, 200)
}
