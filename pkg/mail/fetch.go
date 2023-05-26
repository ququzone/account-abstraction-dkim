package mail

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func Fetch(server, username, password string) ([][]byte, error) {
	// Connect to the IMAP server
	c, err := client.DialTLS(server, nil)
	if err != nil {
		return nil, err
	}
	defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return nil, err
	}

	// Fetch the last 10 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 10 {
		from = mbox.Messages - 9
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	// Get the whole message body
	var section imap.BodySectionName
	items := []imap.FetchItem{section.FetchItem(), imap.FetchUid}

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqSet, items, messages)
	}()

	result := make([][]byte, 0)
	for msg := range messages {
		r := msg.GetBody(&section)
		if r == nil {
			log.Printf("Server didn't returned message body: %d", msg.Uid)
			continue
		}
		data := make([]byte, r.Len())
		_, err = r.Read(data)
		if err != nil {
			return nil, err
		}
		result = append(result, data)

		// Delete the email
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(1)
		operation := imap.FormatFlagsOp(imap.AddFlags, true)
		flags := []interface{}{imap.DeletedFlag}
		if err := c.Store(seqSet, operation, flags, nil); err != nil {
			return nil, err
		}
		err = c.Expunge(nil)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
