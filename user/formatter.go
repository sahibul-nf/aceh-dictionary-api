package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	AvatarURL string `json:"avatar_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{}
	formatter.ID = user.ID
	formatter.Name = user.Name
	formatter.Email = user.Email
	formatter.Token = token
	formatter.AvatarURL = user.AvatarURL

	return formatter
}
