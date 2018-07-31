package http

import (
	"net/http"
	"strings"

	"github.com/open-falcon/mail-provider/config"
	"github.com/open-falcon/mail-provider/smtps"
	//	"../smtps"
	"github.com/toolkits/smtp"
	"github.com/toolkits/web/param"
)

func configProcRoutes() {

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		tos = strings.Replace(tos, ",", ";", -1)

		s := smtp.New(cfg.Smtp.Addr+":25", cfg.Smtp.Username, cfg.Smtp.Password)
		err := s.SendMail(cfg.Smtp.From, tos, subject, content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "success", http.StatusOK)
		}
	})

	http.HandleFunc("/sender/sslmail", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		subject := param.MustString(r, "subject")
		content := param.MustString(r, "content")
		toss := strings.Split(tos, ",")

		err := smtps.SendMail(toss, cfg.Smtp.From, subject, content, cfg.Smtp.Addr, cfg.Smtp.Password, cfg.Smtp.Username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "success", http.StatusOK)
		}
	})

}
