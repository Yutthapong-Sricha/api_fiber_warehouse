package sessions

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"api_fiber_warehouse/initializers"
	modelsStruc "api_fiber_warehouse/models/struc/session"
)

var store = session.New(session.Config{
	CookieHTTPOnly: true,
	// CookieSecure: true, for https
	Expiration: time.Hour * 12,
})

var (
	//store     *session.Store = x
	AUTH_KEY  string = "authenticated"
	USER_ID   string = "id_user"
	BRANCH_ID string = "id_branch"
)

func GenSession(c *fiber.Ctx, LogLogin modelsStruc.LogLogin) (string, error) {

	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return "error", sessErr
	}

	if sess.Fresh() {
		sess.Set(AUTH_KEY, true)
		sess.Set(USER_ID, LogLogin.Id_user)
		sess.Set(BRANCH_ID, LogLogin.Id_branch)

		sessErr = sess.Save()
		if sessErr != nil {
			return "error", sessErr
		}

		clientIp := c.IP()
		agent := string(c.Request().Header.UserAgent())
		session_data, err := json.Marshal(LogLogin)
		if err != nil {
			log.Fatalln(err.Error())
		}

		//res := c.Response().String()
		resCookie := c.GetRespHeader("set-cookie")

		findSession := strings.Split(resCookie, "session_id=")[1] // ebe8529d-b8bd-4876-abb7-f75b82d93415; max-age=180; path=/; HttpOnly; SameSite=Lax
		sessionToken := strings.Split(findSession, ";")[0]        // ebe8529d-b8bd-4876-abb7-f75b82d93415

		sqlStmt := `INSERT INTO oauth_sessions(session_id, session_start, session_last_activity, session_ip_address, session_user_agent, session_data)
			VALUES (?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), ?, ?, ? )`
		stmt, err := initializers.DB.Prepare(sqlStmt)
		if err != nil {
			log.Fatalln(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(sessionToken, clientIp, agent, session_data)
		if err != nil {
			log.Fatalln(err)
		}
		return sessionToken, nil
	} else {
		myError := errors.New("session error")
		return "error", myError
	}

}

func SessionLogin(c *fiber.Ctx, LogLogin modelsStruc.LogLogin) (string, error) {

	now := time.Now()
	var timestamp = strconv.FormatInt(now.Unix(), 10)
	login_time, _ := strconv.ParseInt(timestamp, 10, 0)
	LogLogin.Is_success1_fail2_flag = 1
	LogLogin.Login_time = int(login_time)

	sessionToken, err := GenSession(c, LogLogin)
	if err != nil {
		return "error", err
	}

	sqlStmt := `INSERT INTO oauth_log_user_login(session_id, id_user, id_branch, email_login, is_success1_fail2_flag, login_time) 
				VALUES (?, ?, ?, ?, ?, UNIX_TIMESTAMP())`
	stmt, err := initializers.DB.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionToken, LogLogin.Id_user, LogLogin.Id_branch, LogLogin.Email_login, 1)
	if err != nil {
		log.Fatalln(err)
	}

	return sessionToken, nil

}
