package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func decodeToken(token string, pretty bool) {
	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		fail("invalid token format")
	}

	decoded, err := decodeSegments(segments)
	if err != nil {
		fail(fmt.Sprintf("error decoding token: %v", err))
	}
	if pretty {
		decoded[0] = formatJSON(decoded[0])
		decoded[1] = formatJSON(decoded[1])
	}
	fmt.Printf("%s\n", decoded[0])
	fmt.Printf("%s\n", decoded[1])
}

func decodeSegments(segments []string) ([][]byte, error) {
	decoded := make([][]byte, 3)
	for i, seg := range segments {
		dec, err := jwt.DecodeSegment(seg)
		if err != nil {
			return nil, err
		}
		decoded[i] = dec
	}
	return decoded, nil
}

func formatJSON(plain []byte) []byte {
	dst := bytes.NewBuffer(nil)
	json.Indent(dst, plain, "", "  ")
	return dst.Bytes()
}
