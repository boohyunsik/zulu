package signer

type signer interface {
	sign(tx interface{}, k key)
}
