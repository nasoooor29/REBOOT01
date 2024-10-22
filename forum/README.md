# FORUM

## Description

This project involves creating a web-based forum that facilitates communication
between users, categorization of posts, and interaction through likes and
dislikes. The forum also features post filtering and user authentication.

## Features

- Post Creation: Registered users can create posts and associate them with one
  or more categories.
- Commenting: Registered users can comment on posts.
- Visibility: Posts and comments are visible to all users, including
  non-registered visitors.
- Users can associate categories with posts, allowing organization and filtering
  of content.
- Interaction: Registered users can like or dislike posts and comments.
- Categories: Filter posts by category.

## User Authentication

- Users can register by providing an email, username, and password.
- Passwords are encrypted.
- Users can log in using their credentials.
- Use cookies to manage user sessions with an expiration date.

## Technology Stack

### Backend

- Golang standard packages along with sqlite3, bcrypt and UUID.
- Javascript
- SQLite: Store and manage user data, posts, comments, likes, dislikes, and categories.
- Docker: Containerize the application for consistent development and deployment
  environments.

### Frontend

HTML & CSS (No frontend libraries or frameworks)

## Usage

Run the build.sh script by executing the below command in a terminal

```bash
bash ./build.sh
```

Once the build is completed successfully, verify if the container is running:

```docker
docker ps -a
```

Open a web browser of your choice and navigate to the below link:

`http://localhost:8080/`

## Authors

- [nhussain](https://learn.reboot01.com/git/nhussain)
- [yabuzuha](https://learn.reboot01.com/git/yabuzuha)
- [etarada](https://learn.reboot01.com/git/etarada)
- [fsayedsa](https://learn.reboot01.com/git/fsayedsa)
