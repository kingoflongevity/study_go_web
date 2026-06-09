package middlewares

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()

	// 注册自定义规则：username_len
	err := v.RegisterValidation("username_len", func(fl validator.FieldLevel) bool {
		name := fl.Field().String()
		length := len(name)
		// 限制长度在 3 到 12 之间
		return length >= 3 && length <= 12
	})
	if err != nil {
		return nil
	}

	return &CustomValidator{
		v,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
