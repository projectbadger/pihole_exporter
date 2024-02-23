package test_data

var (
	PiholeClient = []testData{
		{
			Name: "Empty_default",
			InputStr: []string{
				"",
			},
			ExpectedStr: []string{"", ""},
		},
	}
	MetricsURL = []testData{
		{
			Name: "Empty_missing_listen",
			InputStr: []string{
				"", // ListenAddress
				"", // APIPath
				"", // CACertificate
				"", // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr:  []string{"", ""},
			ExpectedBool: []bool{true, true}, // Expect error
		},
		{
			Name: "Valid_domain",
			InputStr: []string{
				"pi.hole", // ListenAddress
				"",        // APIPath
				"",        // CACertificate
				"",        // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_apiPath",
			InputStr: []string{
				"pi.hole",    // ListenAddress
				"/test/path", // APIPath
				"",           // CACertificate
				"",           // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/test/path?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_apiPath_ca",
			InputStr: []string{
				"pi.hole",     // ListenAddress
				"/test/path",  // APIPath
				"ca.cert.pem", // CACertificate
				"",            // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/test/path?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_ca",
			InputStr: []string{
				"pi.hole",     // ListenAddress
				"",            // APIPath
				"ca.cert.pem", // CACertificate
				"",            // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_insecure",
			InputStr: []string{
				"pi.hole",     // ListenAddress
				"",            // APIPath
				"ca.cert.pem", // CACertificate
				"",            // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_apiPath_insecure",
			InputStr: []string{
				"pi.hole",     // ListenAddress
				"/test/path",  // APIPath
				"ca.cert.pem", // CACertificate
				"",            // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/test/path?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Invalid_apiPath_ca",
			InputStr: []string{
				"",            // ListenAddress
				"/test/path",  // APIPath
				"ca.cert.pem", // CACertificate
				"",            // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr:  []string{"", ""},
			ExpectedBool: []bool{true, true}, // Expect error
		},
		{
			Name: "Invalid_apiPath_ca_insecure",
			InputStr: []string{
				"",            // ListenAddress
				"/test/path",  // APIPath
				"ca.cert.pem", // CACertificate
				"",            // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr:  []string{"", ""},
			ExpectedBool: []bool{true, true}, // Expect error
		},
		{
			Name: "Invalid_insecure",
			InputStr: []string{
				"", // ListenAddress
				"", // APIPath
				"", // CACertificate
				"", // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr:  []string{"", ""},
			ExpectedBool: []bool{true, true}, // Expect error
		},
		// http/https
		{
			Name: "Valid_http_domain_insecure",
			InputStr: []string{
				"http://pi.hole", // ListenAddress
				"",               // APIPath
				"",               // CACertificate
				"",               // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_http_domain_insecure_ca",
			InputStr: []string{
				"http://pi.hole", // ListenAddress
				"",               // APIPath
				"ca.cert.pem",    // CACertificate
				"",               // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_https_domain_ca",
			InputStr: []string{
				"https://pi.hole", // ListenAddress
				"",                // APIPath
				"ca.cert.pem",     // CACertificate
				"",                // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_https_domain",
			InputStr: []string{
				"https://pi.hole", // ListenAddress
				"",                // APIPath
				"",                // CACertificate
				"",                // APIToken
			},
			InputInt: []int{
				0, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		// NumResults
		{
			Name: "Valid_https_domain_numResults",
			InputStr: []string{
				"https://pi.hole", // ListenAddress
				"",                // APIPath
				"",                // CACertificate
				"",                // APIToken
			},
			InputInt: []int{
				15, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?getForwardDestinations=15&getQuerySources=15&getQueryTypes=15&jsonForceObject=15&overTimeData=15&recentItems=15&summaryRaw=15&topItems=15",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_https_domain_numResults",
			InputStr: []string{
				"https://pi.hole", // ListenAddress
				"",                // APIPath
				"",                // CACertificate
				"",                // APIToken
			},
			InputInt: []int{
				-1, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?getForwardDestinations=30&getQuerySources=30&getQueryTypes=30&jsonForceObject=30&overTimeData=30&recentItems=30&summaryRaw=30&topItems=30",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_http_domain_numResults",
			InputStr: []string{
				"http://pi.hole", // ListenAddress
				"",               // APIPath
				"",               // CACertificate
				"",               // APIToken
			},
			InputInt: []int{
				1, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?getForwardDestinations=1&getQuerySources=1&getQueryTypes=1&jsonForceObject=1&overTimeData=1&recentItems=1&summaryRaw=1&topItems=1",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_http_domain_numResults_apiToken",
			InputStr: []string{
				"http://pi.hole", // ListenAddress
				"",               // APIPath
				"",               // CACertificate
				"token",          // APIToken
			},
			InputInt: []int{
				1, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?auth=token&getForwardDestinations=1&getQuerySources=1&getQueryTypes=1&jsonForceObject=1&overTimeData=1&recentItems=1&summaryRaw=1&topItems=1",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_http_domain_numResults_apiToken",
			InputStr: []string{
				"http://pi.hole", // ListenAddress
				"",               // APIPath
				"",               // CACertificate
				"1",              // APIToken
			},
			InputInt: []int{
				1, // NumResults
			},
			InputBool: []bool{
				false, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?auth=1&getForwardDestinations=1&getQuerySources=1&getQueryTypes=1&jsonForceObject=1&overTimeData=1&recentItems=1&summaryRaw=1&topItems=1",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_http_domain_insecure_numResults_apiToken",
			InputStr: []string{
				"http://pi.hole", // ListenAddress
				"",               // APIPath
				"",               // CACertificate
				"token",          // APIToken
			},
			InputInt: []int{
				15, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"http://pi.hole/admin/api.php?auth=token&getForwardDestinations=15&getQuerySources=15&getQueryTypes=15&jsonForceObject=15&overTimeData=15&recentItems=15&summaryRaw=15&topItems=15",
				"",
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_insecure_apiToken_basic_auth",
			InputStr: []string{
				"pi.hole", // ListenAddress
				"",        // APIPath
				"",        // CACertificate
				"token",   // APIToken,
				"user",    // Username for basic auth
				"pass",    // Password for basic auth
			},
			InputInt: []int{
				15, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?auth=token&getForwardDestinations=15&getQuerySources=15&getQueryTypes=15&jsonForceObject=15&overTimeData=15&recentItems=15&summaryRaw=15&topItems=15",
				"Basic dXNlcjpwYXNz", // Authorization header
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
		{
			Name: "Valid_domain_insecure_apiToken_basic_auth",
			InputStr: []string{
				"pi.hole", // ListenAddress
				"",        // APIPath
				"",        // CACertificate
				"token",   // APIToken,
				"",        // Username for basic auth
				"",        // Password for basic auth
			},
			InputInt: []int{
				15, // NumResults
			},
			InputBool: []bool{
				true, // Insecure
			},
			ExpectedStr: []string{
				"https://pi.hole/admin/api.php?auth=token&getForwardDestinations=15&getQuerySources=15&getQueryTypes=15&jsonForceObject=15&overTimeData=15&recentItems=15&summaryRaw=15&topItems=15",
				"", // Authorization header
			},
			ExpectedBool: []bool{false, false}, // Expect error
		},
	}
)
