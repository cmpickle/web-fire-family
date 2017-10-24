package app

import "testing"

func TestCanary(t *testing.T){
		if (!(true == true)) {
		t.error("Expected True got somthing doesn't exist")
	}

}