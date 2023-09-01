# Ascii-Art-Web-Stylize
🎨 Ascii-Art-Web-Stylize is an exciting project that merges ASCII art generation with a user-friendly web interface. It builds upon the foundation of the original ascii-art project, elevating it with various captivating banner styles. This project harnesses a Go-based web server to craft a harmonious user experience.


## Description

Ascii-Art-Web-Stylize empowers users to effortlessly generate ASCII art banners via a web interface. It offers a selection of banner styles, including shadow, standard, and thinkertoy. Users can easily input their desired text, choose a banner style, and with a single click, create visually appealing ASCII art.

## Authors

🚀 **Captain:** 
- **Hatham Radhi**
  - Reboot01 ID: [@hradhi](https://learn.reboot01.com/hradhi)

🌟 **Team Member:** 
- **Amir Iqbal**
  - Reboot01 ID: [@aiqbal](https://learn.reboot01.com/aiqbal)



## Usage

To run the Ascii-Art-Web-Stylize application:

1. Clone this repository to your local machine.
2. Make sure you have Go installed.
3. Navigate to the project directory in your terminal.
4. Run the following command to start the application:
```go
go run .
```


5. Open your web browser and access the Ascii-Art-Web-Stylize interface at http://localhost:8080.

## Implementation Details
### ASCII Art Generation Algorithm

The ASCII art generation process follows these steps:

1. **Banner Format**:
   - Each character is represented by an 8-line pattern in the banner style.
   - Characters are separated by new lines (\n).

2. **User Input Processing**:
   - The application takes user input, which includes the text and the chosen banner style.

3. **Loading Banner Styles**:
   - The chosen banner style's ASCII art template is loaded from the corresponding text file in the "banners" directory.

4. **Text to ASCII Art Conversion**:
   - The input text is split into lines, where each line represents a line of ASCII art.
   - For each line of text:
     - The line is split into individual characters.
     - For each character:
       - The character's ASCII value is checked to ensure it falls within the Basic Latin character range (ASCII values 32 to 126).
       - If the character is within this range, its corresponding pattern from the banner style's ASCII art template is retrieved.
       - The patterns of all characters in the line are concatenated to form the ASCII art representation for that line.
       - The process is repeated for each of the 8 lines in the character's pattern.
       - The application maps each character to its corresponding pattern in the banner style, resulting in the ASCII art representation for each line of text.

5. **Rendering the Result**:
   - The generated ASCII art, along with the original input text and chosen banner style, is rendered on the result page for the user to view.


### Web Application Framework

Ascii-Art-Web is built on the Go programming language and the net/http package. It consists of the following main components:

- **Main Page Handler**: Renders the main page where users input text and select a banner style.

- **ASCII Art Handler**: Processes user input, generates ASCII art, and renders the result page.

- **Not Found Handler**: Handles cases where users access non-existent routes.

## Extensible and Scalable

The application is designed to be extensible, allowing for easy addition of new banner styles by creating corresponding text files in the "banners" directory. 
- To add custome Banner File make sure:
  - Banner Format

    - Each character has a height of 8 lines.
    - Characters are separated by a new line \n.

  Here is an example of  '!' (one dot represents one space) :

```sh
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....
```
The flexible nature of Go's web framework enables further feature enhancements and improvements.
 