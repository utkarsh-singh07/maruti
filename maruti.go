package maruti

import (
	"github.com/utkarsh-singh07/hanuman"
)

func Chant() string {
	return "Jai shree Ram "
}

func Chants() string {
	return "Jai shree Ram!, Jai shree Ram!"
}

func Roar() string {
	return hanuman.Shouts(Chant())
}
