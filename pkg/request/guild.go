package request

import "time"

type CreateGuildRequest struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	JoinedAt time.Time `json:"-"`
}

type UpdateGuildRequest struct {
	GlobalXP   *bool      `json:"global_xp"`
	LogChannel *string    `json:"log_channel"`
	Active     *bool      `json:"active"`
	LeftAt     *time.Time `json:"left_at"`
}

type HandleGuildDeleteRequest struct {
	GuildID   string `json:"guild_id"`
	GuildName string `json:"guild_name"`
	IconURL   string `json:"icon_url"`
}

type ValidateUserRequest struct {
	Ids     string `form:"ids"`
	GuildId string `form:"guild_id"`
}
