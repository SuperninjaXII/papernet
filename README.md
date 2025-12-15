# PAPERNET

papernet is a book downloading site written in go with htmx.it is wriiten for
students and everybody alike to use.

## what it offers is

- fast download
- fast experience
- beautiful user interface

# how to run 
#### requirements 
1. git
2. Deno
3. An editor neovim, vs code or any
4. Gcc 
5. Make


##### on pc

to run this on your own windows first install go lang and deno then use git to clone using the command below 


```
git clone https://github.com/One-Developerz/project-papernet.git
```

then in your terminal  while in the   `project-papernet` directory run

```
make install
```

Then for development 
```
make dev
```
##### on mobile 


For this you will need the application userland from playstore
> It's recommended to use the arch distribution 
Install dependencies
`pacman -Syu go deno git gcc make`

Then you can use GitHub cli to signup using auth.
[https://cli.github.com/](link to GitHub cli documentation)

Then, clone:

`git clone https://github.com/One-Developerz/project-papernet.git`

Then enter into directory and install. 
> Note : it will download somethings 

```
cd project-papner&&make install
```
 Then run it
 ```
 Make dev
```
 

# DEVELOPMENT 
1. **Code quality** 
	 - The naming of functions, variables and a sorts should specify the use , what it is doing in a simple manner.
	 - functions need to be 3 to 4 lines max 5.
	 - use tabs and not spaces
	 - do not use static numbers 
	 - add comments when necessary not everything.
	 - the use of AI should be to a lesser amount in short audit everything AI writes most of it get rewritten either way 
	 
2.**CSS rules**
	- use the units `svw` for width relative to screen width and `svh` for height relative to the screen height 
	- clear class names e.g .*nav-bar, download-button etc*
	- use of sass or less is okey just compile to CSS
	-

  

# DESIGN


### visual signals:
Every action taken needs visual ques 
It might be a subtitle color change , a growth in size or something micro but enables any users to know there actions are not being ignore or stack in case of low net work usage;
### speed:
The speed of the application needs to lightening that's why I used htmx,but this means a great deal of cooperation between back and front for cooperation.

### visual clarity:
Every position of anything needs to be in a reasonable way in a manner that it won't disturb the users actions, 
Only one font is needed 
A font hierarchy needs to made.

### color:
They is a file named `variables.css` in `public/CSS/` this files holds the colors only use variables from this, random color placement are not allowed.
This is for easy of maintenance and enhancements.
> Also follow the three color rule

# PaperNet Routes Documentation

## Overview
PaperNet implements two main routing groups:
- Public routes (`Routes`)
- Admin-only routes (`Admin`)

## Public Routes (`homeRoute.go`)

### HTML Endpoints

#### GET `/`
- Handler: `controllers.Home`
- Description: Renders the home page

#### GET `/search`
- Handler: `controllers.SearchPage`
- Description: Renders the search interface

#### POST `/SearchResults`
- Handler: `controllers.BooksPage`
- Description: Returns book results page after search

#### POST `/result`
- Handler: `controllers.SearchResult`
- Description: Processes and returns search results

#### GET `/download/:id`
- Handler: `controllers.DownloadPage`
- Description: Renders the download page for a specific book
- Parameters:
  - `id`: Book identifier

#### GET `/related/:tag`
- Handler: `controllers.GetBookByCartegoryBooks`
- Description: Displays books related by category/tag
- Parameters:
  - `tag`: Category identifier

### JSON API Endpoints

#### GET `/books`
- Handler: `controllers.GetAllBooks`
- Description: Returns a list of all books

#### GET `/book/:id`
- Handler: `controllers.GetBookByID`
- Description: Retrieves details for a specific book
- Parameters:
  - `id`: Book identifier

#### POST `/search`
- Handler: `controllers.SearchBooks`
- Description: Performs a book search and returns results as JSON

#### GET `/category/:tag`
- Handler: `controllers.GetBookByCartegory`
- Description: Retrieves books by category
- Parameters:
  - `tag`: Category identifier

## Admin Routes (`admin.go`)

All admin routes require JWT authentication obtained through the login endpoint.

### Endpoints

#### GET `/admin`
- Handler: `controllers.Admin`
- Description: Admin dashboard access
- Authentication: Required

#### POST `/admin/addBook`
- Handler: `controllers.CreateBook`
- Description: Adds a new book to the system
- Authentication: Required

#### POST `/authentication/createuser`
- Handler: `controllers.CreateUser`
- Description: Creates a new user account
- Authentication: Required
- here is an example 
```json
{
	"username": "JohnDoe",
	"email": "john.doe@example.com",
	"number": "1234567890",
	"password": "hello"
}
```

#### POST `/authentication/login`
- Handler: `controllers.Login`
- Description: Authenticates user and returns JWT token
- Authentication: Not required
- the Json needs to look like
```json
{
"email":"string",
"password": "string"
}
```
## Authentication

- Admin routes require a valid JWT token in the Authorization header
- JWT tokens are obtained through the `/authentication/login` endpoint
- Public routes do not require authentication

## Usage Notes

1. For public access, use endpoints defined in `Routes()`
2. For administrative tasks:
   - Login using `/authentication/login`
   - Use the returned JWT token in subsequent admin requests
   - Include token in Authorization header

## Error Handling
- Routes return appropriate HTTP 
- status codes
- JSON endpoints follow consistent error response format
- Authentication failures return 401 Unauthorized

## Security Considerations

1. Always validate JWT tokens for admin routes
2. Keep login credentialsystatus codes
- JSON endpoints follow consistent error response format
- Authentication failures return 401 Unauthorized

## Security Considerations

1. Always validate JWT tokens for admin routes
2. Keep login credentials secure
3. Use HTTPS in production
4. Implement rate limiting for authentication endpoints
5. Use local storage for keeping the jwt
