
// FieldsTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type FieldsTag struct {
    internal *sdkgen.TagAbstract
}



// Create 
func (client *FieldsTag) Create(baseId string, tableId string, payload Field) (Field, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableId"] = tableId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/meta/bases/:baseId/tables/:tableId/fields", pathParams))
    if err != nil {
        return Field{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Field{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return Field{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Field{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Field{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Field
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Field{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        case 403:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        case 404:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        case 500:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        default:
            return Field{}, errors.New("the server returned an unknown status code")
    }
}

// Update 
func (client *FieldsTag) Update(baseId string, tableId string, columnId string, payload Field) (Field, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableId"] = tableId
    pathParams["columnId"] = columnId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/meta/bases/:baseId/tables/:tableId/fields/:columnId", pathParams))
    if err != nil {
        return Field{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Field{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("PATCH", u.String(), reqBody)
    if err != nil {
        return Field{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Field{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Field{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Field
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Field{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        case 403:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        case 404:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        case 500:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Field{}, err
            }

            return Field{}, &ErrorException{
                Payload: response,
            }
        default:
            return Field{}, errors.New("the server returned an unknown status code")
    }
}



func NewFieldsTag(httpClient *http.Client, parser *sdkgen.Parser) *FieldsTag {
	return &FieldsTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
