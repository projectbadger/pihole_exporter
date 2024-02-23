package test_data

var (
	PiholeAPIPath = []testData{
		{
			Name: "Empty_default",
			InputStr: []string{
				"",
			},
			ExpectedStr: []string{"/admin/api.php"},
		},
		{
			Name: "Set",
			InputStr: []string{
				"/some/path",
			},
			ExpectedStr: []string{"/some/path"},
		},
	}
	PiholeNumResults = []testData{
		{
			Name: "Empty_default",
			InputInt: []int{
				0,
			},
			ExpectedInt: []int{30},
		},
		{
			Name: "Not_empty_default",
			InputInt: []int{
				30,
			},
			ExpectedInt: []int{30},
		},
		{
			Name: "Not_empty_invalid",
			InputInt: []int{
				-20,
			},
			ExpectedInt: []int{30},
		},
		{
			Name: "Not_empty_valid",
			InputInt: []int{
				20,
			},
			ExpectedInt: []int{20},
		},
		{
			Name: "Not_empty_valid_2",
			InputInt: []int{
				1,
			},
			ExpectedInt: []int{1},
		},
		{
			Name: "Not_empty_valid_3",
			InputInt: []int{
				100,
			},
			ExpectedInt: []int{100},
		},
	}

	PiholeRepliesMap = []testData{
		{
			Name: "Empty_default",
			InputStr: []string{
				"NODATA",   // NODATA
				"NXDOMAIN", // NXDOMAIN
				"CNAME",    // CNAME
				"IP",       // IP
				"Domain",   // Domain
				"RRNAME",   // RRNAME
				"SERVFAIL", // SERVFAIL
				"REFUSED",  // REFUSED
				"OTHER",    // OTHER
				"DNSSEC",   // DNSSEC
				"NONE",     // NONE
				"BLOB",     // BLOB
			},
			ExpectedStr: []string{
				"NODATA",
				"NXDOMAIN",
				"CNAME",
				"IP",
				"Domain",
				"RRNAME",
				"SERVFAIL",
				"REFUSED",
				"OTHER",
				"DNSSEC",
				"NONE",
				"BLOB",
			},
		},
	}
)
