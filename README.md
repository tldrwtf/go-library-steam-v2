### Example Usage and Responses for All Functions

#### `Login`
**Usage:**
```go
err := bot.Login("username", "password")
if err != nil {
    log.Fatalf("Login failed: %v", err)
}
log.Println("Login successful")
```
**Response:**
```
Login successful
```

#### `GenerateCode`
**Usage:**
```go
code := steamGuard.GenerateCode()
log.Printf("Steam Guard Code: %s", code)
```
**Response:**
```
Steam Guard Code: ABCDE
```

#### `GenerateConfirmationKey`
**Usage:**
```go
confirmationKey := steamGuard.GenerateConfirmationKey("tag")
log.Printf("Confirmation Key: %s", confirmationKey)
```
**Response:**
```
Confirmation Key: dummy_confirmation_key
```

#### `makeAPIRequest`
**Usage:**
```go
params := map[string]string{"key": "your_api_key", "steamid": "123456789"}
response, err := bot.makeAPIRequest("ISteamUser/GetPlayerSummaries/v2", params)
if err != nil {
    log.Fatalf("API Request failed: %v", err)
}
log.Printf("API Response: %s", response)
```
**Response:**
```
API Response: {"response":{"players":[{"steamid":"123456789","personaname":"John Doe"}]}}
```

#### `makeAPIPostRequest`
**Usage:**
```go
data := map[string]string{"key": "your_api_key", "steamid": "123456789"}
response, err := bot.makeAPIPostRequest("ISteamUser/AddFriend/v1", data)
if err != nil {
    log.Fatalf("API POST Request failed: %v", err)
}
log.Printf("API POST Response: %s", response)
```
**Response:**
```
API POST Response: {"response":"success"}
```

#### `AddFriend`
**Usage:**
```go
err := bot.AddFriend("someSteamID")
if err != nil {
    log.Fatalf("Error adding friend: %v", err)
}
log.Println("Friend added successfully")
```
**Response:**
```
Friend added successfully
```

#### `RemoveFriend`
**Usage:**
```go
err := bot.RemoveFriend("someSteamID")
if err != nil {
    log.Fatalf("Error removing friend: %v", err)
}
log.Println("Friend removed successfully")
```
**Response:**
```
Friend removed successfully
```

#### `AcceptFriendRequest`
**Usage:**
```go
err := bot.AcceptFriendRequest("someSteamID")
if err != nil {
    log.Fatalf("Error accepting friend request: %v", err)
}
log.Println("Friend request accepted successfully")
```
**Response:**
```
Friend request accepted successfully
```

#### `SendTradeOffer`
**Usage:**
```go
items := []steam.Item{
    {ID: "1", Name: "Item1", Price: 10.0},
    {ID: "2", Name: "Item2", Price: 20.0},
}
err := bot.SendTradeOffer("someSteamID", items)
if err != nil {
    log.Fatalf("Error sending trade offer: %v", err)
}
log.Println("Trade offer sent successfully")
```
**Response:**
```
Trade offer sent successfully
```

#### `ListItemsOnMarket`
**Usage:**
```go
items := []steam.Item{
    {ID: "1", Name: "Item1", Price: 10.0},
    {ID: "2", Name: "Item2", Price: 20.0},
}
err := bot.ListItemsOnMarket(items)
if err != nil {
    log.Fatalf("Error listing items on market: %v", err)
}
log.Println("Items listed on market successfully")
```
**Response:**
```
Items listed on market successfully
```

#### `GetInventory`
**Usage:**
```go
inventory, err := bot.GetInventory("someSteamID")
if err != nil {
    log.Fatalf("Error retrieving inventory: %v", err)
}
log.Printf("Inventory: %+v", inventory)
```
**Response:**
```
Inventory: {Items:[{ID:1 Name:Item1 Price:10} {ID:2 Name:Item2 Price:20}]}
```

#### `GetUserData`
**Usage:**
```go
userData, err := bot.GetUserData("someSteamID")
if err != nil {
    log.Fatalf("Error retrieving user data: %v", err)
}
log.Printf("User Data: %+v", userData)
```
**Response:**
```
User Data: {SteamID:123456789 PersonaName:John Doe ProfileURL:https://steamcommunity.com/id/johndoe Avatar:https://avatar.url AvatarMedium:https://avatar.url/medium AvatarFull:https://avatar.url/full}
```

#### `HandleCaptchaChallenge`
**Usage:**
```go
solution, err := bot.HandleCaptchaChallenge("captchaURL")
if err != nil {
    log.Fatalf("CAPTCHA handling failed: %v", err)
}
log.Printf("CAPTCHA solution: %s", solution)
```
**Response:**
```
CAPTCHA solution: dummy_captcha_solution
```

#### `SubmitCaptchaSolution`
**Usage:**
```go
err := bot.SubmitCaptchaSolution("captchaID", "dummy_captcha_solution")
if err != nil {
    log.Fatalf("CAPTCHA submission failed: %v", err)
}
log.Println("CAPTCHA solution submitted successfully")
```
**Response:**
```
CAPTCHA solution submitted successfully
```
