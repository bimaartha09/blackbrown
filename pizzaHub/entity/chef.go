package entity

type Chef struct {
	ID          uint64
	NIP         string
	Name        string
	Description string
}

var Chefs = []Chef{}

func AddChef(chef Chef) Chef {
	chef.ID = generateIDChef()
	Chefs = append(Chefs, chef)

	return chef
}

func GetChefByNIP(nip string) Chef {
	for _, c := range Chefs {
		if c.NIP == nip {
			return c
		}
	}

	return Chef{}
}

func generateIDChef() uint64 {
	if len(Chefs) == 0 {
		return 1
	}

	return uint64(Chefs[len(Chefs)-1].ID) + 1
}
