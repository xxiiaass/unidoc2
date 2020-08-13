// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package zippkg_test

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/xxiiaass/unidoc2/zippkg"
)

type TestStruct struct {
	Foo string
}

func TestMarshal(t *testing.T) {
	buf := &bytes.Buffer{}
	zw := zip.NewWriter(buf)
	f := TestStruct{Foo: "bar"}

	fname := "/test/foo.xml"
	if err := zippkg.MarshalXML(zw, fname, &f); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if err := zw.Close(); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if len(zr.File) != 1 {
		t.Errorf("expected one file in zip, got %d", len(zr.File))
	}
	zf := zr.File[0]
	if zf.Name != fname {
		t.Errorf("expected name = %s, got %s", fname, zf.Name)
	}
	rc, err := zf.Open()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	out, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	exp := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n" + `<TestStruct><Foo>bar</Foo></TestStruct>` + "\r\n"
	if got := string(out); got != exp {
		t.Errorf("expected\n%s\n, got \n%s\n", exp, got)
	}
}
