# Exchange Authorization Code for Access Token

This tool exchanges your authorization code for an actual access token.

## Quick Steps

### 1. Get Your Credentials

Go to https://www.strava.com/settings/api and note:
- **Client ID** (e.g., `12345`)
- **Client Secret** (e.g., `abc123def456...`)

### 2. Edit main.go

Open `main.go` and replace:
```go
clientID := "YOUR_CLIENT_ID"        // Put your actual Client ID
clientSecret := "YOUR_CLIENT_SECRET" // Put your actual Client Secret
```

### 3. Get Authorization Code

If you don't have the code yet:

1. Replace YOUR_CLIENT_ID in this URL and open in browser:
```
https://www.strava.com/oauth/authorize?client_id=YOUR_CLIENT_ID&response_type=code&redirect_uri=http://localhost:8080/callback&approval_prompt=force&scope=read,activity:read_all
```

2. Authorize the app
3. You'll see an error page with a URL like:
```
http://localhost:8080/callback?state=&code=XXXXXXXXXXXX&scope=...
```

4. Copy the `code=XXXXXXXXXXXX` part

### 4. Update the Code

In `main.go`, update:
```go
authorizationCode := "paste_your_code_here"
```

### 5. Run

```bash
go run main.go
```

### 6. Use the Access Token

Copy the access token from the output and use it:

```bash
export STRAVA_ACCESS_TOKEN="the_token_from_output"
cd ../get_activities
go run main.go
```

## Important Notes

⚠️ **Authorization codes are single-use and expire in 10 minutes!**

If you get an error:
- The code might have already been used
- The code might have expired
- Your Client ID/Secret might be wrong

Just get a new authorization code and try again.
