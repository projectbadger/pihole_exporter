package test_data

var (
	TLSIsSet = []testData{
		{
			Name: "Not_set",
			InputStr: []string{
				"", // CACertificate
				"", // Certificate
				"", // Key
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate",
			InputStr: []string{
				"",           // CACertificate
				"./cert.crt", // Certificate
				"",           // Key
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key",
			InputStr: []string{
				"",          // CACertificate
				"",          // Certificate
				"./key.pem", // Key
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_only_CA",
			InputStr: []string{
				"./ca.cert.pem", // CACertificate
				"",              // Certificate
				"",              // Key
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_insecure",
			InputStr: []string{
				"", // CACertificate
				"", // Certificate
				"", // Key
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"",         // Key
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_key_and_insecure",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"key.pem", // Key
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
	}

	WebTLSIsSet = []testData{
		{
			Name: "Not_set",
			InputStr: []string{
				"", // CACertificate
				"", // Certificate
				"", // Key
				"", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate",
			InputStr: []string{
				"",           // CACertificate
				"./cert.crt", // Certificate
				"",           // Key
				"",           // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key",
			InputStr: []string{
				"",          // CACertificate
				"",          // Certificate
				"./key.pem", // Key
				"",          // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_CA",
			InputStr: []string{
				"./ca.cert.pem", // CACertificate
				"",              // Certificate
				"",              // Key
				"",              // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_insecure",
			InputStr: []string{
				"", // CACertificate
				"", // Certificate
				"", // Key
				"", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_ca_and_cert",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"",         // Key
				"",         // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_ca_and_cert_and_key",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"",         // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"",         // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Not_set_only_key_and_insecure",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"key.pem", //
				"",        // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_listen_domain",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"",        // Key
				"pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate_listen_domain",
			InputStr: []string{
				"",           // CACertificate
				"./cert.crt", // Certificate
				"",           // Key
				"pi.hole",    // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key_listen_domain",
			InputStr: []string{
				"",          // CACertificate
				"",          // Certificate
				"./key.pem", // Key
				"pi.hole",   // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_CA_listen_domain",
			InputStr: []string{
				"./ca.cert.pem", // CACertificate
				"",              // Certificate
				"",              // Key
				"pi.hole",       // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_insecure_listen_domain",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"",        // Key
				"pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_ca_and_cert_listen_domain",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"",         // Key
				"pi.hole",  // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_ca_and_cert_and_key_listen_domain",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"pi.hole",  // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure_listen_domain",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"pi.hole",  // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Not_set_only_key_and_insecure_listen_domain",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"key.pem", //
				"pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"./cert.crt",      // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"./key.pem",       // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_CA_listen_domain_https",
			InputStr: []string{
				"./ca.cert.pem",   // CACertificate
				"",                // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_insecure_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_ca_and_cert_listen_domain_https",
			InputStr: []string{
				"./ca.crt",        // CACertificate
				"cert.crt",        // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_ca_and_cert_and_key_listen_domain_https",
			InputStr: []string{
				"./ca.crt",        // CACertificate
				"cert.crt",        // Certificate
				"key.pem",         // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure_listen_domain_https",
			InputStr: []string{
				"./ca.crt",        // CACertificate
				"cert.crt",        // Certificate
				"key.pem",         // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Not_set_only_key_and_insecure_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"key.pem",         //
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{false},
		},
	}

	PiholeTLSIsSet = []testData{
		{
			Name: "Not_set",
			InputStr: []string{
				"", // CACertificate
				"", // Certificate
				"", // Key
				"", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate",
			InputStr: []string{
				"",           // CACertificate
				"./cert.crt", // Certificate
				"",           // Key
				"",           // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key",
			InputStr: []string{
				"",          // CACertificate
				"",          // Certificate
				"./key.pem", // Key
				"",          // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_only_CA",
			InputStr: []string{
				"./ca.cert.pem", // CACertificate
				"",              // Certificate
				"",              // Key
				"",              // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_insecure",
			InputStr: []string{
				"", // CACertificate
				"", // Certificate
				"", // Key
				"", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"",         // Key
				"",         // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"",         // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"",         // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_key_and_insecure",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"key.pem", //
				"",        // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true}, // Because of insecure
		},
		{
			Name: "Not_set_listen_domain",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"",        // Key
				"pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate_listen_domain",
			InputStr: []string{
				"",           // CACertificate
				"./cert.crt", // Certificate
				"",           // Key
				"pi.hole",    // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key_listen_domain",
			InputStr: []string{
				"",          // CACertificate
				"",          // Certificate
				"./key.pem", // Key
				"pi.hole",   // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_only_CA_listen_domain",
			InputStr: []string{
				"./ca.cert.pem", // CACertificate
				"",              // Certificate
				"",              // Key
				"pi.hole",       // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_insecure_listen_domain",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"",        // Key
				"pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_listen_domain",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"",         // Key
				"pi.hole",  // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_listen_domain",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"pi.hole",  // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure_listen_domain",
			InputStr: []string{
				"./ca.crt", // CACertificate
				"cert.crt", // Certificate
				"key.pem",  // Key
				"pi.hole",  // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_key_and_insecure_listen_domain",
			InputStr: []string{
				"",        // CACertificate
				"",        // Certificate
				"key.pem", //
				"pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Not_set_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_certificate_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"./cert.crt",      // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Not_set_only_key_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"./key.pem",       // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{false},
		},
		{
			Name: "Set_only_CA_listen_domain_https",
			InputStr: []string{
				"./ca.cert.pem",   // CACertificate
				"",                // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_insecure_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_listen_domain_https",
			InputStr: []string{
				"./ca.crt",        // CACertificate
				"cert.crt",        // Certificate
				"",                // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_listen_domain_https",
			InputStr: []string{
				"./ca.crt",        // CACertificate
				"cert.crt",        // Certificate
				"key.pem",         // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{false}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_ca_and_cert_and_key_and_insecure_listen_domain_https",
			InputStr: []string{
				"./ca.crt",        // CACertificate
				"cert.crt",        // Certificate
				"key.pem",         // Key
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
		{
			Name: "Set_only_key_and_insecure_listen_domain_https",
			InputStr: []string{
				"",                // CACertificate
				"",                // Certificate
				"key.pem",         //
				"https://pi.hole", // Listen
			},
			InputBool:    []bool{true}, // Insecure
			ExpectedBool: []bool{true},
		},
	}
)
