package fakedata

import (
	"strconv"
	"strings"
)

var MobileNumberHeader = []string{
	"130",
	"131",
	"132",
	"133",
	"134",
	"135",
	"136",
	"137",
	"138",
	"139",

	"145",
	"146",
	"147",
	"148",
	"149",

	"150",
	"151",
	"152",
	"153",
	"155",
	"156",
	"157",
	"158",
	"159",

	"162",
	"165",
	"166",
	"167",

	"170",
	"171",
	"172",
	"173",
	"174",
	"175",
	"176",
	"177",
	"178",

	"180",
	"181",
	"182",
	"183",
	"184",
	"185",
	"186",
	"187",
	"188",
	"189",

	"190",
	"191",
	"192",
	"193",
	"195",
	"196",
	"197",
	"198",
	"199",
}

func MobileNumber() string {
	var sb strings.Builder
	header := MobileNumberHeader[globalRand.Intn(len(MobileNumberHeader))]
	sb.WriteString(header)
	for i := 0; i < 8; i++ {
		sb.WriteString(strconv.Itoa(globalRand.Intn(10)))
	}
	return sb.String()
}
