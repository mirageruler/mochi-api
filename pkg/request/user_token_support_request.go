package request

type CreateUserTokenSupportRequest struct {
	UserDiscordID string `json:"user_discord_id,omitempty" binding:"required"`
	ChannelID     string `json:"channel_id,omitempty" binding:"required"`
	MessageID     string `json:"message_id,omitempty" binding:"required"`
	TokenName     string `json:"token_name,omitempty" binding:"required"`
	TokenAddress  string `json:"token_address,omitempty" binding:"required"`
	TokenChain    string `json:"token_chain,omitempty" binding:"required"`
}