package corviper

import (
	"flag"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestStringValueLoad(t *testing.T) {
	os.Args = []string{"cmd", "-param1", "someNewValue", "-param3", "someOtherValue"}
	sflag := NewStringFlag("param1", "testValue", "some test", "my.test")
	sflag2 := NewStringFlag("param2", "", "some test2", "my.test2")
	sflag3 := NewStringFlag("param3", "testValue3", "some test3", "my.test3")

	viper.BindFlagValue(sflag.Name(), sflag)
	viper.BindFlagValue(sflag2.Name(), sflag2)
	viper.BindFlagValue(sflag3.Name(), sflag3)

	flag.Parse()

	// sflag
	if sflag.HasChanged() == false {
		t.Error("sflag.HasChanged() == false")
	}
	if sflag.Name() != "my.test" {
		t.Error(`sflag.Name() != "my.test"`)
	}
	if sflag.ValueString() != "someNewValue" {
		t.Error(`sflag.ValueString() != "someNewValue"`)
	}

	// sflag2
	if sflag2.HasChanged() == true {
		t.Error("sflag2.HasChanged() == true")
	}
	if sflag2.Name() != "my.test2" {
		t.Error(`sflag2.Name() != "my.test2"`)
	}
	if sflag2.ValueString() != "" {
		t.Error(`sflag2.ValueString() != ""`)
	}

	// sflag3
	if sflag3.HasChanged() == false {
		t.Error("sflag3.HasChanged() == false")
	}
	if sflag3.Name() != "my.test3" {
		t.Error(`sflag3.Name() != "my.test3"`)
	}
	if sflag3.ValueString() != "someOtherValue" {
		t.Error(`sflag3.ValueString() != "someOtherValue"`)
	}

	if viper.GetString("my.test") != "someNewValue" {
		t.Error(`my.test != "someNewValue"`)
	}
	if viper.GetString("my.test2") != "" {
		t.Error(`my.test2 != ""`)
	}
	if viper.GetString("my.test3") != "someOtherValue" {
		t.Error(`my.test3 != "someOtherValue"`)
	}
}
