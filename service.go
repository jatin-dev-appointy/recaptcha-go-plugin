package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"recaptcha"
	"strings"
)

func GenerateService(g *protogen.GeneratedFile, service *protogen.Service) {

	serviceServer := service.GoName + "Server"
	unexposedRecaptchaServer := strings.ToLower(service.GoName[:1]) + service.GoName[1:] + "RecaptchaServer"

	g.P()
	g.P("type ", unexposedRecaptchaServer, " struct {")
	g.P("\t", serviceServer)
	g.P("\trecaptchaServer recaptcha_module.ReCaptchaServer")
	g.P("}")
	g.P()
	g.P("func New", service.GoName, "RecaptchaServer(srv ", serviceServer, ", rSrv recaptcha_module.ReCaptchaServer) ", serviceServer, " {")
	g.P("\treturn &", unexposedRecaptchaServer, "{")
	g.P("\t\t", serviceServer, ": srv,")
	g.P("\t\trecaptchaServer: rSrv,")
	g.P("\t}")
	g.P("}")

	for _, method := range service.Methods {
		reqRecaptcha := proto.GetExtension(method.Desc.Options(), recaptcha.E_RequireRecaptcha).(bool)

		g.P()
		if reqRecaptcha {
			g.P("// Method: ", method.GoName, " (This method requires reCAPTCHA verification)")
		} else {
			g.P("// Method: ", method.GoName, " (This method does not require reCAPTCHA verification)")
		}

		if reqRecaptcha {
			g.P("func (s *", unexposedRecaptchaServer, ") ", method.GoName, "(ctx context.Context, req *", g.QualifiedGoIdent(method.Input.GoIdent), ") (*", g.QualifiedGoIdent(method.Output.GoIdent), ", error) {")
			g.P()
			g.P("\tif err := s.recaptchaServer.ValidateRecaptcha(ctx); err != nil {")
			g.P("\t\treturn nil, err")
			g.P("\t}")
			g.P()

			g.P("\treturn s.", serviceServer, ".", method.GoName, "(ctx, req)")
			g.P("}")
		}
	}
}
