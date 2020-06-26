package authorized_keys

type SSHKey struct {
	Key   string
	Email string
}

func Parse(data []byte) ([]SSHKey, error) {
	keysStrings := strings.Split(rawToString, "\n")
	for i, s := range keysStrings {
		if len(s) == 0 {
			continue
		}
		k := strings.Fields(s)
		var sshKey SSHKey
		sshKey.Num = i + 1
		sshKey.Key = k[0] + " " + k[1]
		if len(k) > 2 {
			sshKey.Email = k[2]
		}
		sshKeys = append(sshKeys, sshKey)
	}
	return sshKeys, nil
}

func Generate(keys []SSHKey) ([]byte) {
	var out *bytes.Buffer
	for _, k := range keys {
		fmt.Fprintln(out, k.Key+" "+k.Email)
	}
	return out.Bytes()
}

