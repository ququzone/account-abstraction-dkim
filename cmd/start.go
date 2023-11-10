package cmd

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ququzone/account-abstraction-dkim/pkg/db"
	"github.com/ququzone/account-abstraction-dkim/pkg/dkim"
	"github.com/ququzone/account-abstraction-dkim/pkg/mail"
	"github.com/ququzone/account-abstraction-dkim/pkg/recovery"
)

func start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start DKIM service",
		Action: func(ctx *cli.Context) error {
			repository := db.NewRepository()
			err := repository.AutoMigrate(&db.Recovery{})
			if err != nil {
				log.Fatalf("database migrate error: %v", err)
			}

			recovery, err := recovery.NewRecovery(
				os.Getenv("PRIVATE_KEY"),
				os.Getenv("RPC"),
			)
			if err != nil {
				log.Fatalf("new recovery error: %v\n", err)
			}

			go func() {
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
						account, hash, err := recovery.PendingRecover(
							header.Server,
							header.Subject,
							header.HeaderData(),
							header.Signature,
						)
						if err != nil {
							log.Printf("recovery account error: %v\n", err)
							continue
						}

						recovery := db.Recovery{
							TxHash:  hash,
							Account: account,
							Status:  0,
						}
						err = repository.Save(&recovery).Error
						if err != nil {
							log.Printf("create recovery account error: %v\n", err)
							continue
						}

						log.Printf("Success recovery hash: %s\n", hash)
					}
				}
			}()

			go func() {
				for {
					now := time.Now().Add(-24 * time.Hour).Add(-time.Minute)
					var recoveries []db.Recovery
					err := repository.Model(&db.Recovery{}).Where("status = ? and created_at < ?", 0, now).Limit(10).Scan(&recoveries).Error
					if err != nil {
						log.Printf("query recoveries account error: %v\n", err)
						continue
					}
					for _, r := range recoveries {
						hash, err := recovery.Recover(r.Account)
						if err != nil {
							log.Printf("recover account error: %v\n", err)
							r.Status = 2
						} else {
							r.Status = 1
							r.ComfirmHash = hash
						}
						err = repository.Save(r).Error
						if err != nil {
							log.Printf("save recovery account error: %v\n", err)
						}
					}
					if len(recoveries) < 10 {
						time.Sleep(10 * time.Second)
					}
				}
			}()

			select {}
		},
	}
}
