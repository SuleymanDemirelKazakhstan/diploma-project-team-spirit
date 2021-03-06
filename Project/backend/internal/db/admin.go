package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"secondChance/internal/models"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type AdminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

func (a *AdminRepo) Create(user *models.Owner) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	sqlStatement := `INSERT INTO shop (phone, email, name, password, address) VALUES ($1, $2, $3, $4, $5)`
	if err := a.db.QueryRow(sqlStatement, user.Phone, user.Email, user.Name, user.Password, user.Address); err != nil {
		return err.Err()
	}
	return nil
}

func (a *AdminRepo) Get(param *models.IdReg) (*models.Owner, error) {
	var user models.Owner
	sqlStatement := `SELECT shop_id,name,email,phone,address,image FROM shop WHERE shop_id=$1 and is_deleted=false`

	row := a.db.QueryRow(sqlStatement, param.Id)
	// unmarshal the row object to user
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Address, &user.Image); err != nil {
		return &models.Owner{}, err
	}
	if err := godotenv.Load(); err != nil {
		return &models.Owner{}, err
	}
	_url := os.Getenv("baseUrl")
	user.Image = _url + user.Image
	return &user, nil
}

func (a *AdminRepo) Delete(param *models.IdReg) error {
	sqlStatement := `update shop set is_deleted=true WHERE shop_id=$1`
	if _, err := a.db.Exec(sqlStatement, param.Id); err != nil {
		return err
	}
	return nil
}

//TODO: for the admin to return all stores even those that were deleted
func (a *AdminRepo) GetAll() ([]models.Owner, error) {
	var users []models.Owner
	sqlStatement := `SELECT name,email,password,phone,address, image, description FROM shop where is_deleted=false`

	rows, err := a.db.Query(sqlStatement)
	if err != nil {
		return []models.Owner{}, err
	}
	defer rows.Close()

	if err := godotenv.Load(); err != nil {
		return []models.Owner{}, err
	}
	_url := os.Getenv("baseUrl")
	for rows.Next() {
		var user models.Owner
		var description sql.NullString
		if err := rows.Scan(&user.Name, &user.Email, &user.Password, &user.Phone, &user.Address, &user.Image, &description); err != nil {
			return []models.Owner{}, err
		}
		user.Description = description.String
		user.Image = _url + user.Image
		users = append(users, user)
	}
	return users, nil
}

func (a *AdminRepo) Update(user *models.Owner) error {
	if user.Id == 0 {
		return fmt.Errorf("id is not specified")
	}
	str := []string{}
	if user.Email != "" {
		str = append(str, user.Email)
	}
	if user.Name != "" {
		str = append(str, user.Name)
	}
	if user.Phone != "" {
		str = append(str, user.Phone)
	}
	if user.Address != "" {
		str = append(str, user.Address)
	}
	if user.Description != "" {
		str = append(str, user.Description)
	}
	if user.Password != "" {
		var err error
		user.Password, err = HashPassword(user.Password)
		if err != nil {
			return err
		}
		str = append(str, user.Password)
	}

	if len(str) == 0 {
		return fmt.Errorf("there is no data to update")
	}

	sqlStatement := fmt.Sprintf(`UPDATE shop SET %v WHERE shop_id=$1`, strings.Join(str, ","))
	_, err := a.db.Exec(sqlStatement, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AdminRepo) GetLogin(param *models.Login) (*models.Owner, error) {
	var user models.Owner
	sqlStatement := `SELECT shop_id,email,password FROM shop WHERE shop_id=$1`

	row := a.db.QueryRow(sqlStatement, param.Email)
	// unmarshal the row object to user
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return &models.Owner{}, err
	}

	if !CheckPasswordHash(param.Password, user.Password) {
		return nil, errors.New("login error")
	}

	return &user, nil
}

func (a *AdminRepo) SaveImage(id *models.IdReg, file string) (string, error) {
	// generate new uuid for image name
	uniqueId := uuid.New()
	// remove "- from imageName"
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	// extract image extension from original file filename
	fileExt := strings.Split(file, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	path := fmt.Sprintf("/images/shop/%d/%s", id.Id, image)
	if err := os.MkdirAll(fmt.Sprintf("./images/shop/%d", id.Id), os.ModePerm); err != nil {
		return "", err
	}

	sqlStatement := `UPDATE shop SET image=$2 WHERE shop_id=$1`
	_, err := a.db.Exec(sqlStatement, id.Id, path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (a *AdminRepo) DeleteImage(id *models.IdReg) error {
	sqlStatement := `UPDATE shop SET image=NULL WHERE shop_id=$1`
	_, err := a.db.Exec(sqlStatement, id.Id)
	if err != nil {
		return err
	}
	return nil
}
