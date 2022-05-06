package services

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"secondChance/internal/db"
	"secondChance/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type CustomerService struct {
	repo db.Customer
}

func NewCustomerService(repo db.Customer) *CustomerService {
	return &CustomerService{repo: repo}
}

func (c *CustomerService) Get(param string) (*models.Customer, error) {
	user, err := c.repo.Get(param)
	if err != nil {
		return &models.Customer{}, err
	}
	return user, nil
}

func (c *CustomerService) Create(user *models.Customer) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("Hash Password: %w", err)
	}
	user.Password = hash
	fmt.Printf("%+v\n", user)
	if err := c.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) Login(param *models.LoginInput) (t string, id int, err error) {
	customer, err := c.repo.Get(param.Email)
	if err != nil {
		return "", 0, err
	}

	if !CheckPasswordHash(param.Password, customer.Password) {
		return "", 0, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = customer.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err = token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", 0, err
	}
	return t, customer.Id, nil
}

func (c *CustomerService) CreateOrder(order *models.Order) (err error) {
	if err := c.repo.CreateOrder(order); err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) GetOrder(id *models.IdReg) (*[]models.Product, error) {
	products, err := c.repo.GetOrder(id)
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}

func (c *CustomerService) SaveImage(id *models.IdReg, file string) (string, error) {
	path, err := c.repo.SaveImage(id, file)
	if err != nil {
		return "", err
	}
	return path, nil
}
func (c *CustomerService) DeleteImage(id *models.IdReg) error {
	if err := c.repo.DeleteImage(id); err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) GmailCode(email *models.EmailRequest) (int, error) {
	if err := godotenv.Load(); err != nil {
		return -1, err
	}

	// Sender data.
	from := os.Getenv("sender-email")
	password := os.Getenv("sender-password")

	// Receiver email address.
	to := []string{
		"180107192@stu.sdu.edu.kz",
	}

	// smtp server configuration.
	smtpHost := os.Getenv("smtpHost")
	smtpPort := os.Getenv("smtpPort")

	// Message.
	code := rand.Intn(899999) + 100000
	msg := fmt.Sprintf("From: %s\r\n To: %s\r\n Subject: verify code\r\n\r\n Code: %d\r\n", from, to[0], code)
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return -1, err
	}
	return code, nil
}
