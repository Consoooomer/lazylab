package endpoint

type Endpoint string

func ToString(endpoint Endpoint) string {
	return string(endpoint)
}

func FromString(endpoint string) Endpoint {
	return Endpoint(endpoint)
}
