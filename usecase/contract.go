package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	fileVm "gitlab.com/s2.1-backend/shm-file-management-svc/domain/view_models"
	"gitlab.com/s2.1-backend/shm-package-svc/jwe"
	"gitlab.com/s2.1-backend/shm-package-svc/jwt"
	"gitlab.com/s2.1-backend/shm-package-svc/redis"
	"gitlab.com/s2.1-backend/shm-package-svc/responses"
	"gitlab.com/s2.1-backend/shm-product-svc/domain/view_models"
	"gorm.io/gorm"
)

type Contract struct {
	UserID        string
	RoleID        string
	App           *fiber.App
	DB            *gorm.DB
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	Validate      *validator.Validate
	Redis         redis.RedisClient
	Translator    ut.Translator
}

const (
	// Default limit for pagination
	defaultLimit = 10

	// Max limit for pagination
	maxLimit = 100

	// Default order by
	defaultOrderBy = "created_at"

	// Default sort
	defaultSort = "desc"

	// Default last page for pagination
	defaultLastPage = 0

	// Status Enum Brand
	StatusBrandActive = "Active"
	StatusBrandBanned = "Banned"
)

func (uc Contract) SetPaginationParameter(page, limit int64, order, sort string) (int64, int64, int64, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) SetPaginationResponse(page, limit, total int64) (res view_models.PaginationVm) {
	var lastPage int64

	if total > 0 {
		lastPage = total / limit
		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}

func (uc Contract) OpenFile(f string) (file *os.File, err error) {
	r, err := os.Open("../../domain/files/" + f)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc Contract) Upload(client *http.Client, values map[string]io.Reader) (data *responses.FactoryBaseResponse, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", os.Getenv("FILE_ENDPOINT"), &b)
	if err != nil {
		return nil, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc Contract) ExportBase(filename string) (link *fileVm.FileVm, err error) {
	file, err := uc.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	values := map[string]io.Reader{
		"file":      file,
		"file_type": strings.NewReader("spreadsheet"),
	}
	res, err := uc.Upload(client, values)
	if err != nil {
		return nil, err
	}
	data := res.Data.(map[string]interface{})
	var filesVm fileVm.FileVm
	for k, v := range data {
		switch k {
		case "file_id":
			filesVm.ID = v.(string)
		case "name":
			filesVm.Name = v.(string)
		case "extension":
			filesVm.Ext = v.(string)
		case "path":
			path := strings.Split(v.(string), "?")
			filesVm.Path = path[0]
		}
	}
	return &filesVm, nil
}
