package response

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
)

type User struct {
	ID                     string                  `json:"id"`
	Username               string                  `json:"username"`
	InDiscordWalletAddress *string                 `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  *int64                  `json:"in_discord_wallet_number"`
	GuildUsers             []*GetGuildUserResponse `json:"guild_users"`
}

type GetGuildUserResponse struct {
	GuildID   string `json:"guild_id"`
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	InvitedBy string `json:"invited_by"`
}

type HandleUserActivityResponse struct {
	ChannelID    string    `json:"channel_id"`
	GuildID      string    `json:"guild_id"`
	UserID       string    `json:"user_id"`
	Action       string    `json:"action"`
	AddedXP      int       `json:"added_xp"`
	CurrentXP    int       `json:"current_xp"`
	CurrentLevel int       `json:"current_level"`
	Timestamp    time.Time `json:"timestamp"`
	LevelUp      bool      `json:"level_up"`
}

type TopUser struct {
	Author      *model.GuildUserXP  `json:"author"`
	Leaderboard []model.GuildUserXP `json:"leaderboard"`
}

type GetUserProfileResponse struct {
	ID           string               `json:"id"`
	AboutMe      string               `json:"about_me"`
	CurrentLevel *model.ConfigXpLevel `json:"current_level"`
	NextLevel    *model.ConfigXpLevel `json:"next_level"`
	GuildXP      int                  `json:"guild_xp"`
	NrOfActions  int                  `json:"nr_of_actions"`
	Progress     float64              `json:"progress"`
	Guild        *model.DiscordGuild  `json:"guild"`
	GuildRank    int                  `json:"guild_rank"`
	UserWallet   *model.UserWallet    `json:"user_wallet"`
}

// For swagger

type GetMyInfoResponse struct {
	Data *discordgo.User `json:"data"`
}

type GetUserResponse struct {
	Data User `json:"data"`
}

type GetTopUsersResponse struct {
	Data TopUser `json:"data"`
}
type GetUserCurrentGMStreakResponse struct {
	Data *model.DiscordUserGMStreak `json:"data"`
}

type GetUserCurrentUpvoteStreakResponse struct {
	UserID         string    `json:"discord_id"`
	ResetTime      float64   `json:"minutes_until_reset"`
	SteakCount     int       `json:"streak_count"`
	TotalCount     int       `json:"total_count"`
	LastStreakTime time.Time `json:"last_streak_time"`
}
