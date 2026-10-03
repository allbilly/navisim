package runner

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCollectorDump(t *testing.T) {
	dir, err := ioutil.TempDir("", "navisim-metrics-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	c := collector{}
	c.Collect("driver", "kernel_time", 0.000005352)
	name := filepath.Join(dir, "metrics")
	c.Dump(name)

	data, err := ioutil.ReadFile(name + ".csv")
	if err != nil {
		t.Fatal(err)
	}
	want := ", where, what, value\n0, driver, kernel_time, 0.000005352000\n"
	if string(data) != want {
		t.Fatalf("metrics file = %q, want %q", data, want)
	}
}
