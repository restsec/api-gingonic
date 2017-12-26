package controllers

import (
	"fmt"
	"log"
	"pessoalAPI/db"

	"github.com/gin-gonic/gin"
)

type Pessoal struct {
	ID                 int    `db:"id, primarykey, autoincrement" json:"id"`
	siape              string `db:"siape" json:"siape"`
	id_pessoa          string `db:"id_pessoa" json:"id_pessoa"`
	Nome               string `db:"nome" json:"nome"`
	matricula_interna  string `db:"matricula_interna" json:"matricula_interna"`
	nome_identificacao string `db:"nome_identificacao" json:"nome_identificacao"`
	data_nascimento    string `db:"data_nascimento" json:"data_nascimento"`
	sexo               string `db:"sexo" json:"sexo"`
}

type PessoalController struct{}

func (ctrl PessoalController) GetPessoal(c *gin.Context) { // Hello
	//func (ctrl PessoalController) getPessoal(c *gin.Context) (pessoal Pessoal, err error) {
	// q := `select s.id_servidor, s.siape, s.id_pessoa, s.matricula_interna,
	// 	s.id_foto, s.nome_identificacao,
	// 	p.nome, p.data_nascimento, p.sexo from rh.servidor s
	// inner join comum.pessoa p on (s.id_pessoa = p.id_pessoa) `

	q2 := "select id_servidor from rh.servidor"

	var pessoas []int64
	row, err := db.GetDB().Query(q2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(row.Scan())
	//	err = db.GetDB().Select(&pessoal, q)
	fmt.Print(pessoas)

	c.JSON(200, gin.H{
		"message": "eh",
	})
	if err != nil {
		return
	}

	return
}