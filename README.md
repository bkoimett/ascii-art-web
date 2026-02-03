ascii-art-web/
├── main.go
├── go.mod
├── go.sum
├── README.MD
├── run_tests.sh
├── coverage.html                 # Generated coverage report
├── coverage.out                  # Generated coverage data
│
├── asciiart/                     # Your existing ASCII art logic (with tests)
│   ├── asciiart.go
│   ├── asciiart_test.go          # Your existing unit tests
│   └── banner_loader.go
│
├── handlers/                     # HTTP handlers
│   ├── handlers.go
│   └── handlers_test.go          # Handler unit tests
│
├── templates/                    # HTML templates
│   ├── index.html
│   ├── result.html
│   ├── error.html
│   └── templates_test.go         # Template validation tests
│
├── static/                       # Static assets
│   └── css/
│       └── style.css
│
├── banner/                       # Banner files (not in git)
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── integration/                  # Integration tests
│   └── integration_test.go
│
├── edgecases/                    # Edge case tests
│   └── edgecases_test.go
│
├── load/                         # Load/concurrent tests
│   └── load_test.go
│
├── testdata/                     # Test data files
│   └── test_banners/
│       ├── standard_test.txt
│       ├── shadow_test.txt
│       └── thinkertoy_test.txt
│
└── scripts/                      # Helper scripts (optional)
    └── setup_test_env.sh