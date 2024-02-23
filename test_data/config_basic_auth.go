package test_data

var (
	IsBasicAuthStr = []testData{
		{
			Name: "Invalid_missing_username_and_password",
			InputStr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{false},
		},
		{
			Name: "Invalid_missing_username",
			InputStr: []string{
				"",
				"pass",
			},
			ExpectedBool: []bool{false},
		},
		{
			Name: "Invalid_missing_password",
			InputStr: []string{
				"user",
				"",
			},
			ExpectedBool: []bool{false},
		},
		{
			Name: "Valid",
			InputStr: []string{
				"user",
				"pass",
			},
			ExpectedBool: []bool{true},
		},
	}
)
