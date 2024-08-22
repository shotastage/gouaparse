# gouaparse

>> Currently, gouaparse is now under alpha and initial version!

gouaparse is a simple and lightweight user agent string parser written in Go. It parses user agent strings sent by web browsers and other HTTP clients, providing structured data for easy analysis.

## Features

- Simple and easy-to-use API
- Detection of major browsers, operating systems, and device types
- Identification of rendering engines
- Bot/crawler detection
- Lightweight with minimal dependencies

## Installation

To install gouaparse, run the following command:

```
go get github.com/shotastage/gouaparse
```

## Usage

Here's a basic example of how to use gouaparse:

```go
package main

import (
    "fmt"
    "github.com/shotastage/gouaparse"
)

func main() {
    uaString := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
    parsedUA := gouaparse.ParseUserAgent(uaString)

    fmt.Printf("Browser: %s %s\n", parsedUA.BrowserName, parsedUA.BrowserVersion)
    fmt.Printf("OS: %s\n", parsedUA.OperatingSystem)
    fmt.Printf("Device Type: %s\n", parsedUA.DeviceType)
    fmt.Printf("Rendering Engine: %s %s\n", parsedUA.RenderingEngine, parsedUA.RenderingEngineVersion)
    fmt.Printf("Is Bot: %v\n", parsedUA.IsBot)
}
```

## Contributing

We welcome contributions to the project. You can contribute in the following ways:

1. If you find a bug, please open a GitHub Issue.
2. If you have a suggestion for a new feature, create an Issue as well.
3. For code improvements or new feature implementations, please submit a Pull Request.

## License

This project is released under the MIT License. See the [LICENSE](LICENSE) file for details.

## Disclaimer

This parser is designed to provide basic functionality. It does not guarantee complete accuracy for all user agent strings. For advanced features or higher accuracy, consider using commercial solutions or more comprehensive libraries.

## Contact

If you have any questions or suggestions, please reach out through GitHub Issues.
