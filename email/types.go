package email

type Auth struct {
	identity string
	username string
	password string
	hostname string
}

type Server struct {
	name string
	port int
}

type Sender struct {
	addr string
}

type Recipients struct {
	addr []string
}

type LoginAuth struct {
	username string
	password string
}

type Message struct {
	subject string
	body    string
	msg     []byte
}

// @dev The keys must be written in Capitalized case.
type Env struct {
	EmailUsername  string `mapstructure:"EMAIL_USERNAME"`
	EmailPassword  string `mapstructure:"EMAIL_PASSWORD"`
	EmailRecipient string `mapstructure:"EMAIL_RECIPIENT"`
}
