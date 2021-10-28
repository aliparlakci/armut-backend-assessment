# Armut Backend Assessment

## Notes

**Cache control** headers are missing in the project. They can be implemented to have more control over the content and bandwith optimization.

# API Reference

## Models

### Message
- **id** string
- **to** string
- **from** string
- **body** string
- **send_at** string
- **is_read** bool

### Activity
- **id** string
- **event** "signin" | "signout" | "fail_signin"
- **username** string
- **ip** string
- **when** string

## Endpoints

All the endpoints return **HTTP 401 Status Unauthorized** if the endpoint requires authorization and request does not have `session` cookie or the provided one does not exist.

### GET /api/messages
Returns all the messages (send or received) of the user. Needs authorization.
  
Returns **HTTP 200** if successful. Return type is `Message[]`.

### GET /api/messages/new
Returns only the unread received messages of the user. Needs authorization. 

Returns **HTTP 200** if successful. Return type is `Message[]`.

### GET /api/messages/check
Returns the number of unread received messages of the user. Return type is number. Needs authorization.

Returns **HTTP 200** if successful. Return type is `Message[]`.

### POST /api/messages/send
Sends a message to a user. Needs authorization.

- Content-Type: Multipart Form
- Fields:
  - **to**: Username of the receiver
  - **body**: Message body
  
Returns **HTTP 201** if successful. Returns **HTTP 400** if either of the fields are missing or provided username does not belong to a user.
  
### PUT /api/messages/:messageId/read
Marks the message with messageId read. Message needs be received by the logged in user. Needs authorization.

Returns **HTTP 200** if successful. Returns **HTTP 400** if messageId is missing or it does not correspond to a message.

### PUT /api/messages/user/:username/read
Marks the messages whose sender is the **username**. Message needs be received by the logged-in user. Needs authorization.

Returns **HTTP 200** if successful. Returns **HTTP 400** if username is missing.

### POST /api/signup
Sends a message to a user. Needs authorization.

- Content-Type: **Multipart Form**
- Fields:
    - **username**
    - **password**

Returns **HTTP 201** if successful.

### POST /api/signin
Creates a new session for the user and sets session id as `session` cookie. `session` cookie must not exist on the request.
  
- Content-Type: **Multipart Form**
- Fields:
    - **username**
    - **password**

Returns **HTTP 200** if successful. Returns **HTTP 400** if user another user already logged-in or either of the fields are missing or username and password mismatch.

### POST /api/signout
Revokes the session of the user and unsets the `session` cookie.

Returns **HTTP 200** if successful. Returns **HTTP 400** if no one is signed in or `session` does not correspond to a session.

### GET /api/me
Returns the username of the user which is signed in on the provided session. Needs authorization.

Returns **HTTP 200** if successful.

### GET /api/activity
Returns the authorization activity of the current user. Needs authorization.

Returns **HTTP 200** if successful. Return type is `Activity[]`