https://github.com/foxcpp/go-mockdns/blob/master/example_test.go

    srv, _ := mockdns.NewServer(map[string]mockdns.Zone{
        "example.org.": {
            A: []string{"1.2.3.4"},
        },
    }, false)
    defer srv.Close()

    srv.PatchNet(net.DefaultResolver)
    defer mockdns.UnpatchNet(net.DefaultResolver)

    addrs, err := net.LookupHost("example.org")
    fmt.Println(addrs, err)
