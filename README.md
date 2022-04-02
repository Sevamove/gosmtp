# Send Microsoft Outlook email through SMTP in Go

## Used Libraries
1. [Viper](github.com/spf13/viper)

## Installation

Clone the repo:

```bash
$ git clone https://github.com/vsevdrob/
```

## Prerequisites

1. After the `git clone` change directory and create `secret.env`:

```bash
$ cd go-smtp
$ touch secret.env
```

2. In `secret.env` file, insert the outlook email address from which you are going to send emails. Also provide the password of that email (account) address. Then provide an email address of the recipient.

```bash
EMAIL_USERNAME="<account_name>@<domain_name>"
EMAIL_PASSWORD="UnCrACKa6LE-Pa$$W00RD"
EMAIL_RECIPIENT="<account_name>@<domain_name>"
```

(optional) In `./email/types.go` append additional keys:

```go
...

type Env struct {
    EmailUsername  string `mapstructure:"EMAIL_USERNAME"`
    EmailPassword  string `mapstructure:"EMAIL_PASSWORD"`
    EmailRecipient string `mapstructure:"EMAIL_RECIPIENT"`

    // Append the rest here from that ./secret.env
}
```

## Usage

Run commando after your `secret.env` is set up:

```bash
$ go run main.go
```

## Resources

1. Go [SMTP Standard Library](https://pkg.go.dev/net/smtp) documentation
1. [POP, IMAP, and SMTP](https://support.microsoft.com/en-us/office/pop-imap-and-smtp-settings-for-outlook-com-d088b986-291d-42b8-9564-9c414e2aa040) settings for Outlook.com
2. [How to send](https://mailtrap.io/blog/golang-send-email/) emails in Go
3. [Load config](https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d) from file & environment variables in Golang with Viper

## Licence

MIT
