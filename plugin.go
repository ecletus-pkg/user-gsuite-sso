package user_gsuite_sso

import (
	gsuite_sso_name_id_finder "github.com/ecletus-pkg/user/gsuite-sso-name-id-finder"
	"github.com/ecletus/plug"
	ect_samlidp "github.com/moisespsena/go-ecletus-samlidp"
)

type Plugin struct {
	SamlIdpKey string
}

func (this Plugin) Init(options *plug.Options) {
	idp := options.GetInterface(this.SamlIdpKey).(*ect_samlidp.SamlIDP)
	idp.NameIdFinders.Append(gsuite_sso_name_id_finder.Finder())
}
