# strip-html-go
This is a basic web UI (using html and javascript) with Golang/Go Programming.
# What it does
This project:
1. Shows a simple HTML UI with some javascript in it
2. Grabs the content of URL provided by user in the url field
3. Loads the data from the url
4. Strips all HTML tag from the url content, so that only the real-content is left
5. Gets the count of all words from the remaining string
6. Presents the words with their counts in the grid manner
# Running the project
1. Please make sure the project follows the correct path order (`github.com/rohsa/strip-html-go`)
2. Please make sure you're not using/blocking the port `8080` on your localhost
2. `cd` to the project directory and run `go build main.go`
3. Open your web browser and hit the url `www.localhost:8080`
