package tests

import(
	"testing"
	"github.com/stretchr/testify/require"
	"server/src/models"
)


func TestValidateFieldsCompany(t *testing.T){
	model := models.CompanyModel{}

	company := models.Company{Name: "Company", ZipCode: "12345", Website: "http://teste.com.br"}
	result := model.Prepare(&company)
	require.Exactly(t, true, result)

	//Wrong Company Name
	company = models.Company{Name: "", ZipCode: "12345", Website: "http://teste.com.br"}
	result = model.Prepare(&company)
	require.Exactly(t, false, result)

	// Wrong Company ZipCode
	company = models.Company{Name: "Company", ZipCode: "1234", Website: "http://teste.com.br"}
	result = model.Prepare(&company)
	require.Exactly(t, false, result)

	// Wrong Company Website
	company = models.Company{Name: "Company", ZipCode: "1234", Website: "http://teste"}
	result = model.Prepare(&company)
	require.Exactly(t, false, result)
}
