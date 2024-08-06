
// CommentsTag automatically generated by SDKgen please do not edit this file manually
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

type CommentsTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll 
func (client *CommentsTag) GetAll(baseId string, tableIdOrName string, recordId string) (CommentCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId/comments", pathParams))
    if err != nil {
        return CommentCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return CommentCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return CommentCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return CommentCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response CommentCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return CommentCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentCollection{}, err
            }

            return CommentCollection{}, &ErrorException{
                Payload: response,
            }
        case 403:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentCollection{}, err
            }

            return CommentCollection{}, &ErrorException{
                Payload: response,
            }
        case 404:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentCollection{}, err
            }

            return CommentCollection{}, &ErrorException{
                Payload: response,
            }
        case 500:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentCollection{}, err
            }

            return CommentCollection{}, &ErrorException{
                Payload: response,
            }
        default:
            return CommentCollection{}, errors.New("the server returned an unknown status code")
    }
}

// Create 
func (client *CommentsTag) Create(baseId string, tableIdOrName string, recordId string, payload Comment) (Comment, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId/comments", pathParams))
    if err != nil {
        return Comment{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Comment{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return Comment{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Comment{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Comment{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Comment
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Comment{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        case 403:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        case 404:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        case 500:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        default:
            return Comment{}, errors.New("the server returned an unknown status code")
    }
}

// Update 
func (client *CommentsTag) Update(baseId string, tableIdOrName string, recordId string, rowCommentId string, payload Comment) (Comment, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId
    pathParams["rowCommentId"] = rowCommentId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId/comments/:rowCommentId", pathParams))
    if err != nil {
        return Comment{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return Comment{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("PATCH", u.String(), reqBody)
    if err != nil {
        return Comment{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Comment{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Comment{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response Comment
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return Comment{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        case 403:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        case 404:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        case 500:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return Comment{}, err
            }

            return Comment{}, &ErrorException{
                Payload: response,
            }
        default:
            return Comment{}, errors.New("the server returned an unknown status code")
    }
}

// Delete 
func (client *CommentsTag) Delete(baseId string, tableIdOrName string, recordId string, rowCommentId string) (CommentDeleteResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["baseId"] = baseId
    pathParams["tableIdOrName"] = tableIdOrName
    pathParams["recordId"] = recordId
    pathParams["rowCommentId"] = rowCommentId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/v0/:baseId/:tableIdOrName/:recordId/comments/:rowCommentId", pathParams))
    if err != nil {
        return CommentDeleteResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("DELETE", u.String(), nil)
    if err != nil {
        return CommentDeleteResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return CommentDeleteResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return CommentDeleteResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response CommentDeleteResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return CommentDeleteResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentDeleteResponse{}, err
            }

            return CommentDeleteResponse{}, &ErrorException{
                Payload: response,
            }
        case 403:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentDeleteResponse{}, err
            }

            return CommentDeleteResponse{}, &ErrorException{
                Payload: response,
            }
        case 404:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentDeleteResponse{}, err
            }

            return CommentDeleteResponse{}, &ErrorException{
                Payload: response,
            }
        case 500:
            var response Error
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommentDeleteResponse{}, err
            }

            return CommentDeleteResponse{}, &ErrorException{
                Payload: response,
            }
        default:
            return CommentDeleteResponse{}, errors.New("the server returned an unknown status code")
    }
}



func NewCommentsTag(httpClient *http.Client, parser *sdkgen.Parser) *CommentsTag {
	return &CommentsTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
