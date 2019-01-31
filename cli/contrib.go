// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*  This file handles third party contributions and libraries.*/
package main

import (
	"net/http"

	dataset "github.com/cvpm-contrib/dataset"
	"github.com/gin-gonic/gin"
)

// GET /datasets
func GetAllDatasets(c *gin.Context) {
	c.JSON(http.StatusOK, dataset.FetchAllDatasets())
}
