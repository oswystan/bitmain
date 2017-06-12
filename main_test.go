//===============================================================================
//                      Copyright (C) 2017 wystan
//
//        filename: main_test.go
//     description:
//         created: 2017-06-12 15:25:04
//          author: wystan
//
//===============================================================================

package main_test

import (
	"net/http"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var root = "http://localhost:8000"

func TestSpec(t *testing.T) {
	//student features
	Convey("==>students", t, func() {
		Convey("post", func() {
			data := `
				{
					"id":"10000",
					"classNumber":1,
					"score": 100
				}`
			res, err := http.Post(root+"/register-student", "application/json", strings.NewReader(data))
			So(err, ShouldEqual, nil)
			defer res.Body.Close()
			So(res.StatusCode, ShouldEqual, http.StatusOK)
		})
		Convey("post-update", func() {
			data := `
				{
					"id":"10000",
					"classNumber":1,
					"score": 90
				}`
			res, err := http.Post(root+"/register-student", "application/json", strings.NewReader(data))
			So(err, ShouldEqual, nil)
			defer res.Body.Close()
			So(res.StatusCode, ShouldEqual, http.StatusOK)
		})
	})
	Convey("==>classes", t, func() {
		Convey("post", func() {
			data := `
				{
					"classNumber":1,
					"teacher": "jason born"
				}`
			res, err := http.Post(root+"/register-class", "application/json", strings.NewReader(data))
			So(err, ShouldEqual, nil)
			defer res.Body.Close()
			So(res.StatusCode, ShouldEqual, http.StatusOK)
		})
		Convey("post-update", func() {
			data := `
				{
					"classNumber":1,
					"teacher": "Lucy Li"
				}`
			res, err := http.Post(root+"/register-class", "application/json", strings.NewReader(data))
			So(err, ShouldEqual, nil)
			defer res.Body.Close()
			So(res.StatusCode, ShouldEqual, http.StatusOK)
		})
	})

	Convey("===>get-top-score", t, func() {
		Convey("normal case", func() {
			res, err := http.Get(root + "/get-class-total-score/10000")
			So(err, ShouldEqual, nil)
			defer res.Body.Close()
			So(res.StatusCode, ShouldEqual, http.StatusOK)
		})
		Convey("abnormal case", func() {
			res, err := http.Get(root + "/get-class-total-score/1000")
			So(err, ShouldEqual, nil)
			defer res.Body.Close()
			So(res.StatusCode, ShouldNotEqual, http.StatusOK)
		})
	})
}

//==================================== END ======================================
