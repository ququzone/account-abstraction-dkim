package cmd

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ququzone/account-abstraction-dkim/pkg/dkim"
	"github.com/ququzone/account-abstraction-dkim/pkg/mail"
	"github.com/ququzone/account-abstraction-dkim/pkg/recovery"
)

func start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start DKIM service",
		Action: func(ctx *cli.Context) error {
			go func() {
				recovery, err := recovery.NewRecovery(
					os.Getenv("KEY_FILE"),
					os.Getenv("KEY_PASSPHRASE"),
					os.Getenv("RPC"),
				)
				if err != nil {
					log.Fatalf("new recovery error: %v\n", err)
				}

				for {
					mails, err := mail.Fetch(
						os.Getenv("IMAP_SERVER"),
						os.Getenv("IMAP_USERNAME"),
						os.Getenv("IMAP_PASSWORD"),
					)
					if err != nil {
						log.Fatalf("fetch emails error: %v\n", err)
					}
					if len(mails) == 0 {
						time.Sleep(10 * time.Second)
						continue
					}
					for _, mail := range mails {
						header, err := dkim.Parse(mail, true)
						if err != nil {
							log.Printf("parse email error: %v\n", err)
							continue
						}
						hash, err := recovery.Recover(
							header.Server,
							header.Subject,
							header.HeaderData(),
							header.Signature,
						)
						if err != nil {
							log.Printf("recovery account error: %v\n", err)
							continue
						}
						log.Printf("Success recovery hash: %s\n", hash)
					}
				}
			}()

			select {}
		},
	}
}
