
// RecordsTag automatically generated by SDKgen please do not edit this file manually
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

type RecordsTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll List records in a table. Note that table names and table ids can be used interchangeably. We recommend using table IDs so you don&#039;t need to modify your API request when your table name changes.
func (client *RecordsTag) GetAll(baseId string, tableIdOrName string, timeZone string, userLocale string, pageSize int, maxRecords int, offset string, view string, sort string, filterByFormula string, cellFormat string, fields string, returnFieldsByFieldId bool, recordMetadata string) (RecordCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName

    queryParams := make(map[string]interface{})
    queryParams["timeZone"] = timeZone
    queryParams["userLocale"] = userLocale
    queryParams["pageSize"] = pageSize
    queryParams["maxRecords"] = maxRecords
    queryParams["offset"] = offset
    queryParams["view"] = view
    queryParams["sort"] = sort
    queryParams["filterByFormula"] = filterByFormula
    queryParams["cellFormat"] = cellFormat
    queryParams["fields"] = fields
    queryParams["returnFieldsByFieldId"] = returnFieldsByFieldId
    queryParams["recordMetadata"] = recordMetadata

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName", pathParams))
    if err != nil {
        return RecordCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return RecordCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return RecordCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return RecordCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response RecordCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return RecordCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return RecordCollection{}, errors.New("the server returned an unknown status code")
    }
}

// Get Retrieve a single record. Any &quot;empty&quot; fields (e.g. &quot;&quot;, [], or false) in the record will not be returned.
func (client *RecordsTag) Get(baseId string, tableIdOrName string, recordId string) (Record, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId", pathParams))
    if err != nil {
        return Record{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return Record{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Record{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Record{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Record
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Record{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return Record{}, errors.New("the server returned an unknown status code")
    }
}

// Replace Updates a single record. Table names and table ids can be used interchangeably. We recommend using table IDs so you don&#039;t need to modify your API request when your table name changes.
func (client *RecordsTag) Replace(baseId string, tableIdOrName string, recordId string, payload Record) (Record, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId", pathParams))
    if err != nil {
        return Record{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Record{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("PUT", u.String(), reqBody)
    if err != nil {
        return Record{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Record{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Record{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Record
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Record{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return Record{}, errors.New("the server returned an unknown status code")
    }
}

// Update Updates a single record. Table names and table ids can be used interchangeably. We recommend using table IDs so you don&#039;t need to modify your API request when your table name changes.
func (client *RecordsTag) Update(baseId string, tableIdOrName string, recordId string, payload Record) (Record, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId", pathParams))
    if err != nil {
        return Record{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Record{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("PATCH", u.String(), reqBody)
    if err != nil {
        return Record{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Record{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Record{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Record
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Record{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return Record{}, errors.New("the server returned an unknown status code")
    }
}



func NewRecordsTag(httpClient *http.Client, parser *sdkgen.Parser) *RecordsTag {
	return &RecordsTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}