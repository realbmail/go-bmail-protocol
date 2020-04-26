package test

import (
	"fmt"
	"github.com/BASChain/go-bmail-protocol/bmprotocol"
	"github.com/BASChain/go-bmail-protocol/translayer"
	"math/rand"
	"testing"
)

func Test_EnvelopeHead(t *testing.T) {
	eh := &bmprotocol.EnvelopeHead{}

	eh.From = "a@bas"
	eh.RecpAddr = "b@bas"

	pubkey := make([]byte, 32)

	for {
		n, _ := rand.Read(pubkey)
		if n != len(pubkey) {
			continue
		}
		break
	}
	eh.LPubKey = pubkey

	data, _ := eh.Pack()

	fmt.Println(eh.String())

	ehUnpack := &bmprotocol.EnvelopeHead{}
	ehUnpack.UnPack(data)

	fmt.Println(ehUnpack.String())

	if eh.String() == ehUnpack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}

}

func Test_EnvelopeContent(t *testing.T) {
	ec := &bmprotocol.EnvelopeContent{}

	ec.To = []string{"toa@bas", "tob@bas", "toc@bas"}
	ec.CC = []string{"cca@bas", "ccb@bas"}
	ec.BC = []string{"bca@bas"}

	ec.Subject = "test a ec"
	ec.Data = "test e content"

	data, _ := ec.Pack()

	fmt.Println(ec.String())

	ecUnpack := &bmprotocol.EnvelopeContent{}
	ecUnpack.UnPack(data)
	fmt.Println(ecUnpack.String())

	if ec.String() == ecUnpack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}
}

func Test_EnvelopeTail(t *testing.T) {
	et := &bmprotocol.EnvelopeTail{}

	iv := make([]byte, 16)

	for {
		n, _ := rand.Read(iv)
		if n != len(iv) {
			continue
		}
		break
	}

	sig := make([]byte, 32)

	for {
		n, _ := rand.Read(sig)
		if n != len(sig) {
			continue
		}
		break
	}

	et.IV = iv
	et.Sig = sig

	data, _ := et.Pack()
	fmt.Println(et.String())

	etUnpack := &bmprotocol.EnvelopeTail{}
	etUnpack.UnPack(data)

	fmt.Println(etUnpack.String())

	if et.String() == etUnpack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}

}

func fillEH(eh *bmprotocol.EnvelopeHead) {
	eh.From = "a@bas"
	eh.RecpAddr = "b@bas"

	pubkey := make([]byte, 32)

	for {
		n, _ := rand.Read(pubkey)
		if n != len(pubkey) {
			continue
		}
		break
	}
	eh.LPubKey = pubkey
}

func fillET(et *bmprotocol.EnvelopeTail) {
	iv := make([]byte, 16)

	for {
		n, _ := rand.Read(iv)
		if n != len(iv) {
			continue
		}
		break
	}

	sig := make([]byte, 32)

	for {
		n, _ := rand.Read(sig)
		if n != len(sig) {
			continue
		}
		break
	}

	et.IV = iv
	et.Sig = sig
}

func Test_SendEnvelope(t *testing.T) {
	se := bmprotocol.NewSendEnvelope()

	eh := &se.EnvelopeHead

	fillEH(eh)

	et := &se.EnvelopeTail

	fillET(et)

	cipher := make([]byte, 32)

	for {
		n, _ := rand.Read(cipher)
		if n != len(cipher) {
			continue
		}
		break
	}

	se.CipherTxt = cipher

	data, _ := se.Pack()

	fmt.Println(se.String())

	seUnpack := &bmprotocol.SendEnvelope{}

	bmtl := &translayer.BMTransLayer{}
	n, _ := bmtl.UnPack(data)

	seUnpack.BMTransLayer = *bmtl

	seUnpack.UnPack(data[n:])

	fmt.Println(seUnpack.String())

	if se.String() == seUnpack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}

}

func Test_RespSendEnvelope(t *testing.T) {
	rse := bmprotocol.NewRespSendEnvelope()
	eh := &rse.EnvelopeHead
	fillEH(eh)

	iv := make([]byte, 16)

	for {
		n, _ := rand.Read(iv)
		if n != len(iv) {
			continue
		}
		break
	}

	rse.IV = iv

	data, _ := rse.Pack()

	fmt.Println(rse.String())

	rseUnpack := &bmprotocol.RespSendEnvelope{}

	bmtl := &translayer.BMTransLayer{}
	n, _ := bmtl.UnPack(data)

	rseUnpack.BMTransLayer = *bmtl

	rseUnpack.UnPack(data[n:])

	fmt.Println(rseUnpack.String())

	if rse.String() == rseUnpack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}

}

func Test_SendEnvelopeFail(t *testing.T) {
	sef := bmprotocol.NewSendEnvelopeFail()

	eh := &sef.EnvelopeHead

	fillEH(eh)

	et := &sef.EnvelopeTail

	fillET(et)

	cipher := make([]byte, 32)

	for {
		n, _ := rand.Read(cipher)
		if n != len(cipher) {
			continue
		}
		break
	}

	sef.CipherTxt = cipher

	sef.ErrorCode = 1

	data, _ := sef.Pack()

	fmt.Println(sef.String())

	bmtl := &translayer.BMTransLayer{}
	n, _ := bmtl.UnPack(data)

	sefUnpack := &bmprotocol.SendEnvelopeFail{}
	sefUnpack.BMTransLayer = *bmtl
	sefUnpack.UnPack(data[n:])

	fmt.Println(sefUnpack.String())

	if sef.String() == sefUnpack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}
}

func Test_RespSendEnvelopeFail(t *testing.T) {
	rsef := bmprotocol.NewRespSendEnvelopeFail()
	eh := &rsef.EnvelopeHead

	fillEH(eh)

	iv := make([]byte, 16)

	for {
		n, _ := rand.Read(iv)
		if n != len(iv) {
			continue
		}
		break
	}

	rsef.IV = iv

	data, _ := rsef.Pack()
	fmt.Println(rsef.String())

	bmtl := &translayer.BMTransLayer{}
	n, _ := bmtl.UnPack(data)
	rsefunPack := &bmprotocol.RespSendEnvelopeFail{}
	rsefunPack.BMTransLayer = *bmtl
	rsefunPack.UnPack(data[n:])

	fmt.Println(rsefunPack.String())

	if rsef.String() == rsefunPack.String() {
		t.Log("pass")
	} else {
		t.Fatal("failed")
	}

}
