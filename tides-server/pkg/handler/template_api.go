package handler

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	"tides-server/pkg/restapi/operations/template"

	"tides-server/pkg/config"
	"tides-server/pkg/models"
)

func AddTemplateHandler(params template.AddTemplateParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return template.NewAddTemplateUnauthorized()
	}

	body := params.ReqBody
	newTem := models.Template{
		Compatibility:    body.Compat,
		GuestOS:          body.Os,
		MemorySize:       body.Memsize,
		Name:             body.Name,
		ProvisionedSpace: body.Space,
		TemplateType:     body.Source,
		VmName:           body.VMName,
	}

	db := config.GetDB()
	err := db.Create(&newTem).Error
	if err != nil {
		return template.NewAddTemplateBadRequest()
	}

	return template.NewAddTemplateOK().WithPayload(&template.AddTemplateOKBody{
		Message: "success",
		ID:      int64(newTem.Model.ID),
	})
}

func ListTemplateHandler(params template.ListTemplateParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return template.NewListTemplateUnauthorized()
	}

	db := config.GetDB()
	var templates []*models.Template
	db.Find(&templates)
	var result []*template.ListTemplateOKBodyResultsItems0

	for _, tem := range templates {
		newItem := template.ListTemplateOKBodyResultsItems0{
			Compatibility:    tem.Compatibility,
			DateAdded:        time.Time.String(tem.Model.CreatedAt),
			GuestOS:          tem.GuestOS,
			MemorySize:       tem.MemorySize,
			Name:             tem.Name,
			ProvisionedSpace: tem.ProvisionedSpace,
			TemplateType:     tem.TemplateType,
		}
		result = append(result, &newItem)
	}

	return template.NewListTemplateOK().WithPayload(&template.ListTemplateOKBody{
		Message: "success",
		Results: result,
	})
}

func DeleteTemplateHandler(params template.DeleteTemplateParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return template.NewDeleteTemplateUnauthorized()
	}

	uid, _ := ParseUserIdFromToken(params.HTTPRequest)
	db := config.GetDB()

	if db.Unscoped().Where("id = ? AND user_id = ?", params.ID, uid).Delete(&models.Template{}).RowsAffected == 0 {
		return template.NewDeleteTemplateForbidden()
	}

	return template.NewDeleteTemplateOK().WithPayload(&template.DeleteTemplateOKBody{
		Message: "success",
	})
}
