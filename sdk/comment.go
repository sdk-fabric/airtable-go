
// comment automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


type Comment struct {
    Id string `json:"id"`
    CreatedTime string `json:"createdTime"`
    LastUpdatedTime string `json:"lastUpdatedTime"`
    Text string `json:"text"`
    ParentCommentId string `json:"parentCommentId"`
    Reactions string `json:"reactions"`
    Author *CommentAuthor `json:"author"`
}
