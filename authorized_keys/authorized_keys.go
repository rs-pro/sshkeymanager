package authorized_keys

func GenerateAuthorizedKeys(keys []SSHKey) ([]byte) {
	var out *bytes.Buffer
	for _, k := range keys {
		fmt.Fprintln(out, k.Key+" "+k.Email)
	}
	return out.Bytes()
}
