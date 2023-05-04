package controllers

import (
	"api_fiber_warehouse/initializers"
	modelsStruc "api_fiber_warehouse/models/struc/session"
	"encoding/json"
	"log"

	modelSess "api_fiber_warehouse/models/sessions"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

func Login(c *fiber.Ctx) error {

	var Login modelsStruc.Login

	if err := c.BodyParser(&Login); err != nil {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
	}

	var LogLogin modelsStruc.LogLogin
	var err error
	var pwd_verification string

	statement, err := initializers.DB.Prepare("SELECT id_staff_core AS id_user, id_branch,  tmp_branch_name, staff_email AS email_login, pwd_verification, staff_fullname, IFNULL(tmp_position_name,'') AS tmp_position_name FROM staff_core WHERE id_branch=? and staff_email=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	row_auth := statement.QueryRow(Login.Id_branch, Login.Username)
	err = row_auth.Scan(
		&LogLogin.Id_user,
		&LogLogin.Id_branch,
		&LogLogin.Tmp_branch_name,
		&LogLogin.Email_login,
		&pwd_verification,
		&LogLogin.Staff_fullname,
		&LogLogin.Tmp_position_name,
	)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "ไม่พบข้อมูลในระบบ"})
	}

	err_valid := bcrypt.CompareHashAndPassword([]byte(pwd_verification), []byte(Login.Password))
	if err_valid != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "รหัสผ่านไม่ถูกต้อง"})
	}

	sessionToken, err := modelSess.SessionLogin(c, LogLogin)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "error", "message": "something went wrong: " + err.Error(),
		})
	}
	session_data, err := json.Marshal(LogLogin)
	if err != nil {
		log.Fatalln(err.Error())
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	resCookie := c.GetRespHeader("set-cookie")
	cookie, err := json.Marshal(resCookie)
	if err != nil {
		log.Fatalln(err.Error())
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Login Success", "token": sessionToken, "data": string(session_data), "response": string(cookie)})

}
