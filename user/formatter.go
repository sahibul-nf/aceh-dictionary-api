package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{}
	formatter.ID = user.ID
	formatter.Name = user.Name
	formatter.Email = user.Email
	formatter.AvatarURL = user.AvatarURL

	return formatter
}

type AuthUser struct {
	Token string `json:"token"`
}

func FormatAuthUser(token string) AuthUser {
	formatter := AuthUser{}
	formatter.Token = token

	return formatter
}
