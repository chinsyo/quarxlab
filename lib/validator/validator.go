package validators

import (
    //"regexp"
    //"errors"
)

type Validator struct {

}

func (v *Validator) IsEmail(email string) (bool, error) {
    return true, nil
}

func (v *Validator) IsStrongPassword(password string) (bool, error) {
    return true, nil
}





