# Writing a Go client for your RESTful API

https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c

## Your first Go client SDK

First off, you are going to need a Client struct to hold information about where to find the API you are going to consume. It will also contain a reference to the http.Client used to make request to our API.

```
package chatsby
type Client struct {
    BaseURL   *url.URL
    UserAgent string
 
    httpClient *http.Client
}
```

Pro tip: By making the base URL configurable we can make it testable by passing the URL of a httptest.Server.

Next, let us create a method for listing users. We create a complete URL from our base URL and relative path, and use it to build a request that we can send using our http.Client.

```
func (c *Client) ListUsers() ([]User, error) {
    rel := &url.URL{Path: "/users"}
    u := c.BaseURL.ResolveReference(rel)
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)
 
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    var users []User
    err = json.NewDecoder(resp.Body).Decode(&users)
    return users, err
}
```

Note: Keep in mind that the Transport is only able to reuse the HTTP connection if the response body has been drained completely. We can ensure this is the case by calling io.Copy(ioutil.Discard, resp.Body) right before resp.Body.Close is called. Check out this thread to learn more.

Edit: As pointed out by some, in practice you probably want to use io.CopyN in order to not drain an unbounded number of bytes. Note though that draining the body is really only necessary if you can realistically expect junk data to follow your response.

When we implement the rest of the API operations however, we see that there are some things that we can extract to their own methods. By extracting our own newRequest() and do() methods, we can reuse them for all our API operations:

```
func (c *Client) ListUsers() ([]User, error) { 
    req, err := c.newRequest("GET", "/users", nil)
    if err != nil {
        return nil, err
    }
    var users []User
    _, err = c.do(req, &users)
    return users, err
}
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
    rel := &url.URL{Path: path}
    u := c.BaseURL.ResolveReference(rel)
    var buf io.ReadWriter
    if body != nil {
        buf = new(bytes.Buffer)
        err := json.NewEncoder(buf).Encode(body)
        if err != nil {
            return nil, err  
        }
    }
    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }
    if body != nil {  
        req.Header.Set("Content-Type", "application/json") 
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)
    return req, nil
}
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    err = json.NewDecoder(resp.Body).Decode(v)
    return resp, err
}
```

You can even provide Go methods for each HTTP method as a further convenience. Depending on how consistent your requests and responses are, you can conveniently encode/decode them in the newRequest() and do() methods, or where you do the actual request. When we are done, the complete API will look something like this:

```
func (c *Client) SendMessage(channel, msg string) error
func (c *Client) History(channel string) ([]Message, error)
func (c *Client) CreateChannel(channel string) error
func (c *Client) ListUsers() ([]User, error)
```

Often, a Client with a bunch of methods is all you are going to need, especially if the REST API is fairly small. Both docker/docker and nlopes/slack are examples of two decent-sized API clients that are taking this pattern to the limit. However, when the API grows larger and we start exposing several services, you might want to consider using service objects.

## Service objects

At some point, you are going to want to break up your package into multiple files, e.g. client.go, chat.go, channels.go, and users.go. docker/docker deals with this by communicating services in the filename, e.g. container_list.go, and by prefixing each method with the service name, e.g. client.ContainerList() (check out their Client object on godoc.org).

If you do not like spreading methods across multiple files, another approach is to break down your API into service objects that group related operations. Each service object would then contain a reference to the Client. This approach is used by for example google/go-github.

```
type ChannelService type {
    client *Client
}
func (s *ChannelService) History(channel string) ([]Message, error)
func (s *ChannelService) Create(channel string) error
```

A common practice is to have service objects as fields in the Client. When creating the client, you can then pass a reference to the Client to each service object in the NewClient function:

```
type Client struct {
    httpClient *http.Client
 
    Chat     *ChatService
    Channels *ChannelService
    Users    *UserService
}
func NewClient(httpClient *http.Client) *Client {
    if httpClient == nil {
        httpClient = http.DefaultClient
    }
    c := &Client {httpClient: httpClient}
    c.ChatService = &ChatService{client: c}
    c.ChannelService = &ChannelService{client: c}
    c.UserService = &UserService{client: c}
}
```

## Configuring the http.Client

Depending on how customizable you want the API to be, you can go for simply passing an http.Client, or you could go for the functional options approach, letting your users pretty much customize all they want. For example, if you provide a library for a hosted service you might rely on a defaultEndpoint, while if you are talking to an on-site installation you might provide a func BaseURL(u string) func(*Client) option.

Regardless of the approach you go with, you should provide a means of letting users configure their own http.Client, e.g. setting sensible timeout values. This also allows you to handle authentication outside your Client. Here is an example from google/go-github of how you can authenticate using OAuth2 access tokens.

```
ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "... your access token ..."},
)
tc := oauth2.NewClient(oauth2.NoContext, ts)
client := github.NewClient(tc)
```

Pro tip: If you have several ways of letting users authenticate themselves, consider providing helper functions for creating authenticated http.Clients.

Keeping a reference to a http.Client lets us handle authentication, timeouts and other things independently from your API. Allowing your users to pass their own HTTP client used by your library lets them configure settings that make sense in their context.

## Tips and tricks

### Dealing with errors

Arguably the simplest way to deal with errors coming from you API is to simply check whether the status code in the response is anywhere between 200≤ and <400, or otherwise return an error.

Depending on how your RESTful API returns errors, consider converting them into proper Go errors instead of generic API errors, e.g. check for 404 and instead return a instance of ErrChannelNotFound. Also make sure to document that ErrChannelNotFound is returned when channel is not found. This makes for a more pleasant and consistent experience for your users.

### Pointer fields

To distinguish between unset fields and zero-value fields, structs can use pointer values.

```
type Channel struct {
    Name    *string `json:"name,omitempty"`
    Private *bool   `json:"private,omitempty"`
}
```

This can be helpful but it can also make for a rather unsightly code when using it. In this case make sure to provide helper functions for creating pointers to common types.

```
ch := &chatsby.Channel{
    Name:    chatsby.String("afterwork"),
    Private: chatsby.Bool(true),
}
```

This is obviously more tedious than using non-pointer values so do this only if it makes sense for your API. Check out google/go-github and aws/aws-sdk-go to see this in action. Compare with tambet/go-asana where values are just as effective.

## Summary

Go makes it easy writing production-ready client packages that allow your users to consume your RESTful services safely and efficiently from their Go applications. In this blog post you have seen how you can get started writing your own. If you have experience writing client SDKs in Go and you feel I left something out, please let me know by commenting on this post!