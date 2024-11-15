# Spixel

**Spixel** (Spy + Pixel) is a simple server built in Golang designed to host a spy pixel. This tracking pixel can be used to detect whether an email or a web page was loaded by the recipient, letting you know when it was read or viewed.

## Features

- **Track Email Opens**: Generate a unique tracking pixel for each target to monitor email opens or web page views.
- **Simple Setup**: Lightweight and easy to deploy with minimal dependencies.
- **Logs Access Information**: Records details of each access, such as timestamp, IP address, and user-agent, if desired.

## Getting Started

### Prerequisites

- **Go** installed (version 1.16 or higher)

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/spixel.git
   cd spixel
   ```

2. Build the server:

   ```bash
   go build -o spixel
   ```

3. Run the server:
   ```bash
   ./spixel
   ```

### Usage

1. Generate a Target: To create a unique target. This will provide you with a unique ID representing the target.

2. Insert the Pixel in an Email or Web Page: Embed the tracking pixel by adding an <img> tag pointing to your server, including the unique ID in the URL:
   ```html
   <img
     src="http://yourserver.com/ID"
     alt="tracking pixel"
     style="display:none;"
   />
   ```
   Each time the pixel is loaded, Spixel will log the access, confirming that the email or page was viewed.

### License

This project is licensed under the MIT License.

### Disclaimer

Ensure compliance with privacy laws when using tracking pixels. Always obtain consent before tracking email opens.
