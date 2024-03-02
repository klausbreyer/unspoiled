
# Unspoiled

Welcome to Unspoiled, an open-source project that helps users find the latest video for a specified YouTube channel. This Go-based web server makes it easy to retrieve and display information about the most recent video posted by your favorite YouTube content creators.

## Features

- Fetch the latest video from a specified YouTube channel
- Display the video link and publish date
- Easy-to-use web interface with responsive design

## Getting Started

### Prerequisites

- Go (at least version 1.15)
- YouTube API Key

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/unspoiled.git
   ```
2. Navigate to the project directory:
   ```
   cd unspoiled
   ```
3. Set your YouTube API Key as an environment variable:
   ```
   export YOUTUBE_API_KEY=your_api_key_here
   ```
4. Run the server:
   ```
   go run .
   ```
5. Open your browser and visit `http://localhost:8080` to start using Unspoiled.

## Usage

To find the latest video from a YouTube channel, navigate to `http://localhost:8080/latest?channelId=your_channel_id_here` in your web browser, replacing `your_channel_id_here` with the ID of the YouTube channel you're interested in.

For example, to find the latest video from the Unspoiled example channel:
```
http://localhost:8080/latest?channelId=UCmaItsxNPLEQ-NIjv5gPScg
```

You can also visit the root URL (`http://localhost:8080`) for instructions and an example link.

## Contributing

We welcome contributions to Unspoiled! Please refer to our contribution guidelines for more information.

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- Thanks to everyone who has contributed to this project!
- Special thanks to the YouTube Data API for making this possible.
