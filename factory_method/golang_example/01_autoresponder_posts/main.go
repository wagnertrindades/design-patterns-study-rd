package main

import (
	"fmt"
	"sort"
)

func main() {
	frase := factoryFilosofos.create("vanDamme")
	fmt.Println(frase.fraseDeImpacto())
}

type vanDamme struct {
	frasesRefletidasDuranteSparcateComCadeiras []string
}

type chuckNorris struct {
	frasesAntesDaPropriaExistenciaDeFrasesSeparadasEmTopicos map[string]string
}

type marratimaGhandi struct {
	asDezFrasesDePaz map[int]string
}

type filosofo interface {
	fraseDeImpacto() string
}

func (v vanDamme) fraseDeImpacto() string {
	fraseSorteada := sort.StringSlice(v.frasesRefletidasDuranteSparcateComCadeiras)
	return fraseSorteada[0]
}

func (c chuckNorris) fraseDeImpacto() string {
	topicos := make([]string, 10)
	for key := range c.frasesAntesDaPropriaExistenciaDeFrasesSeparadasEmTopicos {
		topicos = append(topicos, key)
	}

	topicoSorteado := sort.StringSlice(topicos)
	fraseSorteada := sort.StringSlice(topicoSorteado)
	return fraseSorteada[0]
}

func (m marratimaGhandi) fraseDeImpacto() string {
	dezFrases := make([]int, 10)
	for key := range m.asDezFrasesDePaz {
		dezFrases = append(dezFrases, key)
	}

	fraseSorteada := sort.IntSlice(dezFrases)
	indice := fraseSorteada[0]
	return m.asDezFrasesDePaz[indice]
}

type factoryFilosofos struct {
}

func (fac factoryFilosofos) create(soldier string) filosofo {
	var object struct

	if soldier == "vanDamme" {
		object = vanDamme{[]string{
			"As Baleias são seres sensacionais",
			"O sparcate é a melhor forma de não entrar em depressão."}

	}

	return object
}
