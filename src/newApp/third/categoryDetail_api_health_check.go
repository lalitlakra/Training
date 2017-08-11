package third

type CategoryDetailsHealthCheck struct {
}

func (n CategoryDetailsHealthCheck) GetName() string {
	return "Hello category Details"
}

func (n CategoryDetailsHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
