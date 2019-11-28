package wx_helper

import (
	"fmt"
	"testing"
)

func TestDecryptUserInfo(t *testing.T) {
	iv := "9qlYBw0AaUMhdxWWQ5oEqw=="
	encData := "uw3lYDSJLiJZX8BAvYOEbqNaG+GmdOPDFWQvG2qEcsZLFV8ZOAw+qExx7P+R9V4ho4k8wh4jiKkzFaA8btEZ23rfi+gWfSMjV/ixCr81pRPg75mG08LDY2c5CWGHxyoi5v64LdvmvFo8RzkolrXL4Kc6PLmyJTO9AIUIfDR9vQXVV2BxSKiQpM28arGUyiFpliDnEuTi/Z3qDp40q1a+6TCLJSq1YhPoPfss8ZT3YPE6IHm0ik7SdEmZ3LjQJLswmt4qnGqRls6KaIdVGBkMlJrPABB2T1Nf62D9/nnjkjSVM71TOFH3C69Dr/KgI2AjzM1Tp2uLpgyXOc9MVz+TtlMxP3A8js4v/+uDpWNeRgInOyIvOIwRONLJ9NDWkKcXdtNRPDDfooIq4gVkmzUlPv4+eIVm4xnfyp6aiH4YodkzdruN+G1urP8v9kxSmX9HX56K6mkgWW3O4vfZxre153P5xjrh4LGWOMbktR3sTJ0="
	sessionKey := "Xd7a0aAOexQ8DxQermKvvg=="
	userInfo, err := DecryptUserInfo(encData, iv, sessionKey)
	fmt.Println(userInfo, err)
}
