package build

import (
	"encoding/json"
	"os"
	"reflect"
)

type SkinPath struct {
	Index  string `json:"index"`
	Post   string `json:"post"`
	Header string `json:"header"`
	Footer string `json:"footer"`
	Nav    string `json:"nav"`
}

type SkinInfo struct {
	Name    string   `json:"name"`
	Author  string   `json:"author"`
	Version string   `json:"version"`
	Paths   SkinPath `json:"paths"`
}

type Skin struct {
	Info SkinInfo
}

func MakeSkin() Skin {
	return Skin{}
}

func (s *Skin) Get_skin() {
	skinpath, perr := os.Getwd()

	if perr != nil {
		panic(perr)
	}

	ctx, ferr := os.ReadFile(skinpath + "/skin/skin.json")

	if ferr != nil {
		panic(ferr)
	}

	var sinfo SkinInfo = SkinInfo{}
	jerr := json.Unmarshal(ctx, &sinfo)

	if jerr != nil {
		panic(jerr)
	}

	s.Info = sinfo

	var infos SkinInfo = SkinInfo{}

	ref := reflect.ValueOf(&sinfo.Paths).Elem()
	sref := reflect.ValueOf(&infos.Paths).Elem()

	for i := 0; i < ref.NumField(); i++ {
		pelm_val := ref.Field(i)
		pelm_typ := ref.Type().Field(i)

		srefval := sref.FieldByName(pelm_typ.Name)
		pval := skinpath + (pelm_val.Interface().(string))
		srefval.SetString(pval)
	}

	s.Info.Paths = infos.Paths

}
