package users

type user struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

func All() []user {
	var users []user
	users = append(users, user{
		First: "Jason",
		Last:  "McCallister",
	})
	users = append(users, user{
		First: "Another",
		Last:  "McCallister",
	})
	return users
}

func New(f, l string) user {
	return user{
		First: f,
		Last:  l,
	}
}
