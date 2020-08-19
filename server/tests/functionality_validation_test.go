package tests

import(
	"testing"
	"github.com/stretchr/testify/require"
	"server/src/models"
)

func TestInsertCompany(t *testing.T){
	model := models.CompanyModel{}

	company := models.Company{Name: "Company", ZipCode: "12345", Website: "http://teste.com.br"}
	_, err := model.Insert(&company)
	require.Exactly(t, nil, err)

	//Wrong Company Name
	company = models.Company{Name: "", ZipCode: "1234", Website: "http://teste.com.br"}
	_, err = model.Insert(&company)
	require.Error(t, err)

	//Wrong Company ZipCode
	company = models.Company{Name: "Company", ZipCode: "1234", Website: "http://teste.com.br"}
	_, err = model.Insert(&company)
	require.Error(t, err)

	//Wrong Company Website
	company = models.Company{Name: "Company", ZipCode: "12345", Website: "http://"}
	_, err = model.Insert(&company)
	require.Error(t, err)
}

func TestGetAllCompanies(t *testing.T){
	model := models.CompanyModel{}

	_, err := model.GetAll()
	require.Exactly(t, nil, err)
}

func TestGetCompanyByNameAndZipCode(t *testing.T){
	model := models.CompanyModel{}

	company := models.Company{Name: "Company", ZipCode: "12345"}
	_, err := model.GetByNameAndZipCode(&company)
	require.Exactly(t, nil, err)

	//Wrong Company Name
	company = models.Company{Name: "A123321&&", ZipCode: "12345"}
	_, err = model.GetByNameAndZipCode(&company)
	require.Error(t, err)

	//Wrong Company ZipCode
	company = models.Company{Name: "Company", ZipCode: "12345678"}
	_, err = model.GetByNameAndZipCode(&company)
	require.Error(t, err)

	//Wrong Company Name and Company ZipCode
	company = models.Company{Name: "A123321&&", ZipCode: "12345678"}
	_, err = model.GetByNameAndZipCode(&company)
	require.Error(t, err)
}

func TestUpdateCompany(t *testing.T){
	model := models.CompanyModel{}

	company := models.Company{Name: "Company", ZipCode: "12345", Website: "http://novosite.com.br"}
	_, err := model.UpdateWebsite(&company)
	require.Exactly(t, nil, err)

	//Wrong Company Name
	company = models.Company{Name: "A123321&&", ZipCode: "12345", Website: "http://novosite.com.br"}
	_, err = model.UpdateWebsite(&company)
	require.Error(t, err)

	//Wrong Company ZipCode
	company = models.Company{Name: "Company", ZipCode: "12345678", Website: "http://novosite.com.br"}
	_, err = model.UpdateWebsite(&company)
	require.Error(t, err)

	//Wrong Company WebSite
	company = models.Company{Name: "A123321&&", ZipCode: "12345678", Website: "Teste"}
	_, err = model.UpdateWebsite(&company)
	require.Error(t, err)
}