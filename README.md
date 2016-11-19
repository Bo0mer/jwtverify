jwtverify
=========
Verify and decode JWTs. For more info on JWT, see [this](http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html).

## User's manual
```
Usage: jwtverify is a tool to verify and decode JSON Web Tokens.

    jwtverify <action> <token> [-k key] [-s secret] [-p]

Actions:
    verify	Verify whether the token is valid.
    decode	Only decode the token, skip verification.

    help	Print this.

Options:
    -k		Key (path to pem file) used to sign the token. To use with RSA.
    -s		Secret used to sign the token. To use with HS.
    -p		Pretty output.
```
