package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/utils"
	"golang.org/x/crypto/bcrypt"
)

type MemberRepo interface {
	FindAll() ([]entity.Member, error)
	FindOne(id int) (entity.Member, error)
	Create(newMember *entity.Member) (entity.Member, error)
	Update(member *entity.Member) (entity.Member, error)
	Delete(id int) error
	LoginCheck(username string, password string) (string, error)
}

type memberRepo struct {
	db *sql.DB
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *memberRepo) LoginCheck(username string, password string) (string, error) {

	var err error

	u := entity.MemberLogin{}

	query := "SELECT member_id, username, password FROM m_member WHERE username = $1"
	row := r.db.QueryRow(query, username)
	err = row.Scan(&u.Member_Id, &u.Username, &u.Password)

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.Member_Id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *memberRepo) FindAll() ([]entity.Member, error) {
	var members []entity.Member

	query := "SELECT member_id, username, password, email_address, contact_number, wallet_amount, status FROM m_member ORDER BY member_id"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var member entity.Member
		if err := rows.Scan(&member.Member_Id, &member.Username, &member.Password, &member.Email_Address, &member.Contact_Number, &member.Wallet_Amount, &member.Status); err != nil {
			log.Println(err)

			return nil, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)

		return nil, err
	}

	return members, nil
}

func (r *memberRepo) FindOne(id int) (entity.Member, error) {
	var memberInDb entity.Member

	query := "SELECT member_id, username, password, email_address, contact_number, wallet_amount, status FROM m_member WHERE member_id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&memberInDb.Member_Id, &memberInDb.Username, &memberInDb.Password, &memberInDb.Email_Address, &memberInDb.Contact_Number, &memberInDb.Wallet_Amount, &memberInDb.Status)

	if err == sql.ErrNoRows {
		log.Println(err)

		return entity.Member{}, fmt.Errorf("member with id %d not found", id)
	} else if err != nil {
		log.Println(err)

		return entity.Member{}, err
	}

	return memberInDb, nil
}

func (r *memberRepo) Create(newMember *entity.Member) (entity.Member, error) {
	query := "INSERT INTO m_member (username, password, email_address, contact_number) VALUES ($1, $2, $3, $4) RETURNING member_id"

	// turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newMember.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.Member{}, err
	}

	var memberID int
	err = r.db.QueryRow(query, newMember.Username, string(hashedPassword), newMember.Email_Address, newMember.Contact_Number).Scan(&memberID)
	if err != nil {
		log.Println(err)
		return entity.Member{}, err
	}

	newMember.Member_Id = memberID

	return *newMember, nil
}

func (r *memberRepo) Update(member *entity.Member) (entity.Member, error) {
	query := "UPDATE m_member SET username = $1, password = $2, email_address = $3, contact_number = $4, status = $5 WHERE member_id = $6 RETURNING member_id, username, email_address, contact_number, status"
	row := r.db.QueryRow(query, member.Username, member.Password, member.Email_Address, member.Contact_Number, member.Status, member.Member_Id)

	var updatedMember entity.Member
	err := row.Scan(&updatedMember.Member_Id, &updatedMember.Username, &updatedMember.Email_Address, &updatedMember.Contact_Number, &updatedMember.Status)
	if err != nil {
		log.Println(err)
		return entity.Member{}, err
	}

	return updatedMember, nil
}

func (r *memberRepo) Delete(id int) error {
	query := "DELETE FROM m_member WHERE member_id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("member with id %d not found", id)
	}
	return nil
}

func NewMemberRepository(db *sql.DB) MemberRepo {
	repo := new(memberRepo)
	repo.db = db
	return repo
}
