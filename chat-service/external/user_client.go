package external

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type UserInfo struct {
    ID        int64  `json:"id"`
    Username  string `json:"username"`
    AvatarURL string `json:"avatar_url"`
}

func GetUserByID2(id int64) (*UserInfo, error) {
    client := http.Client{Timeout: 3 * time.Second}
    url := fmt.Sprintf("http://auth-service:8080/api/users/%d", id)

    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("auth-service returned status %d", resp.StatusCode)
    }

    var user UserInfo
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        return nil, err
    }

    return &user, nil
}
