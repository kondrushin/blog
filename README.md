# blog

CRUD API (go, gin)

## API Endpoint specification

### Get a post by ID

The endpoint is designed to get a post by specifying its ID.

- **Endpoint URL:** "HTTP GET /v1/api/blog/posts/{id}"
- **Curl Command example:**
  ```
  curl -X GET 'http://localhost:8080/v1/api/blog/posts/1'
  ```
- **Response example:**
  ```json
  {
    "id": 1,
    "title": "Title 1",
    "content": "Content of the post",
    "author": "Author 1"
  }
  ```

### Get all posts

The endpoint is designed to retrieve all posts currently presented in the blog. The endpoint does not support sorting, filtration, pagenation.

- **Endpoint URL:** "HTTP GET /v1/api/blog/posts"
- **Curl Command example:**
  ```
  curl -X GET 'http://localhost:8080/v1/api/blog/posts'
  ```
- **Response example:**
  ```json
  {
    "posts": [
      {
        "Id": 1,
        "Author": "Anton",
        "Title": "On golang",
        "Content": "some content"
      },
      {
        "Id": 2,
        "Author": "Jonny",
        "Title": "On golang again",
        "Content": "some extra content"
      }
    ]
  }
  ```

### Create a new post in the blog

The endpoint is designed to add a new post in the blog. ID is granted automatically based on the next available value. It will be returned in the response body.

- **Endpoint URL:** "HTTP POST /v1/api/blog/posts"
- **Curl Command example:**
  ```
  curl -X POST 'http://localhost:8080/v1/api/blog/posts' \
    --header 'Content-Type: application/json' \
    --data '{
        "title": "On golang",
        "content": "some content",
        "author": "Anton"
        }'
  ```
- **Response example:**
  ```json
  {
    "Id": 3
  }
  ```

### Update post details

The endpoint is designed to update title, content and author values of an existing post in the blog by specifying its ID. ID will be returned in the response body.

- **Endpoint URL:** "HTTP PUT /v1/api/blog/posts/{id}"
- **Curl Command example:**
  ```
    curl -X PUT 'http://localhost:8080/v1/api/blog/posts/3' \
    --header 'Content-Type: application/json' \
    --data '{
        "Author": "Anton NEW",
        "Title": "On golang NEW",
        "Content": "some content NEW"
    }'
  ```
- **Response example:**
  ```json
  {
    "Id": 3
  }
  ```

### Delete a post from the blog

The endpoint is designed to delete a post from the blog by specifying its ID.

- **Endpoint URL:** "HTTP DELETE /v1/api/blog/posts/{id}"
- **Curl Command example:**
  ```
    curl -X DELETE 'http://localhost:8080/v1/api/blog/posts/2'
  ```
