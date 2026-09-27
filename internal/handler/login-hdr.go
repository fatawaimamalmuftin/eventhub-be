package handler

type ILoginSrv interface {
	//isi method di service
}

type LoginHdr struct {
	Lh ILoginSrv
}

func LoginHandler(ls ILoginSrv) *LoginHdr {
	return &LoginHdr{
		Lh: ls,
	}
}
