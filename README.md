# Http GET calls in Go
A base project for Http GET calls in Go language. 
In this project I am using **net/http** package to create a minimal web server and accept GET calls. Then for each call a HTML page is rendered with information.

## Go file explanation

#### Packages
I am importing the following packages:
- "html/template" : for working with HTML pages
- "log" : for logging
- "net/http" : for web server creation and endpoint handling
- "os" : for information on local machine, specifically here for finding working folder

#### Structs
I am going to use two structs. <br/> 
The first one is __Person__ which holds information for a person (Name, Surname, Age). <br/>
The second one is __Page__ which holds information we are using in our html page. (Massage, Persons array). <br/>

### Functions
| Name                | Description                                                                       |
|---------------------|-----------------------------------------------------------------------------------|
| **getWorkingDir**   | Returns the working directory                                                     |
| **main**            | Registers the two endpoints and also listening at port 8080 for incoming requests |
| **indexHandler**    | Handler for "/" endpoint, adds Message information on html page                   |
| **personsHandler**  | Handler for "/persons" endpoint, adds a list of Persons on htm page               |


## HTML file explanation
Remders a Message inside **`<h1>`** tags <br/>
In case the html was handled by **personsHandler**, the Persons list has values so the html renders the Person struct information



