package k8s

type Secrets struct {
	Secrets []Secret
}

type Secret struct {
	Name      string
	Namespace string
}

func (sc *Secrets) AddSecret(secret Secret) []Secret {
	toAdd := true
	for i := range sc.Secrets {
		if sc.Secrets[i] == secret {
			toAdd = false
		}
	}
	if toAdd {
		sc.Secrets = append(sc.Secrets, secret)
	}

	return sc.Secrets
}
