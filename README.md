# blog

CRUD API (go, gin)

## API Endpoint specification

### Get a post by ID

- **Endpoint URL:** "HTTP GET /v1/api/blog/posts/{id}"
- **Curl Command example:**
  ```
  curl --location 'http://localhost:8080/v1/api/blog/posts/1'
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

- **Endpoint URL:** "HTTP GET /v1/api/blog/posts"
- **Curl Command example:**
  ```
  curl --location 'http://localhost:8080/v1/api/blog/posts'
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
