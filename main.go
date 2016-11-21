// Program jwtverify is a tool to verify and decode JSON Web Tokens.
package main

import (
	"flag"
	"fmt"
	"os"
)

const helpText = `Usage: jwtverify is a tool to verify and decode JSON Web Tokens.

    jwtverify <action> <token> [-k key] [-s secret] [-p]

Actions:
    verify	Verify whether the token is valid.
    decode	Only decode the token, skip verification.

    help	Print this.

Options:
    -k		Public key (path to pem file) to verify the token. To use with RSA.
    -s		Secret used to sign the token. To use with HS.
    -p		Pretty output.
`

var cmds = map[string]func(args ...string){
	"decode": decode,
	"verify": verify,
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		usage("")
	}
	if len(os.Args) < 3 {
		usage("insufficient arguments")
	}

	cmd := os.Args[1]
	fn, ok := cmds[cmd]
	if !ok {
		usage("unknown command " + cmd)
	}

	fn(os.Args[2:]...)
}

func decode(args ...string) {
	token := args[0]

	var pretty bool
	flag.BoolVar(&pretty, "p", false, "Pretty output.")
	flag.CommandLine.Parse(args[1:])

	decodeToken(token, pretty)
}

func verify(args ...string) {
	token := args[0]

	var keyFile string
	var secret string

	flag.StringVar(&keyFile, "k", "", "Key (path to pem file) used to sign the token.")
	flag.StringVar(&secret, "s", "", "Secret used to sign the token.")
	flag.CommandLine.Parse(args[1:])

	verifyToken(token, keyFile, secret)
}

func usage(msg string) {
	if msg != "" {
		fmt.Printf("jwtverify: %v\n", msg)
	}
	fmt.Fprintf(os.Stderr, "%v\n", helpText)
	os.Exit(1)
}

func fail(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "jwtverify: %v\n", msg)
	}
	os.Exit(1)
}
