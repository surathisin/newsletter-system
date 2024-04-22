package storage

var emails = make(map[string]bool)

func AddEmail(email string) error {
    emails[email] = true
    return nil
}

func RemoveEmail(email string) error {
    delete(emails, email)
    return nil
}

func GetEmails() ([]string, error) {
    var result []string
    for email := range emails {
        result = append(result, email)
    }
    return result, nil
}
